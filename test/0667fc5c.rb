module Fugaa
  def hoge
    1
  end
end

class User
  extend Fugaa

  def piyo
    p hoge
  end
end

u = User.new
u.piyo
