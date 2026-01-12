a = 
  case variable
  in String => x
    dbtp x
  in Integer => x
    dbtp x
  end

p x


a = 
  case arr
  in [x, y, u, String => o]
    dbtp x
    dbtp y
    dbtp u
    dbtp o
    dbtp j
  end

a = 
  case 1
  in r, z
    dbtp r
    dbtp z
  else 
    nil
  end


a = 
  case 1
  in {name: names, age: ages}
    dbtp names
    dbtp name
    dbtp age
    dbtp ages
  in Integer => x
    dbtp x
  end
