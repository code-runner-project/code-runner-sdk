package context

import (
	"github.com/code-runner-project/code-runner-sdk/pkg/orchestra"
)

type Message struct {
	FlowDefinition string
}

type Context struct {
	orchestrator *orchestra.Orchestrator
}

func (c *Context) Proceed(message Message) {
	c.orchestrator.Execute(message)
}

func NewContext(orchestrator *orchestra.Orchestrator) Context {
	return Context{orchestrator: orchestrator}
}
