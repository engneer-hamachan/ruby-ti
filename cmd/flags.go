package cmd

type ExecuteFlags struct {
	IsDefineInfo    bool
	IsDefineAllInfo bool
	IsLsp           bool
	IsAllType       bool
}

func NewExecuteFlags() *ExecuteFlags {
	return &ExecuteFlags{}
}
