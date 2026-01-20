package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test0ab8abe9(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./0ab8abe9.rb", "--suggest", "--row=1")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `%!:::String.!() -> Bool:::Return logical NOT
%!=:::!=(untyped) -> Bool:::Perform inequality comparison
%*:::String.*(Integer) -> String:::Repeat string specified number of times
%+:::String.+(String) -> String:::Concatenate strings
%->:::Kernel.->(optional untyped) <block_params: untyped> -> Proc:::
%<:::String.<(String) -> Bool:::Check if lexicographically smaller
%<<:::String.<<(Union<String Integer>) -> String:::Append string or character code
%<=:::String.<=(String) -> Bool:::Check if lexicographically less than or equal
%<=>:::<=>(untyped) -> Union<Integer NilClass>:::Comparison operator
%<=>:::String.<=>(String) -> Integer:::Compare lexicographically. Returns -1 if self < other, 0 if self == other, 1 if self > other
%==:::==(untyped) -> Bool:::Perform equality comparison
%==:::String.==(untyped) -> Bool:::Perform equality comparison
%===:::===(untyped) -> Bool:::Judge case equality
%===:::String.===(untyped) -> Bool:::Perform equality comparison (for case statements)
%>:::String.>(String) -> Bool:::Check if lexicographically greater
%>=:::String.>=(String) -> Bool:::Check if lexicographically greater than or equal
%[]:::String.[](untyped) -> Union<String NilClass>:::Get character or substring at specified position
%[]=:::String.[]=(untyped, String) -> String:::Replace character or substring at specified position
%__ljust_rjust_argcheck:::String.__ljust_rjust_argcheck(Union<Integer String>) -> NilClass:::Internal method for argument checking of ljust and rjust
%attr_accessor:::attr_accessor(untyped) -> NilClass:::
%attr_reader:::attr_reader(untyped) -> NilClass:::
%block_given?:::block_given?() -> Bool:::
%bytes:::String.bytes() -> Array<Integer>:::Return byte sequence of string as array
%chars:::String.chars() -> Array<String>:::Return characters of string as array
%chomp:::String.chomp(optional String) -> String:::Return string with trailing newline removed
%chomp!:::String.chomp!(optional String) -> Union<String NilClass>:::Remove trailing newline destructively
%class:::class() -> String:::Return object class
%clear:::String.clear() -> String:::Clear string contents
%downcase:::String.downcase(optional untyped, optional untyped) -> String:::Return string converted to lowercase
%downcase!:::String.downcase!(optional untyped, optional untyped) -> Union<String NilClass>:::Convert to lowercase destructively
%dup:::dup() -> Self:::Duplicate object
%each_byte:::String.each_byte() <block_params: Integer> -> Self:::Execute block for each byte
%each_char:::String.each_char() <block_params: String> -> Self:::Execute block for each character
%each_line:::String.each_line() <block_params: String> -> String:::Execute block for each line
%empty?:::String.empty?() -> Bool:::Check if string is empty
%end_with?:::String.end_with?(untyped) -> String:::Check if ends with specified string
%exit:::exit(optional Integer) -> NilClass:::
%getbyte:::String.getbyte(Integer) -> Union<Integer NilClass>:::Get byte value at specified position
%include:::include(untyped) -> NilClass:::
%include?:::String.include?(String) -> Bool:::Check if contains specified string
%index:::String.index(Union<String Integer>) -> Union<Integer NilClass>:::Return first position of specified string
%inspect:::inspect() -> String:::Return string representation for debugging
%inspect:::String.inspect() -> String:::Return string representation for debugging
%intern:::String.intern() -> Symbol:::Convert to symbol
%is_a?:::is_a?(untyped) -> Bool:::Check if instance of specified class or module
%lambda:::Kernel.lambda() <block_params: untyped> -> Proc:::Create lambda expression
%length:::String.length() -> Integer:::Return string length
%ljust:::String.ljust(Integer, optional String) -> String:::Return string left-justified to specified width
%loop:::loop() <block_params: NilClass> -> NilClass:::
%lstrip:::String.lstrip() -> String:::Return string with leading whitespace removed
%lstrip!:::String.lstrip!() -> Union<String NilClass>:::Remove leading whitespace destructively
%methods:::methods() -> Array<Symbol>:::get object method list
%nil?:::nil?() -> Bool:::Check if nil
%ord:::String.ord() -> Integer:::Return character code of first character
%p:::Kernel.p(untyped) -> SelfArgument:::Inspect and output object
%print:::Kernel.print(optional untyped) -> NilClass:::Output object
%private:::private(optional untyped) -> NilClass:::
%proc:::Kernel.proc() <block_params: untyped> -> Proc:::Create Proc object
%public:::public(optional untyped) -> NilClass:::
%puts:::Kernel.puts(untyped) -> NilClass:::Output object with newline
%raise:::raise(optional untyped, optional String) -> Bot:::
%relinquish:::relinquish() -> Bool:::
%rjust:::String.rjust(Integer, optional String) -> String:::Return string right-justified to specified width
%rstrip:::String.rstrip() -> String:::Return string with trailing whitespace removed
%rstrip!:::String.rstrip!() -> Union<String NilClass>:::Remove trailing whitespace destructively
%sleep:::sleep(Union<Integer Float>) -> Integer:::
%sleep_ms:::sleep_ms(Integer) -> Integer:::
%split:::String.split(optional String, optional Integer) -> Array<String>:::Split by delimiter and return array
%sprintf:::sprintf(String, untyped) -> String:::
%start_with?:::String.start_with?(untyped) -> Bool:::Check if starts with specified string
%strip:::String.strip() -> String:::Return string with leading and trailing whitespace removed
%strip!:::String.strip!() -> Union<String NilClass>:::Remove leading and trailing whitespace destructively
%system:::Kernel.system(String) -> Bool:::
%to_f:::String.to_f() -> Float:::Convert to floating point number
%to_i:::String.to_i(optional Integer) -> Integer:::Convert to integer
%to_s:::to_s() -> String:::Return string representation
%to_s:::String.to_s() -> String:::Return self as string
%to_sym:::String.to_sym() -> Symbol:::Convert to symbol
%tr:::String.tr(String, String) -> String:::Return string with characters replaced
%tr!:::String.tr!(String, String) -> Union<String NilClass>:::Replace characters destructively
%upcase:::String.upcase(optional untyped, optional untyped) -> String:::Return string converted to uppercase
%upcase!:::String.upcase!(optional untyped, optional untyped) -> Union<String NilClass>:::Convert to uppercase destructively
%yield:::Kernel.yield(untyped) -> SelfArgument:::`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
