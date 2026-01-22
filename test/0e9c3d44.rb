module Fugaa
  def yyy x
    1
  end
end

class User
  include Fugaa
end


dbtp User.yyy 1

