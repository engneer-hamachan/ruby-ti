package base

var tPower = make(map[string]int8)

func init() {
	tPower["!"] = 100
	tPower["~"] = 100
	tPower["*"] = 70
	tPower["/"] = 70
	tPower["%"] = 70
	tPower["+"] = 60
	tPower["-"] = 60
	tPower["<<"] = 55
	tPower[">>"] = 55
	tPower["&"] = 50
	tPower["|"] = 45
	tPower["^"] = 45
	tPower[">"] = 40
	tPower[">="] = 40
	tPower["<"] = 40
	tPower["<="] = 40
	tPower["<=>"] = 35
	tPower["=="] = 35
	tPower["==="] = 35
	tPower["!="] = 35
	tPower["=~"] = 35
	tPower["!~"] = 35
	tPower["&&"] = 30
	tPower["||"] = 25
	tPower[".."] = 20
	tPower["..."] = 20
	tPower["?"] = 15
	tPower[":"] = 15
	tPower["="] = 10
	tPower["+="] = 10
	tPower["-="] = 10
	tPower["*="] = 10
	tPower["/="] = 10
	tPower["%="] = 10
	tPower["**="] = 10
	tPower["<<="] = 10
	tPower[">>="] = 10
	tPower["&="] = 10
	tPower["|="] = 10
	tPower["^="] = 10
	tPower["||="] = 10
	tPower["&&="] = 10
	tPower["."] = 95
	tPower["defined?"] = 5
	tPower["not"] = 3
	tPower["or"] = 2
	tPower["and"] = 2
}

func (t *T) IsNotPowerDown(otherT *T) bool {
	return t.GetPower() <= otherT.GetPower()
}

func (t *T) GetPower() int8 {
	power, ok := tPower[t.ToString()]
	if ok {
		return power
	}

	return 0
}
