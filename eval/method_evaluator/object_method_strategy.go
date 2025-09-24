package method_evaluator

import (
	"fmt"
	"slices"
	"strings"
	"ti/base"
)

func init() {
	dynamicStrategies[[2]string{"", "class"}] = &objectClassStrategy{}

	dynamicStrategies[[2]string{"", "attr_accessor"}] =
		&objectAttrAccessorStrategy{}

	dynamicStrategies[[2]string{"", "attr_reader"}] =
		&objectAttrReaderStrategy{}

	dynamicStrategies[[2]string{"", "include"}] =
		&objectIncludeStrategy{}

	dynamicStrategies[[2]string{"", "raise"}] =
		&objectRaiseStrategy{}
}

type objectIncludeStrategy struct{}

func (o *objectIncludeStrategy) evaluate(m *MethodEvaluator) error {
	nextT, err := m.parser.Read()
	if err != nil {
		return err
	}

	classNode := base.ClassNode{Frame: m.ctx.GetFrame(), Class: m.ctx.GetClass()}

	parentFrame, parentNamespace, parentClass := nextT.SeparateNameSpaces()

	parentFrame = base.CalculateFrame(parentFrame, parentNamespace)
	parentNode := base.ClassNode{Frame: parentFrame, Class: parentClass}

	if slices.Contains(base.ClassInheritanceMap[classNode], parentNode) {
		return nil
	}

	base.ClassInheritanceMap[classNode] =
		append(base.ClassInheritanceMap[classNode], parentNode)

	return nil
}

type objectAttrReaderStrategy struct{}

func (o *objectAttrReaderStrategy) evaluate(m *MethodEvaluator) error {
	for {
		nextT, err := m.parser.Read()
		if err != nil {
			return err
		}

		switch nextT.ToString() {
		case "\n":
			m.parser.Unget()
			return nil

		case ",":
			continue

		default:
			if !nextT.IsSymbolType() {
				return fmt.Errorf(
					"expected symbol, but got '%s'",
					nextT.ToString(),
				)
			}

			identifier := strings.TrimPrefix(nextT.ToString(), ":")
			t := base.MakeIdentifier(identifier)
			t.EnableReadOnly()

			currentT :=
				base.GetInstanceValueT(
					m.ctx.GetFrame(),
					m.ctx.GetClass(),
					identifier,
				)

			if currentT == nil {
				base.SetInstanceValueT(
					m.ctx.GetFrame(),
					m.ctx.GetClass(),
					identifier,
					t,
				)
			}
		}
	}
}

type objectAttrAccessorStrategy struct{}

func (o *objectAttrAccessorStrategy) evaluate(m *MethodEvaluator) error {
	for {
		nextT, err := m.parser.Read()
		if err != nil {
			return err
		}

		switch nextT.ToString() {
		case "\n":
			m.parser.Unget()
			return nil

		case ",":
			continue

		default:
			if !nextT.IsSymbolType() && m.ctx.IsCheckRound() {
				return fmt.Errorf("Expected symbol, but got '%s'", nextT.ToString())
			}

			identifier := strings.TrimPrefix(nextT.ToString(), ":")
			nilT := base.MakeNil()

			base.SetInstanceValueT(
				m.ctx.GetFrame(),
				m.ctx.GetClass(),
				identifier,
				nilT,
			)
		}
	}
}

type objectClassStrategy struct{}

func (o *objectClassStrategy) evaluate(m *MethodEvaluator) error {
	methodT := base.GetMethodT("Builtin", "", "class", false)

	if methodT == nil {
		return m.makeNotDefinedMethodError("", "class")
	}

	evaluatedArgs, err := getEvaluatedArgs(m, methodT)
	if err != nil {
		return err
	}

	err = checkAndPropagateArgs(m, "", methodT, evaluatedArgs)
	if err != nil {
		return err
	}

	classT := base.MakeObject("Class")
	m.parser.SetLastEvaluatedT(classT)

	return nil
}

type objectRaiseStrategy struct{}

func (o *objectRaiseStrategy) evaluate(m *MethodEvaluator) error {
	methodT := base.GetMethodT("Builtin", "", "raise", false)

	if methodT == nil {
		return m.makeNotDefinedMethodError("", "raise")
	}

	evaluatedArgs, err := getEvaluatedArgs(m, methodT)
	if err != nil {
		return err
	}

	err = checkAndPropagateArgs(m, "", methodT, evaluatedArgs)
	if err != nil {
		return err
	}

	m.parser.SetLastReturnT(methodT)
	m.parser.SetLastEvaluatedT(methodT)

	return nil
}
