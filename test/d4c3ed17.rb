module Fuga
  def hoge
    1
  end
end

class User
  include Fuga

  def piyo
    hoge
  end
end

u = User.new
u
