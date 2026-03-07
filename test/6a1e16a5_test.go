package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test6a1e16a5(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./6a1e16a5.rb")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `./6a1e16a5.rb:::3:::String
./6a1e16a5.rb:::7:::String
./6a1e16a5.rb:::11:::String
./6a1e16a5.rb:::15:::Integer
./6a1e16a5.rb:::19:::Integer
./6a1e16a5.rb:::23:::String
./6a1e16a5.rb:::27:::String
./6a1e16a5.rb:::31:::Integer
./6a1e16a5.rb:::35:::Integer
./6a1e16a5.rb:::39:::String
./6a1e16a5.rb:::46:::Integer
./6a1e16a5.rb:::50:::Integer
./6a1e16a5.rb:::54:::Integer
./6a1e16a5.rb:::58:::Integer
./6a1e16a5.rb:::62:::String
./6a1e16a5.rb:::66:::String`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
