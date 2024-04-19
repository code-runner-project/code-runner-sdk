package transmitter

import (
	"github.com/code-runner-project/code-runner-sdk/pkg/executor"
)

type Definition struct {
	Version       string
	Type          string
	Name          string
	Description   string
	OnReceiveFunc func(data OnReceiveCallbackData)
}

type OnReceiveCallbackData struct {
	Status  int
	Results []executor.Result
	extras  map[string]interface{}
}
