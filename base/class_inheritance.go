package base

type ClassNode struct {
	Frame    string
	Class    string
	IsExtend bool
}

var ClassInheritanceMap = make(map[ClassNode][]ClassNode)
