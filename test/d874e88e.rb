config = {db: {user: 'admin', password: 1}}
 
case config
in db: {password:} 
  dbtp password
in connection: {username: }
  dbtp username
else
  dbtp "hoge"
end 

