class User
  def self.hoge
    1
  end

  def self.piyo
    hoge
  end
end

p User.piyo
