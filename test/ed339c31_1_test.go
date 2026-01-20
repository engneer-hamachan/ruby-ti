package main

import (
	"os/exec"
	"strings"
	"testing"
)

func TestEd339c31_1(t *testing.T) {
	t.Parallel()
	cmd := exec.Command("../ti", "./ed339c31_1.rb", "-i")

	output, _ := cmd.CombinedOutput()

	expectedOutput := `@./ed339c31_1.rb:::4:::bind: Integer
@./ed339c31_1.rb:::5:::bind: String
@./ed339c31_1.rb:::6:::bind: Float
@./ed339c31_1.rb:::7:::bind: Bool
@./ed339c31_1.rb:::8:::bind: NilClass
@./ed339c31_1.rb:::9:::bind: Symbol
@./ed339c31_1.rb:::12:::bind: Integer
@./ed339c31_1.rb:::13:::bind: Integer
@./ed339c31_1.rb:::14:::bind: Integer
@./ed339c31_1.rb:::15:::bind: Integer
@./ed339c31_1.rb:::18:::bind: String
@./ed339c31_1.rb:::19:::bind: String
@./ed339c31_1.rb:::20:::bind: String
@./ed339c31_1.rb:::21:::bind: String
@./ed339c31_1.rb:::24:::bind: Array<Integer>
@./ed339c31_1.rb:::25:::bind: Array<Integer String Float Bool NilClass Symbol>
@./ed339c31_1.rb:::26:::bind: Array<Array<Integer>>
@./ed339c31_1.rb:::27:::bind: Array<untyped>
@./ed339c31_1.rb:::30:::bind: Hash
@./ed339c31_1.rb:::31:::bind: Hash
@./ed339c31_1.rb:::32:::bind: Hash
@./ed339c31_1.rb:::33:::bind: Hash
@./ed339c31_1.rb:::36:::bind: Range
@./ed339c31_1.rb:::37:::bind: Range
@./ed339c31_1.rb:::40:::() -> NilClass [i/public]
@./ed339c31_1.rb:::44:::(Integer, Integer, Integer) -> Integer [i/public]
@./ed339c31_1.rb:::48:::(Integer, optional Integer, optional Integer) -> Integer [i/public]
@./ed339c31_1.rb:::52:::(untyped) -> Array<Integer> [i/public]
@./ed339c31_1.rb:::56:::() <block_params: NilClass> -> Union<NilClass untyped> [i/public]
@./ed339c31_1.rb:::65:::bind: String
@./ed339c31_1.rb:::64:::(String) -> BasicClass [c/public]
@./ed339c31_1.rb:::68:::() -> String [i/public]
@./ed339c31_1.rb:::73:::bind: String
@./ed339c31_1.rb:::72:::(String) -> String [i/public]
@./ed339c31_1.rb:::76:::() -> String [c/public]
@./ed339c31_1.rb:::84:::bind: Integer
@./ed339c31_1.rb:::82:::(String, Integer) -> ChildClass [c/public]
@./ed339c31_1.rb:::87:::() -> Integer [i/public]
@./ed339c31_1.rb:::91:::() -> String [i/public]
@./ed339c31_1.rb:::98:::() -> String [i/public]
@./ed339c31_1.rb:::106:::() -> String [i/public]
@./ed339c31_1.rb:::112:::bind: Integer
@./ed339c31_1.rb:::113:::bind: Float
@./ed339c31_1.rb:::148:::bind: Integer
@./ed339c31_1.rb:::155:::bind: Integer
@./ed339c31_1.rb:::174:::bind: Integer
@./ed339c31_1.rb:::185:::bind: Integer
@./ed339c31_1.rb:::186:::bind: Integer
@./ed339c31_1.rb:::187:::bind: Integer
@./ed339c31_1.rb:::188:::bind: Integer
@./ed339c31_1.rb:::189:::bind: Integer
@./ed339c31_1.rb:::190:::bind: Integer
@./ed339c31_1.rb:::193:::bind: Bool
@./ed339c31_1.rb:::194:::bind: Bool
@./ed339c31_1.rb:::195:::bind: Bool
@./ed339c31_1.rb:::196:::bind: Bool
@./ed339c31_1.rb:::197:::bind: Bool
@./ed339c31_1.rb:::198:::bind: Bool
@./ed339c31_1.rb:::201:::bind: Bool
@./ed339c31_1.rb:::202:::bind: Bool
@./ed339c31_1.rb:::203:::bind: Bool
@./ed339c31_1.rb:::206:::bind: Integer
@./ed339c31_1.rb:::213:::bind: String
@./ed339c31_1.rb:::214:::bind: String
@./ed339c31_1.rb:::215:::bind: String
@./ed339c31_1.rb:::216:::bind: String
@./ed339c31_1.rb:::219:::bind: Array<Integer>
@./ed339c31_1.rb:::222:::bind: Integer
@./ed339c31_1.rb:::223:::bind: Integer
@./ed339c31_1.rb:::224:::bind: Integer
@./ed339c31_1.rb:::227:::bind: Hash
@./ed339c31_1.rb:::228:::bind: Integer
@./ed339c31_1.rb:::229:::bind: Integer
@./ed339c31_1.rb:::230:::bind: Array<untyped>
@./ed339c31_1.rb:::233:::bind: Array<Integer>
@./ed339c31_1.rb:::234:::bind: Array<Integer>
@./ed339c31_1.rb:::237:::bind: String
@./ed339c31_1.rb:::238:::bind: String
@./ed339c31_1.rb:::239:::bind: String
@./ed339c31_1.rb:::245:::bind: String
@./ed339c31_1.rb:::246:::bind: Integer
@./ed339c31_1.rb:::244:::(String, Integer) -> Person [c/public]
@./ed339c31_1.rb:::249:::() -> String [i/public]
@./ed339c31_1.rb:::254:::bind: Unknown
@./ed339c31_1.rb:::253:::() -> unknown [i/public]
@./ed339c31_1.rb:::257:::() -> Integer [i/public]
@./ed339c31_1.rb:::262:::bind: Unknown
@./ed339c31_1.rb:::261:::() -> unknown [i/public]
@./ed339c31_1.rb:::265:::() -> String [i/public]
@./ed339c31_1.rb:::272:::bind: Integer
@./ed339c31_1.rb:::274:::() -> Counter [c/public]
@./ed339c31_1.rb:::278:::() -> Integer [c/public]
@./ed339c31_1.rb:::282:::() -> Integer [i/public]
@./ed339c31_1.rb:::289:::() -> String [i/public]
@./ed339c31_1.rb:::295:::() -> String [i/private]
@./ed339c31_1.rb:::301:::() -> String [i/protected]
@./ed339c31_1.rb:::307:::() -> String [i/public]
@./ed339c31_1.rb:::314:::() -> String [c/public]
@./ed339c31_1.rb:::318:::() -> String [i/public]
@./ed339c31_1.rb:::324:::bind: String
@./ed339c31_1.rb:::326:::() -> NilClass [i/public]
@./ed339c31_1.rb:::333:::() -> String [i/public]
@./ed339c31_1.rb:::338:::() -> Outer::Inner [i/public]
@./ed339c31_1.rb:::345:::bind: String
@./ed339c31_1.rb:::346:::bind: Integer
@./ed339c31_1.rb:::350:::bind: Proc
@./ed339c31_1.rb:::354:::bind: Integer
@./ed339c31_1.rb:::355:::bind: Integer
@./ed339c31_1.rb:::356:::bind: Integer
@./ed339c31_1.rb:::365:::bind: BasicClass
@./ed339c31_1.rb:::366:::bind: String
@./ed339c31_1.rb:::368:::bind: String
@./ed339c31_1.rb:::370:::bind: ChildClass
@./ed339c31_1.rb:::371:::bind: String
@./ed339c31_1.rb:::374:::bind: ClassWithModule
@./ed339c31_1.rb:::375:::bind: String
@./ed339c31_1.rb:::378:::bind: Person
@./ed339c31_1.rb:::379:::bind: String
@./ed339c31_1.rb:::380:::bind: String
@./ed339c31_1.rb:::381:::bind: Integer
@./ed339c31_1.rb:::384:::bind: Counter
@./ed339c31_1.rb:::385:::bind: Counter
@./ed339c31_1.rb:::386:::bind: Integer
@./ed339c31_1.rb:::389:::bind: VisibilityExample
@./ed339c31_1.rb:::390:::bind: String
@./ed339c31_1.rb:::391:::bind: String
@./ed339c31_1.rb:::394:::bind: String
@./ed339c31_1.rb:::395:::bind: SelfExample
@./ed339c31_1.rb:::396:::bind: String
@./ed339c31_1.rb:::402:::bind: Outer
@./ed339c31_1.rb:::403:::bind: Outer::Inner
@./ed339c31_1.rb:::404:::bind: String
@./ed339c31_1.rb:::407:::bind: String
@./ed339c31_1.rb:::408:::bind: Integer
@./ed339c31_1.rb:::411:::bind: untyped`

	if strings.TrimSpace(string(output)) != strings.TrimSpace(expectedOutput) {
		t.Errorf("Expected output: %s, but got: %s", expectedOutput, string(output))
	}
}
