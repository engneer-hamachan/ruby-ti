class User < ApplicationRecord
  def self.test name
    dbtp name
    where(name: name).all
  end

  def self.test2 name
    dbtp name
    self.where(name: name).all
  end
end

User.test 'hoge'
User.test2 1
