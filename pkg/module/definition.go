package module

import (
	"github.com/code-runner-project/code-runner-sdk/pkg/executor"
	"github.com/code-runner-project/code-runner-sdk/pkg/runner"
)

type ExecutorDefinition struct {
	Version     string
	Id          string
	Type        string
	Name        string
	Description string
	ExecuteFunc func(e executor.Step)
}

type RunnerDefinition struct {
	Version                 string
	Id                      string
	Type                    string
	Name                    string
	Description             string
	OnInitFunc              func(bundle runner.Step) executor.Step
	OnExecutorStartFunc     func(c executor.OnStartCallbackData)
	OnExecutionCompleteFunc func(c executor.OnTestCompleteCallbackData)
	OnExecutorFinishFunc    func(c executor.OnFinishCallbackData)
}
