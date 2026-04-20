class User
  def [] x
    1
  end

  def []= x, y
    '1'
  end
end

u = User.new
dbtp u[0]

dbtp u[0] = 1
 
