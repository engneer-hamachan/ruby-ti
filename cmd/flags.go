package cmd

const Version = "v1.0.11"

type ExecuteFlags struct {
	IsDefineInfo    bool
	IsDefineAllInfo bool
	IsSuggest       bool
	IsAllType       bool
	IsExtends       bool
	IsHover         bool
	IsVersion       bool
}

func NewExecuteFlags() *ExecuteFlags {
	return &ExecuteFlags{}
}
