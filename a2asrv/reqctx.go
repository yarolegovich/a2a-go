package a2asrv

import (
	"context"
	"github.com/a2aproject/a2a-go/a2a"
)

type RequestContextBuilder interface {
	Build(ctx context.Context, p a2a.MessageSendParams, t *a2a.Task) RequestContext
}

type RequestContext struct {
	Request      a2a.MessageSendParams
	TaskID       a2a.TaskID
	Task         *a2a.Task
	RelatedTasks []a2a.Task
	ContextID    string
}
