class User
  def test
    1
  end
end

u = User.new

a = {a: u.test ? 1.5 : 2}

dbtp a[:a]
