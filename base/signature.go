package base

type Sig struct {
	Contents string
	Detail   string
	Frame    string
	Class    string
	IsStatic bool
	FileName string
	Row      int
}

var TSignatures = make(map[string]Sig)
