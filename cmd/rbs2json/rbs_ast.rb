require "rbs"
require "json"

def strip_locations(obj)
  case obj
  when Hash
    obj.reject { |k, _| k == "location" || k == "buffer" }
       .transform_values { |v| strip_locations(v) }
  when Array
    obj.map { |v| strip_locations(v) }
  else
    obj
  end
end

file = ARGV[0]
unless file
  $stderr.puts "Usage: ruby rbs_ast.rb <file.rbs>"
  exit 1
end

source = File.read(file)
result = RBS::Parser.parse_signature(source)
decls = result[2] || result

json_decls = decls.map { |decl| strip_locations(JSON.parse(decl.to_json)) }
puts JSON.pretty_generate(json_decls)
