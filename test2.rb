# mruby 構文テストファイル

# 基本的な変数代入
x = 1
y = 2.5
z = "hello"
flag = true
empty = nil

# 配列とハッシュ
arr = [1, 2, 3, "string", nil]
hash = {a: 1, b: 2, "key" => "value"}
nested = [[1, 2], {x: [3, 4]}]


# 文字列操作
str1 = "Hello"
str2 = 'World'
interpolated = "#{str1} #{str2}!"

dbtp 1


# 数値演算
sum = x + y
diff = x - y
product = x * y
division = y / x
modulo = 10 % 3
power = 2 ** 3


# 比較演算
equal = (x == 1)
not_equal = (x != 2)
less_than = (x < y)
greater_than = (y > x)
less_equal = (x <= 1)
greater_equal = (y >= 2.5)


# 論理演算
and_op = flag && true
or_op = flag || false
not_op = !flag


# 条件分岐
if x > 0
  puts "positive"
elsif x < 0
  puts "negative"
else
  puts "zero"
end

# case文
case x
when 1
  puts "one"
when 2, 3
  puts "two or three"
when 4..10
  puts "four to ten"
else
  puts "other"
end

# 三項演算子
result = x > 0 ? "positive" : "not positive"

# ループ
# while ループ
i = 0
while i < 3
  puts i
  i += 1
end

# until ループ
j = 0
until j >= 3
  puts j
  j += 1
end

# for ループ
for k in [1, 2, 3]
  puts k
end

# times
3.times do |n|
  puts n
end

# each
[1, 2, 3].each do |item|
  puts item
end

# each_with_index
["a", "b", "c"].each_with_index do |char, idx|
  puts "#{idx}: #{char}"
end

# メソッド定義
def simple_method
  "simple"
end

def method_with_args(a, b)
  a + b
end

def method_with_default(x, y = 10)
  x * y
end

def method_with_block
  yield if block_given?
end


def method_with_splat(*args)
  args.length
end

# クラス定義
class Person
  attr_reader :name
  attr_writer :age
  attr_accessor :email

  def initialize(name, age = 0)
    @name = name
    @age = age
    @email = nil
  end

  def greet
    "Hello, I'm #{@name}"
  end

  def self.species
    "Homo sapiens"
  end

  private

  def private_method
    "private"
  end

  protected

  def protected_method
    "protected"
  end
end

# クラスの継承
class Student < Person
  def initialize(name, age, school)
    super(name, age)
    @school = school
  end

  def study
    "#{@name} is studying at #{@school}"
  end
end

# モジュール
module Greetings
  def hello
    "Hello from module"
  end

  def self.farewell
    "Goodbye"
  end
end

# モジュールのインクルード
class Teacher
  include Greetings

  def teach
    "Teaching..."
  end
end

# 定数
PI = 3.14159
MAX_SIZE = 100

# クラス変数
class Counter
  @@count = 0

  def increment
    @@count += 1
  end

  def self.total
    @@count
  end
end

# 例外処理
begin
  result = 10 / 0
rescue ZeroDivisionError => e
  puts "Division by zero: #{e.message}"
rescue StandardError => e
  puts "Error: #{e.message}"
ensure
  puts "Always executed"
end

# ブロックの様々な形
# do-end形式
[1, 2, 3].each do |x|
  puts x * 2
end

# {}形式
[1, 2, 3].map { |x| x * 3 }

# proc
my_proc = Proc.new { |x| x + 1 }
result = my_proc.call(5)

# lambda
my_lambda = lambda { |x| x * 2 }
result2 = my_lambda.call(10)

# 正規表現
pattern = /\d+/
text = "abc123def"
match = pattern.match(text)

# 範囲
range1 = 1..10
range2 = 1...10
range3 = 'a'..'z'

# シンボル
sym1 = :symbol
sym2 = :"dynamic symbol"

# 代入演算子
x += 5
y -= 2
z ||= "default"
w &&= true

# 配列とハッシュの操作
arr << 4
arr[0] = 10
hash[:new_key] = "new_value"
hash.delete(:a)

# 文字列の操作
str = "hello world"
upcase = str.upcase
length = str.length
substring = str[0, 5]

# メソッドチェーン
result = [1, 2, 3, 4, 5]
  .select { |x| x.even? }
  .map { |x| x * 2 }
  .reduce(:+)

# 並列代入
a, b, c = [1, 2, 3]
x, *rest = [1, 2, 3, 4, 5]

# ブロック変数の省略記法（mrubyで利用可能な場合）
numbers = [1, 2, 3, 4, 5]
doubled = numbers.map(&:to_s)

# 条件修飾子
puts "positive" if x > 0
puts "not nil" unless empty

# インスタンス変数とクラス変数の操作
person = Person.new("Alice", 25)
puts person.greet
puts Person.species

student = Student.new("Bob", 20, "University")
puts student.study

teacher = Teacher.new
puts teacher.hello

# カウンターのテスト
counter1 = Counter.new
counter2 = Counter.new
counter1.increment
counter2.increment
puts Counter.total

# 複雑な式
complex_result = (x + y) * (z.length + arr.size) / hash.keys.length rescue 0

# ネストしたブロック
matrix = [[1, 2], [3, 4]]
matrix.each do |row|
  row.each do |cell|
    puts cell
  end
end

puts "mruby syntax test completed"
