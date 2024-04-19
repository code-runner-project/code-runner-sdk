package runner

import "github.com/code-runner-project/code-runner-sdk/pkg/executor"

type Definition struct {
	Version                 string
	Type                    string
	Name                    string
	Description             string
	OnInitFunc              func(s Step) executor.Step
	OnExecutorStartFunc     func(c executor.OnStartCallbackData)
	OnExecutionCompleteFunc func(c executor.OnTestCompleteCallbackData)
	OnExecutorFinishFunc    func(c executor.OnFinishCallbackData)
}

type Step struct {
	Files    []File
	Tests    []Test
	Language string
}

type Test struct {
	Input          string
	ExpectedOutput string
}

type File struct {
	Ref  string
	Path string
	Name string
	Ext  string
}
