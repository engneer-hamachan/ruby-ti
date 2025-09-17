class Hoge
  class << self
    private

    def test2
      1
    end

    public

    def test
      1
    end
  end
end

dbtp Hoge.test2
