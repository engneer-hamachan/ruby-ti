package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test6ad663e7(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./6ad663e7.rb", "--suggest", "--row=10")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `%!=:::!=(untyped) -> Bool:::Perform inequality comparison
%->:::Kernel.->(optional untyped) <block_params: untyped> -> Proc:::
%<=>:::<=>(untyped) -> Union<Integer NilClass>:::Comparison operator
%==:::==(untyped) -> Bool:::Perform equality comparison
%===:::===(untyped) -> Bool:::Judge case equality
%attr_accessor:::attr_accessor(untyped) -> NilClass:::
%attr_reader:::attr_reader(untyped) -> NilClass:::
%block_given?:::block_given?() -> Bool:::
%class:::class() -> String:::Return object class
%dup:::dup() -> Self:::Duplicate object
%exit:::exit(optional Integer) -> NilClass:::
%extend:::extend(untyped) -> NilClass:::
%include:::include(untyped) -> NilClass:::
%inspect:::inspect() -> String:::Return string representation for debugging
%is_a?:::is_a?(untyped) -> Bool:::Check if instance of specified class or module
%lambda:::Kernel.lambda() <block_params: untyped> -> Proc:::Create lambda expression
%loop:::loop() <block_params: NilClass> -> NilClass:::
%methods:::methods() -> Array<Symbol>:::get object method list
%nil?:::nil?() -> Bool:::Check if nil
%p:::Kernel.p(untyped) -> SelfArgument:::Inspect and output object
%print:::Kernel.print(optional untyped) -> NilClass:::Output object
%private:::private(optional untyped) -> NilClass:::
%proc:::Kernel.proc() <block_params: untyped> -> Proc:::Create Proc object
%public:::public(optional untyped) -> NilClass:::
%puts:::Kernel.puts(untyped) -> NilClass:::Output object with newline
%raise:::raise(optional untyped, optional String) -> Bot:::
%relinquish:::relinquish() -> Bool:::
%sleep:::sleep(Union<Integer Float>) -> Integer:::
%sleep_ms:::sleep_ms(Integer) -> Integer:::
%sprintf:::sprintf(String, untyped) -> String:::
%system:::Kernel.system(String) -> Bool:::
%test:::User.test() -> unknown:::
%to_s:::to_s() -> String:::Return string representation
%yield:::Kernel.yield(untyped) -> SelfArgument:::`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
