#!/usr/bin/env ruby
require 'rbs'
require 'json'

######WARNING!!!#######
#this program is WIP!!#
#######################

# Type conversion functions
def extract_type(type)
  case type
  when RBS::Types::Bases::Self, RBS::Types::Bases::Instance
    "Self"
  when RBS::Types::Bases::Bool
    "Bool"
  when RBS::Types::Bases::Void
    "Nil"
  when RBS::Types::Bases::Any
    "Untyped"
  when RBS::Types::ClassInstance
    extract_class_instance_type(type)
  when RBS::Types::Optional
    "Optional#{extract_type(type.type)}"
  when RBS::Types::Union
    type.types.map { |t| extract_type(t) }
  when RBS::Types::Variable
    "Unify"
  when RBS::Types::Alias
    extract_alias_type(type)
  else
    "Untyped"
  end
end

def extract_class_instance_type(type)
  case type.name.name.to_s
  when "Array" then "Array"
  when "Integer" then "Int"
  when "String" then "String"
  else type.name.to_s
  end
end

def extract_alias_type(type)
  case type.name.name.to_s
  when "int" then "Int"
  when "string" then "String"
  when "bool", "boolish" then "Bool"
  else "Untyped"
  end
end

def normalize_type(type)
  result = extract_type(type)
  result.is_a?(Array) ? result : [result]
end

# Argument processing functions
def process_required_args(required_positionals)
  required_positionals.map { |param| { "type" => normalize_type(param.type) } }
end

def process_optional_args(optional_positionals)
  optional_positionals.map do |param|
    types = normalize_type(param.type)
    { "type" => types.map { |t| "Default#{t}" } }
  end
end

def process_arguments(method_type)
  args = []
  args.concat(process_required_args(method_type.required_positionals))
  args.concat(process_optional_args(method_type.optional_positionals))
  args << { "type" => ["Untyped"] } if method_type.rest_positionals
  args
end

def extract_block_params(method_type)
  return nil unless method_type.block

  params = method_type.block.type.required_positionals.map do |param|
    extract_type(param.type)
  end
  [params].flatten
end

# Module method collection
def collect_module_singleton_methods(module_decl)
  methods = []
  module_decl.members.each do |member|
    next unless member.is_a?(RBS::AST::Members::MethodDefinition) && member.singleton?
    methods << member
  end
  methods
end

def build_module_method(member)
  return nil unless member.overloads.any?

  overload = member.overloads.first
  args = process_arguments(overload.method_type.type)
  return_type = normalize_type(overload.method_type.type.return_type)

  method_data = { "name" => member.name.to_s }
  method_data["arguments"] = args unless args.empty?
  method_data["return_type"] = { "type" => return_type }
  method_data
end

def build_module_output(module_decl, module_methods)
  return nil if module_methods.empty?

  {
    "frame" => "Builtin",
    "class" => module_decl.name.name.to_s,
    "instance_methods" => [],
    "class_methods" => module_methods.compact
  }
end

# RBS parsing and declaration extraction
def parse_rbs_file(file_path)
  buffer = RBS::Buffer.new(name: file_path, content: File.read(file_path))
  result = RBS::Parser.parse_signature(buffer)
  result[2] # declarations
end

def extract_module_and_classes(decls)
  module_decls = decls.select { |d| d.is_a?(RBS::AST::Declarations::Module) }

  if module_decls.any?
    frame_name = module_decls.first.name.name.to_s

    all_class_decls = []
    module_decls.each do |mod_decl|
      all_class_decls.concat(mod_decl.members.select { |m| m.is_a?(RBS::AST::Declarations::Class) })
    end

    { module: module_decls.first, classes: all_class_decls, frame: frame_name }
  else
    frame_name = "Builtin"
    class_decls = decls.select { |d| d.is_a?(RBS::AST::Declarations::Class) }
    { module: nil, classes: class_decls, frame: frame_name }
  end
end

# Collect instance methods from nested Methods module
def collect_nested_instance_methods_from_methods_module(klass_decl)
  methods_module = klass_decl.members.find do |m|
    m.is_a?(RBS::AST::Declarations::Module) && m.name.name.to_s == "Methods"
  end

  return [] unless methods_module

  collect_instance_methods(methods_module)
end

# Instance method collection
def group_instance_methods(klass_decl)
  methods_map = {}
  method_order = []

  klass_decl.members.each do |member|
    next unless member.is_a?(RBS::AST::Members::MethodDefinition)
    next unless member.instance?
    next if member.visibility == :private
    next if member.name.to_s == "initialize"

    method_name = member.name.to_s
    unless methods_map.key?(method_name)
      methods_map[method_name] = []
      method_order << method_name
    end
    methods_map[method_name] << member
  end

  { map: methods_map, order: method_order }
end

def merge_method_overloads(method_name, all_overloads)
  all_arg_types = []
  all_return_types = []
  block_params = nil
  has_optional_args = false

  all_overloads.each do |overload|
    method_type = overload.method_type.type
    arg_types = []

    # Collect argument types
    method_type.required_positionals.each { |p| arg_types << normalize_type(p.type) }
    method_type.optional_positionals.each do |p|
      arg_types << normalize_type(p.type)
      has_optional_args = true
    end
    arg_types << ["Untyped"] if method_type.rest_positionals

    all_arg_types << arg_types unless arg_types.empty?
    all_return_types << normalize_type(method_type.return_type)
    block_params ||= extract_block_params(overload.method_type)
  end

  build_instance_method_data(method_name, all_arg_types, all_return_types, block_params, has_optional_args)
end

def build_instance_method_data(method_name, all_arg_types, all_return_types, block_params, has_optional_args)
  method_data = { "name" => method_name }

  # Merge arguments
  if all_arg_types.any?
    max_args = all_arg_types.max_by(&:size)
    arguments = []

    max_args.each_with_index do |arg_type, idx|
      merged_types = all_arg_types.map { |args| args[idx] }.compact.flatten.uniq

      if has_optional_args && idx >= (max_args.size - 1)
        merged_types = merged_types.map { |t| t.start_with?("Default") ? t : "Default#{t}" }
      end

      arguments << { "type" => merged_types }
    end

    method_data["arguments"] = arguments unless arguments.empty?
  else
    method_data["arguments"] = []
  end

  # Add block parameters
  method_data["block_parameters"] = block_params if block_params && !block_params.empty?

  # Build return type
  merged_return_types = all_return_types.flatten.uniq
  return_data = { "type" => merged_return_types }
  return_data["is_conditional"] = true if merged_return_types.size > 1
  return_data["is_destructive"] = true if method_name.end_with?("!")
  method_data["return_type"] = return_data

  method_data
end

def collect_instance_methods(klass_decl)
  grouped = group_instance_methods(klass_decl)
  instance_methods = []

  grouped[:order].each do |method_name|
    members = grouped[:map][method_name]
    all_overloads = members.flat_map(&:overloads)
    next if all_overloads.empty?

    method_data = merge_method_overloads(method_name, all_overloads)
    instance_methods << method_data
  end

  instance_methods
end

# Class method collection
def find_initialize_method(klass_decl)
  klass_decl.members.find do |member|
    member.is_a?(RBS::AST::Members::MethodDefinition) &&
    member.instance? &&
    member.name.to_s == "initialize"
  end
end

def build_new_method(klass_decl, initialize_method)
  return nil unless initialize_method&.overloads&.any?

  overload = initialize_method.overloads.first
  args = process_arguments(overload.method_type.type)

  {
    "name" => "new",
    "arguments" => args,
    "return_type" => { "type" => [klass_decl.name.name.to_s] }
  }
end

def collect_singleton_methods(klass_decl)
  methods = []

  klass_decl.members.each do |member|
    next unless member.is_a?(RBS::AST::Members::MethodDefinition)
    next unless member.singleton?
    next if member.visibility == :private
    next unless member.overloads.any?

    overload = member.overloads.first
    args = process_arguments(overload.method_type.type)
    return_type = normalize_type(overload.method_type.type.return_type)

    method_data = { "name" => member.name.to_s }
    method_data["arguments"] = args unless args.empty?
    method_data["return_type"] = { "type" => return_type }
    methods << method_data
  end

  methods
end

def collect_nested_class_methods_from_class_methods_module(klass_decl)
  class_methods_module = klass_decl.members.find do |m|
    m.is_a?(RBS::AST::Declarations::Module) && m.name.name.to_s == "ClassMethods"
  end

  return [] unless class_methods_module

  collect_instance_methods(class_methods_module)
end

def collect_class_methods(klass_decl)
  methods = []
  initialize_method = find_initialize_method(klass_decl)
  new_method = build_new_method(klass_decl, initialize_method)
  methods << new_method if new_method
  methods.concat(collect_singleton_methods(klass_decl))
  methods.concat(collect_nested_class_methods_from_class_methods_module(klass_decl))
  methods
end

# Class extends collection
def collect_extends(klass_decl)
  extends = []
  klass_decl.members.each do |member|
    extends << member.name.name.to_s if member.is_a?(RBS::AST::Members::Include)
  end
  extends
end

# Build class output
def build_class_output(klass_decl, frame_name)
  output = {
    "frame" => frame_name,
    "class" => klass_decl.name.name.to_s
  }

  extends = collect_extends(klass_decl)
  output["extends"] = extends unless extends.empty?

  instance_methods = collect_instance_methods(klass_decl)
  nested_methods = collect_nested_instance_methods_from_methods_module(klass_decl)
  all_instance_methods = instance_methods + nested_methods
  output["instance_methods"] = all_instance_methods

  class_methods = collect_class_methods(klass_decl)
  output["class_methods"] = class_methods unless class_methods.empty?

  output
end

# Output handling
def generate_filename(frame, class_name)
  if frame == "Builtin"
    "#{class_name.downcase}.json"
  else
    "#{frame.downcase}_#{class_name.downcase}.json"
  end
end

def output_to_files(all_outputs, output_dir = 'myjson')
  Dir.mkdir(output_dir) unless Dir.exist?(output_dir)

  all_outputs.each do |output|
    filename = generate_filename(output["frame"], output["class"])
    filepath = File.join(output_dir, filename)
    File.write(filepath, JSON.pretty_generate(output))
    puts "Generated #{filepath}"
  end

  puts "\nAll JSON files generated in #{output_dir}/ directory"
end

def output_to_console(all_outputs)
  all_outputs.each do |output|
    p JSON.pretty_generate(output)
  end
end

# Main execution
input_file = ARGV[0] || 'input_sample.rbs'
decls = parse_rbs_file(input_file)
parsed = extract_module_and_classes(decls)

if parsed[:classes].empty? && (!parsed[:module] || collect_module_singleton_methods(parsed[:module]).empty?)
  puts "No class declaration found"
  exit 1
end

all_outputs = []

# Process module singleton methods if present
if parsed[:module]
  module_methods = collect_module_singleton_methods(parsed[:module])
  module_methods_data = module_methods.map { |m| build_module_method(m) }
  module_output = build_module_output(parsed[:module], module_methods_data)
  all_outputs << module_output if module_output
end

# Group classes by name to merge duplicates
classes_by_name = {}
parsed[:classes].each do |klass_decl|
  class_name = klass_decl.name.name.to_s
  classes_by_name[class_name] ||= []
  classes_by_name[class_name] << klass_decl
end

# Process all classes (merge duplicates)
classes_by_name.each do |class_name, klass_decls|
  if klass_decls.size == 1
    all_outputs << build_class_output(klass_decls.first, parsed[:frame])
  else
    all_instance_methods = []
    all_class_methods = []
    all_extends = []

    klass_decls.each do |klass_decl|
      single_output = build_class_output(klass_decl, parsed[:frame])
      all_instance_methods.concat(single_output["instance_methods"])
      all_class_methods.concat(single_output["class_methods"] || [])
      all_extends.concat(single_output["extends"] || [])
    end

    merged_output = {
      "frame" => parsed[:frame],
      "class" => class_name,
      "instance_methods" => all_instance_methods.uniq { |m| m["name"] },
      "class_methods" => all_class_methods.uniq { |m| m["name"] }
    }
    merged_output["extends"] = all_extends.uniq unless all_extends.empty?
    merged_output.delete("class_methods") if merged_output["class_methods"].empty?

    all_outputs << merged_output
  end
end

# Execute output
if ARGV.include?("--test")
  output_to_console(all_outputs)
else
  output_to_files(all_outputs)
end
