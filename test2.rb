# mruby comprehensive syntax test file

# Basic literals
number = 42
float_num = 3.14
string = "hello world"
symbol = :test
boolean_true = true
boolean_false = false
nil_value = nil

# String interpolation
name = "mruby"
greeting = "Hello, #{name}!"

# Arrays
array = [1, 2, 3, "string", :symbol]
empty_array = []
nested_array = [[1, 2], [3, 4]]

# Hashes
hash = {key: "value", "string_key" => 42, :symbol_key => true}
empty_hash = {}

# Variables
local_var = "local"
@instance_var = "instance"
@@class_var = "class"
$global_var = "global"

# Constants
CONSTANT = "constant value"

# Method definitions
def simple_method
  "simple"
end

def method_with_args(a, b)
  a + b
end

def method_with_default_args(a, b = 10)
  a + b
end

def method_with_splat(*args)
  args.size
end

def method_with_block(&block)
  block.call if block
end

# Class definition
class TestClass
  def initialize(name)
    @name = name
  end
  
  def get_name
    @name
  end
  
  def set_name(name)
    @name = name
  end
  
  def self.class_method
    "class method"
  end
end

# Module definition
module TestModule
  def module_method
    "from module"
  end
end

# Class with module inclusion
class ExtendedClass
  include TestModule
  
  def test
    module_method
  end
end

# Inheritance
class ChildClass < TestClass
  def child_method
    "child #{get_name}"
  end
end

# Control structures
if true
  result = "if true"
elsif false
  result = "elsif"
else
  result = "else"
end

unless false
  result = "unless false"
end

# Case statement
value = 2
case value
when 1
  result = "one"
when 2, 3
  result = "two or three"
else
  result = "other"
end

# Loops
for i in 1..5
  puts i
end

while false
  break
end

until true
  next
end

# Iterators and blocks
[1, 2, 3].each do |item|
  puts item
end

[1, 2, 3].map { |x| x * 2 }

# Range
range1 = 1..10
range2 = 1...10

# Regular expressions (if supported)
# regex = /test/

# String operations
str = "test string"
str.upcase
str.downcase
str.length

# Array operations
arr = [1, 2, 3]
arr.push(4)
arr.pop
arr.first
arr.last

# Hash operations
h = {a: 1, b: 2}
h[:c] = 3
h.keys
h.values

# Operators
a = 1 + 2
b = 3 - 1
c = 2 * 3
d = 6 / 2
e = 7 % 3
f = 2 ** 3

# Comparison operators
1 == 1
1 != 2
1 < 2
2 > 1
1 <= 1
2 >= 2

# Logical operators
true && false
true || false
!true

# Assignment operators
x = 1
x += 2
x -= 1
x *= 2
x /= 2

# Multiple assignment
a, b = 1, 2
a, *rest = [1, 2, 3, 4]

# Exception handling
begin
  raise "error"
rescue => e
  puts e.message
ensure
  puts "cleanup"
end

# Method calls
obj = TestClass.new("test")
obj.get_name
obj.set_name("new name")
TestClass.class_method

# Block with parameters
def block_method
  yield(1, 2) if block_given?
end

block_method do |a, b|
  a + b
end

# Proc and lambda (if supported)
proc_obj = proc { |x| x * 2 }
lambda_obj = lambda { |x| x * 2 }

# Constants and scope
module Outer
  CONST = "outer"
  
  class Inner
    CONST = "inner"
    
    def test_const
      CONST
    end
  end
end

# Singleton methods
obj = "string"
def obj.singleton_method
  "singleton"
end

# Alias
alias new_name simple_method

# Defined?
defined?(simple_method)
defined?(@instance_var)

# Self
class SelfTest
  def test_self
    self
  end
end

# Return values
def return_test
  return "explicit return"
  "unreachable"
end

def implicit_return
  "implicit return"
end

# Nested method calls
"hello".upcase.downcase.length

# Complex expressions
result = (1 + 2) * (3 - 1) / 2

# Array/Hash access
arr[0]
hash[:key]
hash["string_key"]

# Conditional assignment
x ||= "default value"
y &&= "and assignment"

# Ternary operator
condition = true
result = condition ? "true" : "false"

puts "mruby syntax test complete"
