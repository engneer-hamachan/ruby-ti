package main

import (
	"os/exec"
	"strings"
	"testing"
)

func Test37517e5f(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./37517e5f.rb", "--suggest", "--row=1")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `%!=:::!=(untyped) -> Bool:::Perform inequality comparison
%&:::Array.&(Array<untyped>) -> Self:::
%*:::Array.*(Match) => (Integer) -> Self, (String) -> String:::Repeat array specified number of times and return new array
%+:::Array.+(Array<untyped>) -> Array<untyped>:::Concatenate arrays
%-:::Array.-(Array<untyped>) -> Self:::Return new array excluding specified elements
%->:::Kernel.->(optional untyped) <block_params: untyped> -> Proc:::
%<=>:::<=>(untyped) -> Union<Integer NilClass>:::Comparison operator
%<=>:::Array.<=>(untyped) -> Union<Integer NilClass>:::Compare arrays. Returns -1 if self < other, 0 if self == other, 1 if self > other
%==:::==(untyped) -> Bool:::Perform equality comparison
%===:::===(untyped) -> Bool:::Judge case equality
%[]:::Array.[](Integer) -> Self:::Get element at specified position
%[]=:::Array.[]=(untyped) -> Self:::Set element at specified position
%all?:::Array.all?() <block_params: Unify> -> Bool:::
%at:::Array.at(Integer) -> OptiionalUnify:::Get element at specified position
%attr_accessor:::attr_accessor(untyped) -> NilClass:::
%attr_reader:::attr_reader(untyped) -> NilClass:::
%block_given?:::block_given?() -> Bool:::
%class:::class() -> String:::Return object class
%clear:::Array.clear() -> Array<untyped>:::Remove all elements from array
%collect:::Enumerable.collect() <block_params: Item> -> BlockResultArray:::Execute block for each element and return results as array
%collect!:::Array.collect!() <block_params: Item> -> BlockResultArray:::Execute block for each element and destructively replace
%concat:::Array.concat(Array<untyped>) -> Array<untyped>:::Concatenate arrays
%count:::Array.count(optional untyped) <block_params: Unify> -> Integer:::
%delete_at:::Array.delete_at(Integer) -> OptiionalUnify:::Delete element at specified position
%delete_if:::Array.delete_if() <block_params: Unify> -> Self:::Delete elements for which block returns true
%dup:::dup() -> Self:::Duplicate object
%each:::Array.each() <block_params: Flatten> -> Self:::Execute block for each element
%each_index:::Array.each_index() <block_params: Integer> -> Self:::Execute block for each index
%each_with_index:::Enumerable.each_with_index() <block_params: Unify, Integer> -> Self:::Execute block for each element and index
%empty?:::Array.empty?() -> Bool:::Check if array is empty
%exit:::exit(optional Integer) -> NilClass:::
%first:::Array.first(optional Integer) -> OptiionalUnify:::Get first element
%flatten:::Array.flatten() -> Self:::Flatten nested array
%include:::include(untyped) -> NilClass:::
%include?:::Array.include?(untyped) -> Bool:::Check if element is included
%inspect:::inspect() -> String:::Return string representation for debugging
%inspect:::Array.inspect() -> String:::Return string representation for debugging
%is_a?:::is_a?(untyped) -> Bool:::Check if instance of specified class or module
%join:::Array.join(optional String) -> String:::Join elements and return string
%lambda:::Kernel.lambda() <block_params: untyped> -> Proc:::Create lambda expression
%last:::Array.last(Match) => (OptiionalUnify) -> OptiionalUnify, (Self) -> Self:::Get last element
%length:::Array.length() -> Integer:::Return number of elements in array
%loop:::loop() <block_params: NilClass> -> NilClass:::
%max:::Array.max(Match) => <block_params: Unify, Unify> (OptiionalUnify) -> OptiionalUnify, (Self) -> Self:::
%methods:::methods() -> Array<Symbol>:::get object method list
%nil?:::nil?() -> Bool:::Check if nil
%p:::Kernel.p(untyped) -> SelfArgument:::Inspect and output object
%pop:::Array.pop(Match) => (OptiionalUnify) -> OptiionalUnify, (Self) -> Self:::Remove and return last element
%print:::Kernel.print(optional untyped) -> NilClass:::Output object
%private:::private(optional untyped) -> NilClass:::
%proc:::Kernel.proc() <block_params: untyped> -> Proc:::Create Proc object
%public:::public(optional untyped) -> NilClass:::
%push:::Array.push(untyped) -> Array<untyped>:::Append element to end
%puts:::Kernel.puts(untyped) -> NilClass:::Output object with newline
%raise:::raise(optional untyped, optional String) -> Bot:::
%reject:::Array.reject() <block_params: Unify> -> Self:::Return array excluding elements for which block returns true
%reject!:::Array.reject!() <block_params: Unify> -> Self:::Destructively delete elements for which block returns true
%relinquish:::relinquish() -> Bool:::
%replace:::Array.replace(Array<untyped>) -> Array<untyped>:::
%shift:::Array.shift(Match) => (Unify) -> Unify, (Self) -> Self:::Remove and return first element
%sleep:::sleep(Union<Integer Float>) -> Integer:::
%sleep_ms:::sleep_ms(Integer) -> Integer:::
%slice:::Array.slice(untyped) -> untyped:::Get subarray
%sort:::Array.sort() <block_params: Unify, Unify> -> Self:::Return sorted array
%sort!:::Array.sort!() <block_params: Unify, Unify> -> Self:::Sort destructively
%sprintf:::sprintf(String, untyped) -> String:::
%system:::Kernel.system(String) -> Bool:::
%to_s:::to_s() -> String:::Return string representation
%uniq:::Array.uniq() -> Self:::Return array with duplicates removed
%uniq!:::Array.uniq!() -> Self:::Remove duplicates destructively
%unshift:::Array.unshift(untyped) -> Array<untyped>:::Prepend element to beginning
%yield:::Kernel.yield(untyped) -> SelfArgument:::
%|:::Array.|(Array<untyped>) -> Array<untyped>:::`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
