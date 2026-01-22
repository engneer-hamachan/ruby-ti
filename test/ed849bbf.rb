module Fugaa
  def hoge x
    p x
    dbtp x
  end
end

class User
  extend Fugaa

  def piyo
    hoge 1
  end
end

User.hoge 1

