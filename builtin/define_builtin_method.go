package builtin

import (
	"ti/base"
)

type defineBuiltinMethod struct {
	frame       string
	targetClass string
}

func NewDefineBuiltinMethod(
	frame string,
	class string,
) *defineBuiltinMethod {

	d := &defineBuiltinMethod{
		frame:       frame,
		targetClass: class,
	}

	return d
}

func (d *defineBuiltinMethod) setupMethodArgs(
	method string,
	argTypes []base.T,
) []string {

	var argIdentifiers []string

	for _, argType := range argTypes {
		switch argType.IsKeyValueType() {
		case true:
			argIdentifiers = append(argIdentifiers, argType.GetKey())

			base.SetValueT(
				d.frame,
				d.targetClass,
				method,
				argType.GetRemoveSuffixKey(),
				argType.GetKeyValue(),
			)

		default:
			id := base.GenId()

			if argType.IsBuiltinAsterisk {
				id = "*" + id
			}

			argIdentifiers = append(argIdentifiers, id)
			base.SetValueT(d.frame, d.targetClass, method, id, &argType)
		}
	}

	return argIdentifiers
}

func (d *defineBuiltinMethod) defineBuiltinInstanceMethod(
	frame string,
	method string,
	argTypes []base.T,
	returnT base.T,
) {

	argIdentifiers := d.setupMethodArgs(method, argTypes)
	methodT := base.MakeMethod(frame, method, returnT, argIdentifiers)

	base.SetMethodT(frame, d.targetClass, methodT, false, "unknown", 0)
}

func (d *defineBuiltinMethod) defineBuiltinStaticMethod(
	frame string,
	method string,
	argTypes []base.T,
	returnT base.T,
) {

	argIdentifiers := d.setupMethodArgs(method, argTypes)
	methodT := base.MakeMethod(frame, method, returnT, argIdentifiers)

	base.SetClassMethodT(frame, d.targetClass, methodT, false, "unknown", 0)
}

func (d *defineBuiltinMethod) defineBuiltinConstant(
	frame string,
	class string,
	variable string,
	returnT base.T,
) {

	base.SetConstValueT(frame, class, variable, &returnT)
}

func (d *defineBuiltinMethod) SetDefinedClass() {
	base.SetDefinedClass(d.frame, d.targetClass)
}
