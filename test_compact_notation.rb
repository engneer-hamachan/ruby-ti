# Test compact type notation in builtin JSON

class Test
  # Test 1: Optional type (?String) should accept String or Nil
  def test_optional_string
    t = Test.new
    result = t.compact_optional("hello")  # Should work: String
    result = t.compact_optional()      # Should work: Nil
  end

  # Test 2: Asterisk type (*String) should accept multiple strings
  def test_asterisk
    t = Test.new
    result = t.compact_asterisk("a", "b", "c", 1, 2)  # Should work: *String
  end

  # Test 3: Union type (String|Int) should accept String or Int
  def test_union
    t = Test.new
    result = t.compact_union("hello")  # Should work: String
    result = t.compact_union(123)      # Should work: Int
    result = t.compact_union(1.5)      # Should fail: Float
  end

  # Test 4: Array type ([String]) should accept array of strings
  def test_array
    t = Test.new
    result1 = t.compact_array(["a", "b"])  # Should work: [String]
    result2 = t.compact_array([1, 2])      # Should fail: [Int]
  end

  # Test 5: Complex type ([String|Int]) with optional return (?[String])
  def test_complex
    t = Test.new
    arr = ["hello", 123]
    result = t.compact_complex(arr)  # Should work: [String|Int]
    # result is ?[String] so can be nil or array
    if result
      x = result[0]  # Should be String
    end
  end
end
