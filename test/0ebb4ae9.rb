module Fugaa
  def yyy x
    1
  end
end

class User
  extend Fugaa
end


dbtp User.yyy 1

