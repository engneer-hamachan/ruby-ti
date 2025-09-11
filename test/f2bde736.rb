class Animal
  def initialize(name)
    @name = name
  end

  def str
    @name
  end
end

class Dog < Animal
end

d = Dog.new('Rex')
dbtp d.str
