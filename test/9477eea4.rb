module Fugaa
  def yyy x
    1
  end
end

class User
  include Fugaa
end

u = User.new

dbtp u.yyy 1

