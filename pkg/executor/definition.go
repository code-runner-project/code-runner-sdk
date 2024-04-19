package executor

import (
	"time"
)

type Definition struct {
	Version     string
	Type        string
	Name        string
	Description string
	ExecuteFunc func(s Step)
}

type Step struct {
	Inputs                 []string
	TempDir                string
	WorkDir                string
	Main                   string
	Cmd                    []string
	OnStartCallback        func(c OnStartCallbackData)
	OnEachCompleteCallback func(c OnTestCompleteCallbackData)
	OnFinishCallback       func(c OnFinishCallbackData)
}

type OnStartCallbackData struct {
	Id        string
	StartTime time.Time
	Main      string
	DirSize   int64
}

type OnTestCompleteCallbackData struct {
	Id         string
	Ref        string
	StartTime  time.Time
	FinishTime time.Time
	TotalTime  int64
	Output     string
	StatusCode int8
	Error      bool
}

type OnFinishCallbackData struct {
	Id         string
	FinishTime time.Time
	TotalTime  int64
}

type Result struct {
	Actual   string
	Expected string
	Status   int
}
