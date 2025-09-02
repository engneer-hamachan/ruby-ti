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

p MatzSanArigato.test + 1
p MatzSanArigato.test2

