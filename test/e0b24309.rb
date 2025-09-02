def test
  begin
    raise RuntimeError
  rescue RuntimeError => e
    dbtp e
  end
end

test
