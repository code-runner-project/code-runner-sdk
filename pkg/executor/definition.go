package executor

import (
	"time"
)

type Step struct {
	Inputs                 []string
	TempDir                string
	WorkDir                string
	Main                   string
	Cmd                    []string
	OnStartCallback        func(c OnStartCallbackData)
	OnEachCompleteCallback func(c OnEachCompleteCallbackData)
	OnFinishCallback       func(c OnFinishCallbackData)
}

type OnStartCallbackData struct {
	Id        string
	StartTime time.Time
	Main      string
	DirSize   int64
}

type OnEachCompleteCallbackData struct {
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
