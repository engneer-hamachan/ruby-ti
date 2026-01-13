a = {name: 'hamachang', age: 1, hoge: 'fuga'}

case a
in {name:, age:, **y}
  dbtp name
  dbtp age
  dbtp y
end
