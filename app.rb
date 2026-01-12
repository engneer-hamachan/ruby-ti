a = 
  case variable
  in String => x
    p x
  in Integer => x
    p x
  end


a = 
  case arr
  in [x, y]
    p x
    p y
  end

a = 
  case 1
  in r
    p r
  else 
    nil
  end


a = 
  case 1
  in {name:, age: ages}
    p name
    p age
    p ages
  in Integer => x
    p x
  end
