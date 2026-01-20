module Fugaa
  def hoge
    1
  end
end

class User
  include Fugaa

  def piyo
    hoge
  end
end


u = User
u
