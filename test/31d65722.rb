module Fuga
  def f x
    p x
    dbtp x
  end
end

class Piyo
  extend Fuga

  def y
    f 1
  end
end

Piyo.f 1


class Piya
  include Fuga

  def z
    f '1'
  end
end

piya = Piya.new
piya.f '1'
piya.z
