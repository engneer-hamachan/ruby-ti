# Simple test for compact type notation

class Test
  # Test Union type: String|Int
  def test_union
    t = Test.new
    result = t.compact_union("hello")  # OK: String
    result = t.compact_union(123)      # OK: Int
    result = t.compact_union(1.5)      # ERROR: Float
  end

  # Test Optional return: ?[String] means Union<Array<String>, Nil>
  def test_optional_return
    t = Test.new
    result = t.compact_complex(["hello"])
    # result can be nil or array
    if result
      puts "got array"
    else
      puts "got nil"
    end
  end
end
