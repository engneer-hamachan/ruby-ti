module Fuga
  def f x
    p x
  end
end

class Piyo
  include Fuga

  def y
    f 1
  end
end

piyo = Piyo.new
piyo
