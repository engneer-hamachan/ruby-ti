class User < ActiveRecord::Base
  scope :a, -> {}
  validate :test
  has_many :todos
  

  private

  def test
  end
end

class Fuga < ActiveRecord::Base
  scope :a, -> {}
  validate :test
  has_many :todos
  

  private

  def test
  end
end

user = User.find 1
users = User.find [1, 2, 3]

ans = User.where(id: 1)


users = User.all

dbtp User.where(name: 'hoge', dleted_at: nil).all

dbtp Fuga.where(name: 'hoge', dleted_at: nil).all

User.where(a: 1)
