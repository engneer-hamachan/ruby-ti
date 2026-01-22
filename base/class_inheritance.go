package base

type ClassNode struct {
	Frame     string
	Class     string
	IsInclude bool
	IsExtend  bool
}

var ClassInheritanceMap = make(map[ClassNode][]ClassNode)
