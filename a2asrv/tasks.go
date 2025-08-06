package a2asrv

import (
	"context"
	"github.com/a2aproject/a2a-go/a2a"
)

type RequestContextBuilder interface {
	Build(ctx context.Context, p a2a.MessageSendParams, t *a2a.Task) RequestContext
}

type PushNotifier interface {
	SendPush(ctx context.Context, task a2a.Task) error
}

type PushConfigStore interface {
	Save(ctx context.Context, taskId a2a.TaskID, config a2a.PushConfig) error

	Get(ctx context.Context, taskId a2a.TaskID) ([]a2a.PushConfig, error)

	Delete(ctx context.Context, taskId a2a.TaskID) error
}

type TaskStore interface {
	Save(ctx context.Context, task a2a.Task) error

	Get(ctx context.Context, taskId a2a.TaskID) (a2a.Task, error)
}
