package context

type Context struct {
	frame          string
	class          string
	method         string
	IsPrivate      bool
	IsProtected    bool
	IsBind         bool
	IsCallArg      bool
	IsDefineArg    bool
	IsArrayCollect bool
	IsDefineStatic bool
	round          string
}

func NewContext(class string, method string, round string) Context {
	return Context{
		class:  class,
		method: method,
		round:  round,
	}
}

func GetRounds() []string {
	return []string{"collect", "inference", "check"}
}

func (c *Context) SetFrame(frame string) {
	c.frame = frame
}

func (c *Context) GetFrame() string {
	return c.frame
}

func (c *Context) SetClass(class string) {
	c.class = class
}

func (c *Context) GetClass() string {
	return c.class
}

func (c *Context) SetMethod(method string) {
	c.method = method
}

func (c *Context) GetMethod() string {
	return c.method
}

func (c Context) IsCheckRound() bool {
	return c.round == "check"
}

func (c Context) IsInferenceRound() bool {
	return c.round == "inference"
}

func (c Context) IsCollectRound() bool {
	return c.round == "collect"
}

func (c *Context) StartPrivate() {
	c.IsPrivate = true
	c.IsProtected = false
}

func (c *Context) EndPrivate() {
	c.IsPrivate = false
}

func (c *Context) StartProtected() {
	c.IsProtected = true
	c.IsPrivate = false
}

func (c *Context) EndProtected() {
	c.IsProtected = false
}

func (c *Context) StartCallArg() {
	c.IsCallArg = true
}

func (c *Context) EndCallArg() {
	c.IsCallArg = false
}

func (c *Context) StartDefineArg() {
	c.IsDefineArg = true
}

func (c *Context) EndDefineArg() {
	c.IsDefineArg = false
}

func (c *Context) StartArrayCollect() {
	c.IsArrayCollect = true
}

func (c *Context) EndArrayCollect() {
	c.IsArrayCollect = false
}

func (c *Context) StartDefineStatic() {
	c.IsDefineStatic = true
}

func (c *Context) EndDefineStatic() {
	c.IsDefineStatic = false
}
