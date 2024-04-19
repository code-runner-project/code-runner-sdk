package context

import (
	"github.com/code-runner-project/code-runner-sdk/pkg/executor"
	"github.com/code-runner-project/code-runner-sdk/pkg/receiver"
	"github.com/code-runner-project/code-runner-sdk/pkg/runner"
	"github.com/code-runner-project/code-runner-sdk/pkg/transmitter"
)

type Message struct {
	FlowDefinition string
}

type Context struct {
	Message      Message
	Runners      map[string]runner.Definition
	Executors    map[string]executor.Definition
	Receivers    map[string]receiver.Definition
	Transmitters map[string]transmitter.Definition
}

func (o *Context) RegisterRunner(name string, runner runner.Definition) {
	o.Runners[name] = runner
}

func (o *Context) RegisterExecutor(name string, executor executor.Definition) {
	o.Executors[name] = executor
}

func (o *Context) RegisterReceiver(name string, receiver receiver.Definition) {
	o.Receivers[name] = receiver
}

func (o *Context) RegisterTransmitter(name string, transmitter transmitter.Definition) {
	o.Transmitters[name] = transmitter
}

func (c *Context) Proceed(message Message) {
	c.Message = message
}

func NewContext() *Context {
	return &Context{
		Message: Message{
			FlowDefinition: "",
		},
		Runners:      make(map[string]runner.Definition),
		Executors:    make(map[string]executor.Definition),
		Receivers:    make(map[string]receiver.Definition),
		Transmitters: make(map[string]transmitter.Definition),
	}
}
