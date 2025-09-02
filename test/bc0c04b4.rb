config = {db: {user: 'admin', password: 'abc123'}}

case config
in db: {user:} # ネストしてハッシュにマッチして、その値を変数userに代入する
  puts "Connect with user '#{user}'"
in connection: {username: }
  puts "Connect with user '#{username}'"
else
  puts "Unrecognized structure of config"
end
