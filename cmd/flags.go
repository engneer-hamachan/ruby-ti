package cmd

type ExecuteFlags struct {
	IsDefineInfo    bool
	IsDefineAllInfo bool
	IsSuggest       bool
	IsAllType       bool
	IsExtends       bool
	IsHover         bool
	IsVersion       bool
	IsHelp          bool
	IsLlmInfo       bool
	IsLlmError      bool
	IsLlmDefine     bool
	IsLlmClass      bool
}

func NewExecuteFlags() *ExecuteFlags {
	return &ExecuteFlags{}
}
