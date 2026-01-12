a = 
  case variable
  in String => x
    p x
  in Integer => x
    p x
  end

p x


a = 
  case arr
  in [x, y, u, String => o]
    p x
    p y
    p u
    p o
  end

a = 
  case 1
  in r, z
    p r
    p z
  else 
    nil
  end


a = 
  case 1
  in {name: names, age: ages}
    p names
    p name
    p age
    p ages
  in Integer => x
    p x
  end
