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

# instance method: chain
a8 = 'Hello World'.upcase.length
dbtp a8

# instance method: with operator
a9 = 'abc'.length * 2 + 1
dbtp a9

# instance method: ternary on result
a10 = 'abc'.length > 2 ? 'long' : 'short'
dbtp a10

# args: arithmetic in arguments
def test_add(x, y)
  x + y
end
a11 = test_add(1 + 2, 3 * 4)
dbtp a11

# args: method call in arguments
a12 = test_add('hello'.length, 'world'.length)
dbtp a12

# args: nested parentheses in arguments
a13 = test_add(10 - (3.abs), 2 + (1.abs))
dbtp a13

# args: ternary in arguments
a14 = test_add(true ? 1 : 2, false ? 3 : 4)
dbtp a14

# ternary: instance method in branches
a15 = true ? 'hello'.upcase : 'world'.downcase
dbtp a15

# ternary: comparison + method chain
a16 = 'abc'.length == 'def'.length ? 'same' : 'diff'
dbtp a16
