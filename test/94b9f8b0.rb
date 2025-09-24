class Animal
  attr_accessor :name, :age

  def initialize(name, age = 0)
    @name = name
    @age = age
    @species = self.class.name
  end

  def speak
    "#{@name} makes a sound"
  end

  def self.total_animals
    @@count ||= 0
  end

  def test
    self
  end

  private

  def secret_method
    "This is private"
  end
end

a = Animal.new 'hoge'

dbtp a.test.class.name
