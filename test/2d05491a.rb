class Hoge
  attr_accessor :code

  def initialize
    @code = ''
  end
end

h = Hoge.new

h.code = h.code + '1'

if h.code == '1'
end
