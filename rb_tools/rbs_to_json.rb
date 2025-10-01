#!/usr/bin/env ruby
require 'rbs'
require 'json'

######WARNING!!!#######
#this program is WIP!!#
#######################

def extract_type(type)
  case type
  when RBS::Types::Bases::Self
    "Self"
  when RBS::Types::Bases::Bool
    "Bool"
  when RBS::Types::Bases::Instance
    "Self"
  when RBS::Types::Bases::Void
    "Void"
  when RBS::Types::Bases::Any
    "Untyped"
  when RBS::Types::ClassInstance
    case type.name.name.to_s
    when "Array"
      "Array"
    when "Integer"
      "Int"
    when "String"
      "String"
    else
      # Handle namespaced classes (e.g., JS::Object)
      type.name.to_s
    end
  when RBS::Types::Optional
    inner = extract_type(type.type)
    "Optional#{inner}"
  when RBS::Types::Union
    return type.types.map { |t| extract_type(t) }
  when RBS::Types::Variable
    # Elem -> Unify
    "Unify"
  when RBS::Types::Alias
    case type.name.name.to_s
    when "int"
      "Int"
    when "string"
      "String"
    when "bool"
      "Bool"
    when "boolish"
      "Bool"
    else
      "Untyped"
    end
  else
    "Untyped"
  end
end

def normalize_type(type)
  result = extract_type(type)
  result.is_a?(Array) ? result : [result]
end

def is_optional_param?(param_name)
  param_name && param_name.to_s.start_with?('?')
end

# parsing rbs
input_file = ARGV[0] || 'input_sample.rbs'
buffer = RBS::Buffer.new(name: input_file, content: File.read(input_file))
result = RBS::Parser.parse_signature(buffer)

# result [buffer, _, declarations]
decls = result[2]

# find module or class
module_decl = decls.find { |d| d.is_a?(RBS::AST::Declarations::Module) }
class_decls = []
frame_name = "Builtin"
module_class_methods = []

if module_decl
  # extract all classes from module
  frame_name = module_decl.name.name.to_s
  class_decls = module_decl.members.select { |m| m.is_a?(RBS::AST::Declarations::Class) }

  # extract module's singleton methods (self.xxx methods)
  module_decl.members.each do |member|
    if member.is_a?(RBS::AST::Members::MethodDefinition) && member.singleton?
      module_class_methods << member
    end
  end
else
  # find standalone classes
  class_decls = decls.select { |d| d.is_a?(RBS::AST::Declarations::Class) }
end

if class_decls.empty? && module_class_methods.empty?
  puts "No class declaration found"
  exit 1
end

# process all classes
all_outputs = []

# If module has singleton methods, add module as a class
if module_decl && module_class_methods.any?
  module_output = {
    "frame" => "Builtin",
    "class" => module_decl.name.name.to_s,
    "instance_methods" => [],
    "class_methods" => []
  }

  # Process module's singleton methods as class methods
  module_class_methods.each do |member|
    method_data = {
      "name" => member.name.to_s
    }

    if member.overloads.any?
      overload = member.overloads.first
      args = []

      overload.method_type.type.required_positionals.each do |param|
        args << {"type" => normalize_type(param.type)}
      end

      overload.method_type.type.optional_positionals.each do |param|
        types = normalize_type(param.type)
        args << {"type" => types.map { |t| "Default#{t}" }}
      end

      method_data["arguments"] = args unless args.empty?

      return_type = normalize_type(overload.method_type.type.return_type)
      method_data["return_type"] = {"type" => return_type}
    end

    module_output["class_methods"] << method_data
  end

  all_outputs << module_output
end

class_decls.each do |klass_decl|
  # merge overload methods
  methods_map = {}
  method_order = []

  klass_decl.members.each do |member|
    case member
    when RBS::AST::Members::MethodDefinition
      next unless member.instance?
      next if member.visibility == :private
      next if member.name.to_s == "initialize"  # skip initialize, handle as class method 'new'

      method_name = member.name.to_s
      if !methods_map.key?(method_name)
        methods_map[method_name] = []
        method_order << method_name
      end
      methods_map[method_name] << member
    end
  end

  # collect instance methods
  instance_methods = []

  method_order.each do |method_name|
    members = methods_map[method_name]
    # collect all overloads
    all_overloads = members.flat_map(&:overloads)

    next if all_overloads.empty?

    method_data = {
      "name" => method_name
    }

    # merge
    all_arg_types = []
    all_return_types = []
    block_params = nil
    has_optional_args = false

    all_overloads.each do |overload|
      # 引数の型を収集
      arg_types = []

      overload.method_type.type.required_positionals.each do |param|
        arg_types << normalize_type(param.type)
      end

      overload.method_type.type.optional_positionals.each do |param|
        arg_types << normalize_type(param.type)
        has_optional_args = true
      end

      if overload.method_type.type.rest_positionals
        arg_types << ["Untyped"]
      end

      all_arg_types << arg_types unless arg_types.empty?

      # collect return type
      return_type = normalize_type(overload.method_type.type.return_type)
      all_return_types << return_type

      # collect block params
      if overload.method_type.block && !block_params
        block_params = overload.method_type.block.type.required_positionals.map do |param|
          extract_type(param.type)
        end
        block_params = [block_params].flatten
      end
    end

    # merge arguments
    if all_arg_types.any?
      # 引数の数が最大のものを基準に
      max_args = all_arg_types.max_by(&:size)

      arguments = []
      max_args.each_with_index do |arg_type, idx|
        # calculate overload
        merged_types = all_arg_types.map { |args| args[idx] }.compact.flatten.uniq

        # isDefault
        if has_optional_args && idx >= (max_args.size - 1)
          merged_types = merged_types.map { |t| t.start_with?("Default") ? t : "Default#{t}" }
        end

        arguments << {"type" => merged_types}
      end

      method_data["arguments"] = arguments unless arguments.empty?
    else
      method_data["arguments"] = []
    end

    # insert block params
    if block_params && !block_params.empty?
      method_data["block_parameters"] = block_params
    end

    # merge return type
    merged_return_types = all_return_types.flatten.uniq
    return_type_data = {"type" => merged_return_types}

    # is_conditional
    if merged_return_types.size > 1
      return_type_data["is_conditional"] = true
    end

    # is_destructive
    if method_name.end_with?("!")
      return_type_data["is_destructive"] = true
    end

    method_data["return_type"] = return_type_data

    instance_methods << method_data
  end

  # collect class methods
  class_methods = []

  # find initialize method and convert to 'new'
  initialize_method = klass_decl.members.find do |member|
    member.is_a?(RBS::AST::Members::MethodDefinition) &&
    member.instance? &&
    member.name.to_s == "initialize"
  end

  if initialize_method && initialize_method.overloads.any?
    overload = initialize_method.overloads.first
    args = []

    overload.method_type.type.required_positionals.each do |param|
      args << {"type" => normalize_type(param.type)}
    end

    overload.method_type.type.optional_positionals.each do |param|
      types = normalize_type(param.type)
      args << {"type" => types.map { |t| "Default#{t}" }}
    end

    new_method = {
      "name" => "new",
      "arguments" => args,
      "return_type" => {"type" => [klass_decl.name.name.to_s]}
    }
    class_methods << new_method
  end

  # collect other class methods
  klass_decl.members.each do |member|
    case member
    when RBS::AST::Members::MethodDefinition
      next unless member.singleton?
      next if member.visibility == :private

      method_data = {
        "name" => member.name.to_s
      }

      if member.overloads.any?
        overload = member.overloads.first
        args = []

        overload.method_type.type.required_positionals.each do |param|
          args << {"type" => normalize_type(param.type)}
        end

        overload.method_type.type.optional_positionals.each do |param|
          types = normalize_type(param.type)
          args << {"type" => types.map { |t| "Default#{t}" }}
        end

        method_data["arguments"] = args unless args.empty?

        return_type = normalize_type(overload.method_type.type.return_type)
        method_data["return_type"] = {"type" => return_type}
      end

      class_methods << method_data
    end
  end

  # extends
  extends = []
  klass_decl.members.each do |member|
    case member
    when RBS::AST::Members::Include
      extends << member.name.name.to_s
    end
  end

  output = {
    "frame" => frame_name,
    "class" => klass_decl.name.name.to_s
  }

  output["extends"] = extends unless extends.empty?
  output["instance_methods"] = instance_methods
  output["class_methods"] = class_methods unless class_methods.empty?

  all_outputs << output
end

# output all classes
puts JSON.pretty_generate(all_outputs)
