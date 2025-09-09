dbtp 1
# mruby comprehensive syntax test file

# Basic literals and numeric formats
number = 42
float_num = 3.14
hex_num = 0xFF
octal_num = 0o755
binary_num = 0b1010
string = "hello world"
single_quote_string = 'single quotes'
symbol = :test
quoted_symbol = :"quoted symbol"
boolean_true = true
boolean_false = false
nil_value = nil

# String interpolation and escapes
name = "mruby"
greeting = "Hello, #{name}!"
math_interpolation = "Result: #{2 + 3 * 4}"
escape_sequences = "Line1\nLine2\tTabbed\\"

# Percent notation strings
percent_string = %q(single quoted)

# TODO01
percent_interpolated = %Q(with #{name} interpolation)
word_array = %w(one two three)
symbol_array = %i(sym1 sym2 sym3)

# Arrays with various syntax
array = [1, 2, 3, "string", :symbol]
empty_array = []
nested_array = [[1, 2], [3, 4]]
trailing_comma = [1, 2, 3,]
multiline_array = [
  1,
  2,
  3
]

# Hashes with different syntax
hash = {key: "value", "string_key" => 42, :symbol_key => true}
empty_hash = {}
arrow_hash = { "key" => "value", :sym => 123 }
mixed_hash = { a: 1, "b" => 2, :c => 3 }
trailing_comma_hash = { a: 1, b: 2, }

# Variables and scope
local_var = "local"
@instance_var = "instance"
@@class_var = "class"
$global_var = "global"

# Constants with nested access
CONSTANT = "constant value"
NESTED_CONSTANT = { inner: "nested value" }

# Ranges
inclusive_range = 1..10
exclusive_range = 1...10
string_range = "a".."z"

# Operators - Arithmetic
a = 10 + 5
b = 20 - 3
c = 4 * 6
d = 15 / 3
e = 17 % 5
f = 2 ** 8

# Operators - Comparison
equals = 5 == 5
not_equals = 5 != 3
less_than = 3 < 5
greater_than = 7 > 4
less_equal = 5 <= 5
greater_equal = 8 >= 6
spaceship = 5 <=> 3

# Operators - Logical
and_op = true && false
or_op = true || false
not_op = !true

# Operators - Bitwise
bit_and = 5 & 3
bit_or = 5 | 3
bit_xor = 5 ^ 3
left_shift = 5 << 2
right_shift = 20 >> 2

# Assignment operators
x = 10
x += 5
x -= 2
x *= 2
x /= 3
x %= 4
x **= 2

# Conditional assignment
y ||= "default"
z &&= "conditional"

# Multiple assignment
a, b = 1, 2
c, d, e = [3, 4, 5]
f, *rest = [6, 7, 8, 9]
first, *middle, last = [1, 2, 3, 4, 5]

# Method definitions with various parameters
def simple_method
  "simple"
end

def method_with_args(a, b)
  a + b
end

def method_with_default_args(a, b = 10, c = "default")
  [a, b, c]
end

def method_with_splat(*args)
  args.length
end

def method_with_keyword_args(name:, age: 25)
  "#{name} is #{age}"
end

def method_with_block(&block)
  block.call("hello") if block
end

def method_with_all(a, b = 2, *rest, x:, y: "default", &block)
  result = [a, b, rest, x, y]
  block ? block.call(result) : result
end

# Class definitions
class Animal
  attr_reader :name
  attr_writer :age
  attr_accessor :species
  
  def initialize(name)
    @name = name
    @age = 0
    @@count ||= 0
    @@count += 1
  end
  
  def speak
    "Generic animal sound"
  end
  
  def self.count
    @@count
  end
  
  def self.reset_count
    @@count = 0
  end
  
  private
  
  def private_method
    "This is private"
  end
  
  protected
  
  def protected_method
    "This is protected"
  end
end

class Dog < Animal
  def initialize(name, breed)
    super(name)
    @breed = breed
  end
  
  def speak
    "Woof! I'm #{@name}"
  end
  
  def breed
    @breed
  end
  
  def call_protected
    protected_method
  end
end

# Modules
module Walkable
  def walk
    "#{self.class} is walking"
  end
  
  def run(speed = "normal")
    "#{self.class} is running at #{speed} speed"
  end
end

module Flyable
  def fly(height = 100)
    "Flying at #{height} feet"
  end
end

class Bird
  include Walkable
  include Flyable
  
  def initialize(name)
    @name = name
  end
end

# Module functions and constants
module MathUtils
  PI = 3.14159
  
  def self.circle_area(radius)
    PI * radius * radius
  end
  
  module_function
  
  def square(x)
    x * x
  end
end

# Control structures - Conditionals
if true
  result1 = "if branch"
elsif false
  result1 = "elsif branch"
else
  result1 = "else branch"
end

unless false
  result2 = "unless true"
end

# Ternary operator
condition = true
result3 = condition ? "ternary true" : "ternary false"

# Case statements
value = 2
case value
when 1
  case_result = "one"
when 2, 3
  case_result = "two or three"
when 4..6
  case_result = "four to six"
when String
  case_result = "string type"
else
  case_result = "something else"
end


# Loops
counter = 0
while counter < 3
  counter += 1
end

counter = 0
until counter >= 3
  counter += 1
end

for i in 1..3
  # loop body
end

for item in ["a", "b", "c"]
  # iterate over array
end

# Loop control
5.times do |i|
  next if i == 2
  break if i == 4
end

# Blocks and iterators
[1, 2, 3, 4, 5].each do |item|
  item * 2
end

[1, 2, 3].map { |x| x ** 2 }



# Block with multiple parameters
{a: 1, b: 2, c: 3}.each do |key, value|
  "#{key}: #{value}"
end

# Yield and block_given?
def with_yield
  if block_given?
    yield("param1", "param2")
  else
    "no block given"
  end
end

with_yield do |a, b|
  "received #{a} and #{b}"
end

# Proc and Lambda
my_proc = Proc.new { |x| x * 2 }
my_lambda = lambda { |x| x + 1 }
stabby_lambda = ->(x) { x - 1 }

# Exception handling
begin
  raise StandardError, "Custom error message"
rescue StandardError => e
  error_message = e.message
rescue => general_error
  other_error = general_error.message
else
  no_error = "no exception raised"
ensure
  cleanup = "always executed"
end

# Custom exception
class CustomError < StandardError
  def initialize(message = "Custom error")
    super(message)
  end
end

# String methods and operations
str = "Hello World"
upcase_str = str.upcase
downcase_str = str.downcase
length = str.length
substring = str[0, 5]
gsub_result = str.gsub("World", "mruby")
split_result = str.split(" ")

# Array methods
arr = [3, 1, 4, 1, 5, 9, 2, 6]
sorted_arr = arr.sort
reversed_arr = arr.reverse
unique_arr = arr.uniq
first_elem = arr.first
last_elem = arr.last
arr.push(7)
popped = arr.pop
arr << 8

# Hash methods
h = {a: 1, b: 2, c: 3}
keys = h.keys
values = h.values
h[:d] = 4
deleted = h.delete(:b)

# Enumerable methods
numbers = [1, 2, 3, 4, 5]
doubled = numbers.map { |n| n * 2 }
sum = numbers.reduce(0) { |acc, n| acc + n }
all_positive = numbers.all? { |n| n > 0 }

# Method chaining
chained_result = "hello world".upcase.reverse.split("").join("-")

# Singleton methods
obj = Object.new
def obj.singleton_method
  "I'm a singleton method"
end

class << obj
  def another_singleton
    "Another singleton method"
  end
end

# Alias and alias_method
alias original_method simple_method
# alias_method :new_name, :simple_method

# Constants and nested constants
class Outer
  OUTER_CONST = "outer constant"
  
  class Inner
    INNER_CONST = "inner constant"
    
    def access_constants
      [INNER_CONST, Outer::OUTER_CONST]
    end
  end
  
  module NestedModule
    MODULE_CONST = "module constant"
  end
end

# Method visibility
class VisibilityTest
  def public_method
    "public"
  end
  
  private
  
  def private_method
    "private"
  end
  
  protected
  
  def protected_method
    "protected"
  end
  
  public
  
  def call_private
    private_method
  end
end

# Self keyword
class SelfTest
  def instance_method
    self
  end
  
  def self.class_method
    self
  end
end

# Return statements
def explicit_return
  return "explicit return"
  "unreachable"
end

def implicit_return
  "implicit return"
end

def early_return(condition)
  return "early" if condition
  "normal"
end

# Complex expressions and precedence
complex_expr = (1 + 2) * 3 / (4 - 2) + 5 % 2

# Array and Hash access with various syntax
arr_access = array[0]
arr_negative = array[-1]
hash_access = hash[:key]
hash_string_key = hash["string_key"]

# Defined? operator

# Super keyword
class Parent
  def greet(name)
    "Hello, #{name}!"
  end
end

class Child < Parent
  def greet(name)
    super + " How are you?"
  end
  
  def greet_default
    super("World")
  end
end

# Nested method calls
nested_call = "hello".upcase.downcase.capitalize.length

# Symbol to proc
symbol_to_proc = [1, 2, 3].map(&:to_s)

# Regular expressions (basic, if supported)

# Heredoc syntax (if supported)
heredoc = <<EOF
This is a heredoc
with multiple lines
EOF

# Global variables and special globals
$global_var = "global"
# $0, $1, $2 etc. for regex matches
# $$ for process id
# $? for last exit status

puts "mruby comprehensive syntax test complete"
