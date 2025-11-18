class User < ActiveRecord::Base
  scope :a, -> {}
  validate :test
  has_many :todos
  

  private

  def test
  end
end


dbtp User.where(name: 'hoge', dleted_at: nil)
dbtp User.where(name: 'hoge', dleted_at: nil).first

dbtp User.where(name: 'hoge', dleted_at: nil).all
dbtp User.where(name: 'hoge', dleted_at: nil).all.first

