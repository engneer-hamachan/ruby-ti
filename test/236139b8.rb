class User < ApplicationRecord
  def self.test
    where(name: 'hoge').first
  end
end

user = User.test
dbtp user
user.save!


