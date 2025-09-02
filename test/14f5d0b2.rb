class MatzSanArigato
  class << self
    def test
      test2 
    end

    private

    def test2
      1
    end
  end
end

def test
  1
end

def test x
end

p MatzSanArigato.test.to_i.to_s + 1
p MatzSanArigato.test2

