module Fugaa
  def self.hoge
    1
  end
end

class User
  include Fugaa

  def piyo
    p hoge
  end
end


