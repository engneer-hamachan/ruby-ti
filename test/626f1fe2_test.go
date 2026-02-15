package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test626f1fe2(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./626f1fe2.rb", "-i")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `@./626f1fe2.rb:::4:::bind: Array<String>
@./626f1fe2.rb:::3:::(String) -> Hash [c/public]
@./626f1fe2.rb:::11:::(String) -> Bool [c/private]
@./626f1fe2.rb:::15:::(String) -> Bool [c/private]
@./626f1fe2.rb:::19:::(Union<untyped String>) -> Bool [c/private]
@./626f1fe2.rb:::23:::(String) -> Bool [c/private]
@./626f1fe2.rb:::27:::(String) -> Bool [c/private]
@./626f1fe2.rb:::31:::(unknown) -> Bool [c/private]
@./626f1fe2.rb:::35:::(unknown) -> Bool [c/private]
@./626f1fe2.rb:::41:::bind: Array<String>
@./626f1fe2.rb:::39:::() -> Array<String> [c/private]
@./626f1fe2.rb:::45:::(Array<untyped>) -> Integer [c/private]
@./626f1fe2.rb:::49:::(Array<untyped>) -> String [c/private]
@./626f1fe2.rb:::54:::bind: Array<untyped>
@./626f1fe2.rb:::58:::bind: Array<String>
@./626f1fe2.rb:::53:::() -> Union<Integer Array<String>> [c/private]
@./626f1fe2.rb:::68:::bind: Array<untyped>
@./626f1fe2.rb:::72:::bind: Array<String>
@./626f1fe2.rb:::66:::() -> Union<String Array<String>> [c/private]
@./626f1fe2.rb:::83:::bind: Union<String Array<String>>
@./626f1fe2.rb:::86:::bind: Symbol
@./626f1fe2.rb:::80:::() -> Symbol [c/private]
@./626f1fe2.rb:::96:::bind: Union<NilClass Integer Array<String> String Hash>
@./626f1fe2.rb:::94:::() -> Union<NilClass Integer Array<String> String Hash> [c/private]
@./626f1fe2.rb:::110:::bind: Hash
@./626f1fe2.rb:::114:::bind: Symbol
@./626f1fe2.rb:::116:::bind: Union<NilClass Integer Array<String> String Hash>
@./626f1fe2.rb:::109:::() -> Hash [c/private]
@./626f1fe2.rb:::129:::bind: Hash
./626f1fe2.rb:::130:::Hash`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
