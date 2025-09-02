package eval

import (
	"ti/base"
	"ti/context"
	"ti/parser"
)

var DynamicEvaluators = make(map[string]DynamicEvaluator)

type DynamicEvaluator interface {
	Evaluation(
		e *Evaluator,
		p *parser.Parser,
		ctx context.Context,
		t *base.T,
	) (err error)
}
