package orchestra

import (
	"github.com/code-runner-project/code-runner-sdk/pkg/context"
	"github.com/code-runner-project/code-runner-sdk/pkg/executor"
	"github.com/code-runner-project/code-runner-sdk/pkg/receiver"
	"github.com/code-runner-project/code-runner-sdk/pkg/runner"
	"github.com/code-runner-project/code-runner-sdk/pkg/transmitter"
)

type Orchestrator struct {
	Message      context.Message
	Runners      map[string]runner.Definition
	Executors    map[string]executor.Definition
	Receivers    map[string]receiver.Definition
	Transmitters map[string]transmitter.Definition
}

func (o *Orchestrator) Execute(message context.Message) {
	o.Message = message
}

func (o *Orchestrator) RegisterRunner(name string, runner runner.Definition) {
	o.Runners[name] = runner
}

func (o *Orchestrator) RegisterExecutor(name string, executor executor.Definition) {
	o.Executors[name] = executor
}

func (o *Orchestrator) RegisterReceiver(name string, receiver receiver.Definition) {
	o.Receivers[name] = receiver
}

func (o *Orchestrator) RegisterTransmitter(name string, transmitter transmitter.Definition) {
	o.Transmitters[name] = transmitter
}

func NewOrchestrator() *Orchestrator {
	return &Orchestrator{
		Message: context.Message{
			FlowDefinition: "",
		},
		Runners:      make(map[string]runner.Definition),
		Executors:    make(map[string]executor.Definition),
		Receivers:    make(map[string]receiver.Definition),
		Transmitters: make(map[string]transmitter.Definition),
	}
}
