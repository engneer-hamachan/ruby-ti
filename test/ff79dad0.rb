def test
  begin
    raise RuntimeError
  rescue => e
    return 1
  end

  raise RuntimeError
end

dbtp test
