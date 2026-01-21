# mruby comprehensive syntax test file

# Basic literals and variables
x = 42
y = "hello world"
z = 3.14159
flag = true
empty_val = nil
symbol_val = :test_symbol

# Number formats
decimal_num = 123
hex_num = 0xFF
octal_num = 0777
binary_num = 0b1010

# String literals and interpolation
single_str = 'single quoted'
double_str = "double quoted"
interpolated_str = "Number is #{x}"
escaped_str = "Line 1\nLine 2\tTabbed"

# Arrays
simple_array = [1, 2, 3, 4, 5]
mixed_array = [1, "string", 3.14, true, nil, :symbol]
nested_array = [[1, 2], [3, 4], [5, 6]]
empty_array = []

# Hashes
simple_hash = {"key1" => "value1", "key2" => "value2"}
symbol_hash = {:name => "John", :age => 30}
new_hash_syntax = {name: "Jane", age: 25}
mixed_hash = {"string_key" => 1, :symbol_key => 2}

# Ranges
inclusive_range = 1..10
exclusive_range = 1...10

# Method definitions
def no_args_method
  puts "No arguments method"
end

def with_args_method(a, b, c)
  a + b + c
end

def default_args_method(x, y = 10, z = 20)
  x + y + z
end

def splat_method(*args)
  args.each { |arg| puts arg }
end

def block_method(&block)
  if block
    yield
  end
end

# Class definitions
class BasicClass
  def initialize(name)
    @name = name
  end
  
  def get_name
    @name
  end
  
  def set_name(new_name)
    @name = new_name
  end
  
  def self.class_method
    "This is a class method"
  end
end

class ChildClass < BasicClass
  def initialize(name, age)
    super(name)
    @age = age
  end
  
  def get_age
    @age
  end
  
  def get_info
    "Name: #{@name}, Age: #{@age}"
  end
end

# Module definition
module TestModule
  def module_method
    "Module method called"
  end
end

class ClassWithModule
  include TestModule
  
  def test_include
    module_method
  end
end

# Constants
GLOBAL_CONSTANT = 100
PI_VALUE = 3.14159

# Control flow - if/elsif/else
if x > 40
  puts "x is greater than 40"
elsif x == 40
  puts "x equals 40"
else
  puts "x is less than 40"
end

# unless
unless flag == false
  puts "flag is not false"
end

# Case statement
case x
when 1..10
  puts "x is between 1 and 10"
when 20..30
  puts "x is between 20 and 30"
when 42
  puts "x is exactly 42"
else
  puts "x is something else"
end

# Loops
# for loop
for i in 1..5
  puts "For loop: #{i}"
end

# while loop
counter = 0
while counter < 3
  puts "While loop: #{counter}"
  counter += 1
end

# until loop
countdown = 3
until countdown == 0
  puts "Countdown: #{countdown}"
  countdown -= 1
end

# Iterators and blocks
[1, 2, 3, 4, 5].each do |num|
  puts "Each: #{num}"
end

[1, 2, 3].each { |n| puts "Block: #{n * 2}" }

3.times do |i|
  puts "Times: #{i}"
end

# Exception handling
begin
  result = 10 / 0
rescue ZeroDivisionError => e
  puts "Caught division by zero: #{e.message}"
rescue => e
  puts "Caught other error: #{e.message}"
ensure
  puts "This always runs"
end

# Operators
# Arithmetic
add_result = 5 + 3
sub_result = 10 - 4
mul_result = 6 * 7
div_result = 20 / 4
mod_result = 17 % 5
pow_result = 2 ** 3

# Comparison
eq_result = (5 == 5)
ne_result = (5 != 3)
gt_result = (10 > 5)
lt_result = (3 < 8)
ge_result = (5 >= 5)
le_result = (3 <= 7)

# Logical
and_result = true && false
or_result = true || false
not_result = !true

# Assignment operators
assign_var = 10
assign_var += 5
assign_var -= 2
assign_var *= 2
assign_var /= 3

# String operations
str1 = "Hello"
str2 = "World"
concat_str = str1 + " " + str2
repeat_str = str1 * 3

# Array operations
test_array = [1, 2, 3]
test_array << 4
test_array.push(5)
first_elem = test_array[0]
last_elem = test_array[-1]
array_length = test_array.length

# Hash operations
test_hash = {a: 1, b: 2}
test_hash[:c] = 3
value_a = test_hash[:a]
hash_keys = test_hash.keys

# Multiple assignment
a, b, c = [10, 20, 30]
x, y = 100, 200

# Conditional assignment
var1 ||= "default"
var2 = "existing"
var2 ||= "won't change"

# Method chaining
# Instance variables and accessors
class Person
  def initialize(name, age)
    @name = name
    @age = age
  end

  def name
    @name
  end
  
  def name=(new_name)
    @name = new_name
  end
  
  def age
    @age
  end
  
  def age=(new_age)
    @age = new_age
  end
  
  def greet
    "Hi, I'm #{@name} and I'm #{@age} years old"
  end
end

# Class variables
class Counter
  @@count = 0
  
  def initialize
    @@count += 1
  end
  
  def self.count
    @@count
  end
  
  def get_count
    @@count
  end
end

# Method visibility
class VisibilityExample
  def public_method
    "This is public"
  end
  
  private
  
  def private_method
    "This is private"
  end
  
  protected
  
  def protected_method
    "This is protected"
  end
  
  public
  
  def call_private
    private_method
  end
end

# self keyword
class SelfExample
  def self.class_method
    "Called on class: #{self}"
  end
  
  def instance_method
    "Called on instance: #{self.class}"
  end
end

# Global variables
$global_var = "I'm global"

def access_global
  puts $global_var
end

# Nested classes
class Outer
  class Inner
    def inner_method
      "Inside inner class"
    end
  end
  
  def create_inner
    Inner.new
  end
end

# Module constants
module Constants
  VERSION = "1.0.0"
  MAX_SIZE = 1000
end

# Proc objects (if supported)
simple_proc = proc { |x| x * 2 }

# Method calls and object creation
no_args_method
result1 = with_args_method(1, 2, 3)
result2 = default_args_method(5)
result3 = default_args_method(5, 15)

splat_method(1, 2, 3, 4)

block_method do
  puts "Block executed"
end

# Object instantiation
basic_obj = BasicClass.new("Test")
basic_name = basic_obj.get_name
basic_obj.set_name("New Name")
class_method_result = BasicClass.class_method

child_obj = ChildClass.new("Child", 25)
child_info = child_obj.get_info

# Module usage
module_obj = ClassWithModule.new
module_result = module_obj.test_include

# Person class usage
person = Person.new("Alice", 30)
greeting = person.greet
person.name = "Alicia"
person.age = 31

# Counter class usage
counter1 = Counter.new
counter2 = Counter.new
total_count = Counter.count

# Visibility example
vis_obj = VisibilityExample.new
public_result = vis_obj.public_method
private_result = vis_obj.call_private

# Self example
self_class = SelfExample.class_method
self_obj = SelfExample.new
self_instance = self_obj.instance_method

# Global variable access
access_global

# Nested class usage
outer = Outer.new
inner = outer.create_inner
inner_result = inner.inner_method

# Constants access
version = Constants::VERSION
max_size = Constants::MAX_SIZE

# Proc usage (if supported)
proc_result = simple_proc.call(5)

puts "All tests completed successfully"
