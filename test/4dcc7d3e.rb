class Person
  def initialize(name, age=0)
    dbtp age
    @name = name
    @age = age
  end
end


person = Person.new("Alice", [])
