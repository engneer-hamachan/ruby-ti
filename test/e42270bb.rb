class User < ActiveRecord::Base
  scope :a, -> {}
  scope 
  has_one 

  validate :test


  private

  def test
  end
end
