def hoge(&block)
  block.call if 1
end

class Person
end

class Student < Person
  def initialize(x, y)
  end
end

s = Student.new '1', {}

dbtp s

s = Student.new
