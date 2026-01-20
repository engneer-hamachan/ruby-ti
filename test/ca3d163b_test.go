package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestCa3d163b(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./ca3d163b.rb", "--suggest", "--row=1")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `%!=:::!=(untyped) -> Bool:::Perform inequality comparison
%->:::Kernel.->(optional untyped) <block_params: untyped> -> Proc:::
%<:::Hash.<(Hash) -> Bool:::
%<=:::Hash.<=(Hash) -> Bool:::
%<=>:::<=>(untyped) -> Union<Integer NilClass>:::Comparison operator
%==:::==(untyped) -> Bool:::Perform equality comparison
%==:::Hash.==(untyped) -> Bool:::Perform equality comparison
%===:::===(untyped) -> Bool:::Judge case equality
%>:::Hash.>(Hash) -> Bool:::
%>=:::Hash.>=(Hash) -> Bool:::
%[]:::Hash.[](untyped) -> Self:::Get value for specified key
%[]=:::Hash.[]=(untyped) -> Self:::Set value for specified key
%attr_accessor:::attr_accessor(untyped) -> NilClass:::
%attr_reader:::attr_reader(untyped) -> NilClass:::
%block_given?:::block_given?() -> Bool:::
%class:::class() -> String:::Return object class
%clear:::Hash.clear() -> Hash:::Remove all elements from hash
%collect:::Enumerable.collect() <block_params: Item> -> BlockResultArray:::Execute block for each element and return results as array
%delete:::Hash.delete(Union<String Symbol>) -> Union<Unify NilClass>:::Delete specified key
%dup:::dup() -> Self:::Duplicate object
%each:::Hash.each() <block_params: untyped, Unify> -> Self:::Execute block for each key and value
%each_with_index:::Enumerable.each_with_index() <block_params: Unify, Integer> -> Self:::Execute block for each element and index
%empty?:::Hash.empty?() -> Bool:::Check if hash is empty
%exit:::exit(optional Integer) -> NilClass:::
%has_key?:::Hash.has_key?(Union<String Symbol>) -> Bool:::Check if specified key exists
%has_value?:::Hash.has_value?(untyped) -> Bool:::Check if specified value exists
%include:::include(untyped) -> NilClass:::
%inspect:::inspect() -> String:::Return string representation for debugging
%inspect:::Hash.inspect() -> String:::Return string representation for debugging
%is_a?:::is_a?(untyped) -> Bool:::Check if instance of specified class or module
%key:::Hash.key(untyped) -> Union<String Symbol NilClass>:::
%keys:::Hash.keys() -> Array<untyped>:::Return all keys as array
%lambda:::Kernel.lambda() <block_params: untyped> -> Proc:::Create lambda expression
%length:::Hash.length() -> Bool:::Return number of elements in hash
%loop:::loop() <block_params: NilClass> -> NilClass:::
%merge:::Hash.merge(Hash) <block_params: Symbol, Unify, UnifyArgument> -> Self:::Return new hash merged with another hash
%merge!:::Hash.merge!(Hash) <block_params: Symbol, Unify, UnifyArgument> -> Self:::Merge hash destructively
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
%shift:::Hash.shift() -> Array<untyped>:::Remove and return first key-value pair
%sleep:::sleep(Union<Integer Float>) -> Integer:::
%sleep_ms:::sleep_ms(Integer) -> Integer:::
%sprintf:::sprintf(String, untyped) -> String:::
%system:::Kernel.system(String) -> Bool:::
%to_h:::Hash.to_h() -> Self:::Return self as hash
%to_s:::to_s() -> String:::Return string representation
%values:::Hash.values() -> KeyValueArray:::Return all values as array
%yield:::Kernel.yield(untyped) -> SelfArgument:::`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
