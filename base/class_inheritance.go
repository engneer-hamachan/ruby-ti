package base

type ClassNode struct {
	Frame string
	Class string
}

var ClassInheritanceMap = make(map[ClassNode][]ClassNode)
