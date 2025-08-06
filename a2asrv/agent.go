package a2asrv

import (
	"context"
	"github.com/a2aproject/a2a-go/a2a"
)

type AgentExecutor interface {
	Execute(ctx context.Context, reqCtx RequestContext, queue EventQueue) error

	Cancel(ctx context.Context, reqCtx RequestContext, queue EventQueue) error
}

type AgentCardProducer interface {
	Card() a2a.AgentCard
}
