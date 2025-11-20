class User < ApplicationRecord
  scope :todos, ->(value) {}

  def self.get_all_by_name name:
    where(name:).all
  end
end

users = User.get_all_by_name(name: 'hoge')
p users.first

