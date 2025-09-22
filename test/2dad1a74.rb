def with_block
  yield if block_given?
end

def with_block
  yield 1 if block_given?
end

class Person
  def initialize name, age
    @name = name
    @age = age
  end
end

class Student < Person
end

s = Student.new 'hama', 100

dbtp s
