class Hoge
  def fuga
    1
  end
end

a = []
h = Hoge.new

a.append(h)

a.each do |x|
  x.fuga
  x.piyo
end
