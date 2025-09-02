package lexer

var tbl map[string]Identifier = make(map[string]Identifier)

type Identifier struct {
	name string
}

func newIdentifier(name string) Identifier {
	id := Identifier{name: name}

	tbl[name] = id

	return id
}

func (i *Identifier) GetName() string {
	return i.name
}

func Intern(name string) Identifier {
	val, ok := tbl[name]
	if ok {
		return val
	}

	return newIdentifier(name)
}
