package base

type DefinedClass struct {
	frame string
	class string
}

var DefinedClassTable = make(map[DefinedClass]bool)

func init() {
	class := "Object"

	classNode := ClassNode{Frame: "Builtin", Class: class}
	objectClassNode := ClassNode{Frame: "Builtin", Class: ""}

	ClassInheritanceMap[classNode] =
		append(ClassInheritanceMap[classNode], objectClassNode)

	returnT := MakeObject(class)
	args := "*" + GenId()
	methodT := MakeMethod("Builtin", "new", *returnT, []string{args})
	SetClassMethodT("", class, methodT, false, "unknown", 0)

	DefinedClassTable[DefinedClass{"Builtin", class}] = true
}

func IsClassDefined(frames []string, class string) bool {
	for _, frame := range frames {
		key := DefinedClass{frame: frame, class: class}
		_, ok := DefinedClassTable[key]

		if ok {
			return ok
		}
	}

	key := DefinedClass{frame: "Builtin", class: class}
	_, ok := DefinedClassTable[key]

	return ok
}

func SetDefinedClass(frame, class string) {
	key := DefinedClass{frame: frame, class: class}
	DefinedClassTable[key] = true
}
