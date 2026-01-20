module JsonSan
  class << self
    def parse(str)
      @chars = str.chars

      read_object
    end

    private

    def number_char?(char)
      char.ord >= 48 && char.ord <= 57
    end

    def object_char?(char)
      char == '{'
    end

    def space_char?(char)
      char == ' '
    end

    def quote_char?(char)
      char == "'"
    end

    def comma_char?(char)
      char == ','
    end

    def array_start_char?(char)
      char == '['
    end

    def array_end_char?(char)
      char == ']'
    end

    def skip_empty
      @chars.each_with_index do |char, idx|
        return @chars = @chars[idx..] unless space_char? char
      end
    end

    def make_number(chars)
      chars.join.to_i
    end

    def make_str(chars)
      chars.join
    end

    def read_number
      result = []

      @chars.each_with_index do |char, idx|
        unless number_char? char
          @chars = @chars[idx..]
          return make_number result
        end

        result.push char
      end
    end

    def read_string
      @chars.shift
      result = []

      @chars.each_with_index do |char, idx|
        if quote_char? char
          @chars = @chars[(idx + 1)..]
          return make_str result
        end

        result.push char
      end
    end

    def read_key
      skip_empty

      str = read_string

      if str.is_a?(String)
        result = str.to_sym
      end

      skip_empty

      result
    end

    def read_value
      skip_empty
      result =
        if number_char? @chars[0]
          read_number
        elsif quote_char? @chars[0]
          read_string
        elsif object_char? @chars[0]
          read_object
        end
      skip_empty

      result
    end

    def read_object
      result = {}
      @chars.shift

      loop do
        key = read_key

        result[key] = read_value if @chars.shift == ':'

        break unless comma_char? @chars[0]

        @chars.shift
      end

      @chars.shift
      result
    end
  end
end

json = JsonSan.parse "{'abc': 123, 'cde': 'hoge', 'zzz': {'abc': 1}}"
dbtp json
# => {:abc=>123, :cde=>"hoge", :zzz=>{:abc=>1}}
