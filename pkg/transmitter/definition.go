package transmitter

import (
	"github.com/code-runner-project/code-runner-sdk/pkg/executor"
)

type Definition struct {
	Version         string
	Type            string
	Name            string
	Description     string
	OnEachComplete  func(data OnEachCompleteCallbackData)
	OnFullyComplete func(data OnFullyCompleteCallbackData)
}

type OnEachCompleteCallbackData struct {
	Status  int
	Results []executor.Result
	extras  map[string]interface{}
}

type OnFullyCompleteCallbackData struct {
	Status int
	Result executor.Result
	extras map[string]interface{}
}
