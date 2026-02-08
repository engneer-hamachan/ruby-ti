package cmd

type ExecuteFlags struct {
	IsDefineInfo    bool
	IsDefineAllInfo bool
	IsSuggest       bool
	IsAllType       bool
	IsExtends       bool
	IsHover         bool
	IsVersion       bool
	IsLlmInfo       bool
}

func NewExecuteFlags() *ExecuteFlags {
	return &ExecuteFlags{}
}
