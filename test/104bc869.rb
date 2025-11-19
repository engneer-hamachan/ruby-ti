class User < ApplicationRecord
  has_many :todos, dependent: :destroy

  def self.test
    self.where(name: 'hoge').all
  end
end

user = User.test

dbtp user
