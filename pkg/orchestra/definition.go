package orchestra

import (
	"github.com/code-runner-project/code-runner-sdk/pkg/context"
	"github.com/code-runner-project/code-runner-sdk/pkg/executor"
	"github.com/code-runner-project/code-runner-sdk/pkg/receiver"
	"github.com/code-runner-project/code-runner-sdk/pkg/runner"
	"github.com/code-runner-project/code-runner-sdk/pkg/transmitter"
)

type Orchestrator struct {
	message      context.Message
	runners      map[string]runner.Definition
	executors    map[string]executor.Definition
	receivers    map[string]receiver.Definition
	transmitters map[string]transmitter.Definition
}

func (o *Orchestrator) Execute(message context.Message) {
	o.message = message
}

func (o *Orchestrator) RegisterRunner(name string, runner runner.Definition) {
	o.runners[name] = runner
}

func (o *Orchestrator) RegisterExecutor(name string, executor executor.Definition) {
	o.executors[name] = executor
}

func (o *Orchestrator) RegisterReceiver(name string, receiver receiver.Definition) {
	o.receivers[name] = receiver
}

func (o *Orchestrator) RegisterTransmitter(name string, transmitter transmitter.Definition) {
	o.transmitters[name] = transmitter
}
