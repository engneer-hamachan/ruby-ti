config = {db: {user: 'admin', password: 1}}
 
case config
in db: {user:} 
  dbtp user
in connection: {username: }
  dbtp username
else
  dbtp "hoge"
end 

