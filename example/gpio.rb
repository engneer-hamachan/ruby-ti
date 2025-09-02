loop do
  pin = GPIO.new(19, GPIO::OUT)
  sleep 1
  pin.write 1
  sleep 1
  pin.write 0
end
