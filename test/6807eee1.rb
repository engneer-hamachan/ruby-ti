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

  def test3
    1
  end
end

m = MatzSanArigato.new

m.test3 + '1'

MatzSanArigato.test + 1

