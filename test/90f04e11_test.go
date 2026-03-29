package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test90f04e11(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./90f04e11.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./90f04e11.rb:::2:::untyped
./90f04e11.rb:::3:::Dir
./90f04e11.rb:::12:::Array<String>
./90f04e11.rb:::13:::Array<String>
./90f04e11.rb:::14:::Array<String>
./90f04e11.rb:::15:::Array<String>
./90f04e11.rb:::17:::type mismatch: expected Union<String NilClass>, but got Integer for Dir.glob
./90f04e11.rb:::17:::KeyValue
./90f04e11.rb:::18:::type mismatch: expected Integer, but got String for Dir.glob
./90f04e11.rb:::18:::String`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
