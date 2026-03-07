# ternary: simple
a1 = true ? 'yes' : 'no'
dbtp a1

# ternary: comparison + ternary
a2 = :load == :save ? 'Save' : 'Load'
dbtp a2

# ternary: arithmetic in condition
a3 = 1 + 2 == 3 ? 'match' : 'no'
dbtp a3

# arithmetic: precedence * before +
a4 = 2 * 3 + 10
dbtp a4

# arithmetic: method chain + operator
a5 = 'hello'.length + 1
dbtp a5

# ternary: method chain in condition
a6 = 'abc'.length == 3 ? 'three' : 'other'
dbtp a6

# ternary: nested
a7 = true ? (false ? 'a' : 'b') : 'c'
dbtp a7
