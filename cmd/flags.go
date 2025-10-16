package cmd

type ExecuteFlags struct {
	IsDefineInfo    bool
	IsDefineAllInfo bool
	IsLsp           bool
}

func NewExecuteFlags() *ExecuteFlags {
	return &ExecuteFlags{}
}
