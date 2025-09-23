
def method_with_defaults(required, optional = "default", *rest)
  puts "Required: #{required}"
  puts "Optional: #{optional}"
  puts "Rest args: #{rest}"
end

method_with_defaults("required_value")
