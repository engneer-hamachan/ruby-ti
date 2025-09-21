class Hoge
  def initialize(name, age)
    @name = name
    @age = age
  end

  def name
    @name
  end

  def age
    @age
  end
end

h = Hoge.new 'abc', 1



dbtp h.name
