package parser

import (
	"fmt"
	"ti/base"
	"ti/context"
	"ti/lexer"
)

type Parser struct {
	Lexer               lexer.Lexer
	token               rune
	ungetFlg            bool
	FileName            string
	Row                 int
	ErrorRow            int
	lastEvaluatedT      any
	LastCallT           *base.T
	lastCallFrame       [3]string
	isParsingExpression bool
	lastReturnT         []base.T
	tmpBlockParamaters  []base.T
	tmpEvaluatedArgs    []*base.T
	Debug               bool
	IsDefineInfo        bool
	Errors              []error
	DefineInfos         []string
	BeforeString        string
}

func New(lexer lexer.Lexer, file string) Parser {
	return Parser{
		Lexer:    lexer,
		ungetFlg: false,
		FileName: file,
		Row:      1,
		Debug:    false,
	}
}

func (p *Parser) Fatal(ctx context.Context, err error) {
	if ctx.IsCheckRound() {
		p.Errors =
			append(p.Errors, fmt.Errorf("%v::%d::%v", p.FileName, p.ErrorRow, err))
	}
}

func (p *Parser) SetLastEvaluatedT(some any) {
	p.StartParsingExpression()
	p.lastEvaluatedT = some
}

func (p *Parser) GetLastEvaluatedT() base.T {
	if p.lastEvaluatedT == nil {
		return *base.MakeUntyped()
	}

	var t *base.T

	switch p.lastEvaluatedT.(type) {
	case *base.T:
		t = p.lastEvaluatedT.(*base.T)
	default:
		t = base.MakeUnknown()
	}

	return *t.DeepCopy()
}

func (p *Parser) GetLastEvaluatedTPointer() any {
	return p.lastEvaluatedT
}

func (p *Parser) SetTmpBlockParameters(parameters []base.T) {
	p.tmpBlockParamaters = parameters
}

func (p *Parser) GetTmpBlockParameters() []base.T {
	return p.tmpBlockParamaters
}

func (p *Parser) ClearTmpBlockParameters() {
	p.tmpBlockParamaters = []base.T{}
}

func (p *Parser) SetLastReturnT(t *base.T) {
	p.lastReturnT = append(p.lastReturnT, *t)
}

func (p *Parser) AppendLastReturnT() {
	if p.lastEvaluatedT == nil {
		p.lastReturnT = append(p.lastReturnT, *base.MakeNil())

		return
	}

	for _, candidateT := range p.lastReturnT {
		if candidateT.IsMatchType(p.lastEvaluatedT.(*base.T)) {
			return
		}
	}

	if p.lastEvaluatedT.(*base.T) == nil {
		p.lastReturnT = append(p.lastReturnT, *base.MakeNil())

		return
	}

	lastEvaluatedT := p.lastEvaluatedT.(*base.T)

	if lastEvaluatedT.IsUnionType() {
		p.lastReturnT = append(p.lastReturnT, lastEvaluatedT.GetVariants()...)
		return
	}

	p.lastReturnT = append(p.lastReturnT, *lastEvaluatedT)
}

func (p *Parser) GetLastReturnT() []base.T {
	return p.lastReturnT
}

func (p *Parser) ConsumeLastReturnT() []base.T {
	returnTs := p.lastReturnT

	p.lastReturnT = []base.T{}

	return returnTs
}

func (p *Parser) IsParsingExpression() bool {
	return p.isParsingExpression
}

func (p *Parser) StartParsingExpression() {
	p.isParsingExpression = true
}

func (p *Parser) EndParsingExpression() {
	p.isParsingExpression = false
}

func (p *Parser) SetLastCallFrameDetails(frame, class, method string) {
	p.lastCallFrame = [3]string{frame, class, method}
}

func (p *Parser) GetLastCallFrameDetails() (string, string, string) {
	return p.lastCallFrame[0], p.lastCallFrame[1], p.lastCallFrame[2]
}

func (p *Parser) SetTmpEvaluaetdArgs(args []*base.T) {
	p.tmpEvaluatedArgs = args
}

func (p *Parser) GetTmpEvaluaetdArgs() []*base.T {
	return p.tmpEvaluatedArgs
}

func (p *Parser) ClearTmpEvaluaetdArgs() {
	p.tmpEvaluatedArgs = []*base.T{}
}
