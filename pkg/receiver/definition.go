package receiver

import "github.com/code-runner-project/code-runner-sdk/pkg/context"

type Definition struct {
	Version     string
	Type        string
	Name        string
	Description string
	OnStartFunc func(ctx *context.Context)
}
