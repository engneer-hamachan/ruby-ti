package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test4355a46b(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./4355a46b.rb", "--suggest", "--row=1")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `%%:::Integer.%(Match) => (Integer) -> Integer, (Float) -> Float:::Calculate modulo
%&:::Integer.&(Integer) -> Integer:::Perform bitwise AND operation
%*:::Integer.*(Match) => (Integer) -> Integer, (Float) -> Float:::Perform multiplication
%**:::Integer.**(Union<Integer Float>) -> Integer:::Calculate power
%+:::Integer.+(Match) => (Integer) -> Integer, (Float) -> Float:::Perform addition
%+@:::Integer.+@() -> Integer:::Unary plus operator. Returns self
%-:::Integer.-(Match) => (Integer) -> Integer, (Float) -> Float:::Perform subtraction
%-@:::Integer.-@() -> Integer:::Unary minus operator. Inverts sign
%/:::Integer./(Match) => (Integer) -> Integer, (Float) -> Float:::Perform division
%<:::Integer.<(Union<Integer Float>) -> Bool:::Perform less than comparison
%<<:::Integer.<<(Integer) -> Integer:::Perform left bit shift
%<=:::Integer.<=(Union<Integer Float>) -> Bool:::Perform less than or equal comparison
%<=>:::Integer.<=>(Match) => (Integer) -> Integer, (untyped) -> Union<Integer NilClass>:::Comparison operator. Returns -1 if self < other, 0 if self == other, 1 if self > other
%==:::Integer.==(untyped) -> Bool:::Perform equality comparison
%>:::Integer.>(Union<Integer Float>) -> Bool:::Perform greater than comparison
%>=:::Integer.>=(Union<Integer Float>) -> Bool:::Perform greater than or equal comparison
%>>:::Integer.>>(Integer) -> Integer:::Perform right bit shift
%[]:::Integer.[](untyped) -> Integer:::Get value at specified bit position
%^:::Integer.^(Integer) -> Integer:::Perform bitwise XOR operation
%abs:::Integer.abs() -> Integer:::Get absolute value
%chr:::Integer.chr() -> String:::Return single character string corresponding to character code
%downto:::Integer.downto(Integer) <block_params: Integer> -> Integer:::Execute block while decreasing from self to argument by 1
%times:::Integer.times() <block_params: Integer> -> Integer:::Execute block self times
%to_f:::Integer.to_f() -> Float:::Convert to floating point number
%to_i:::Integer.to_i() -> Integer:::Return self as integer value
%to_s:::Integer.to_s(optional Integer) -> String:::Convert to string. Base can be specified with argument
%|:::Integer.|(Integer) -> Integer:::Perform bitwise OR operation
%~:::Integer.~() -> Integer:::Perform bitwise NOT operation`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
