hash = {}
a, b = 1.5, nil

hash[:x] = 1 % 2 ? a : b
dbtp hash[:x]

hash[:x] = 1, 2, 3
dbtp hash[:x]

