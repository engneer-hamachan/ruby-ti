go build .
mkdir -p ./bin
mv ./ti ./bin

mkdir -p .ti-config
cp test/.ti-config/array.json .ti-config/array.json
cp test/.ti-config/bool.json .ti-config/bool.json
cp test/.ti-config/class.json .ti-config/class.json
cp test/.ti-config/enumerable.json .ti-config/enumerable.json
cp test/.ti-config/float.json .ti-config/float.json
cp test/.ti-config/gpio.json .ti-config/gpio.json
cp test/.ti-config/hash.json .ti-config/hash.json
cp test/.ti-config/identifier.json .ti-config/identifier.json
cp test/.ti-config/integer.json .ti-config/integer.json
cp test/.ti-config/kernel.json .ti-config/kernel.json
cp test/.ti-config/math.json .ti-config/math.json
cp test/.ti-config/nil.json .ti-config/nil.json
cp test/.ti-config/object.json .ti-config/object.json
cp test/.ti-config/proc.json .ti-config/proc.json
cp test/.ti-config/range.json .ti-config/range.json
cp test/.ti-config/runtime_error.json .ti-config/runtime_error.json
cp test/.ti-config/string.json .ti-config/string.json
cp test/.ti-config/symbol.json .ti-config/symbol.json
cp test/.ti-config/untyped.json .ti-config/untyped.json

go build -o ./bin/ti-c2json ./cmd/cpp2json/main.go
