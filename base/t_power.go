package base

var tPower = make(map[string]int8)

func init() {
	tPower["+"] = 20
	tPower["-"] = 20
	tPower["*"] = 20
	tPower["/"] = 20
	tPower["%"] = 20
	tPower["."] = 30
	tPower["=="] = 10
}

func (t *T) IsPowerUp(otherT *T) bool {
	return t.GetPower() <= otherT.GetPower()
}

func (t *T) GetPower() int8 {
	power, ok := tPower[t.ToString()]
	if ok {
		return power
	}

	return 0
}
