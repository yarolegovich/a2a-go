package a2asrv

import (
	"context"
	"iter"

	"github.com/a2aproject/a2a-go/a2a"
)

type RequestHandler interface {
	OnHandleGetTask(ctx context.Context, query a2a.TaskQueryParams) (a2a.Task, error)

	OnHandleCancelTask(ctx context.Context, id a2a.TaskIDParams) (a2a.Task, error)

	OnHandleSendMessage(ctx context.Context, message a2a.MessageSendParams) (a2a.SendMessageResult, error)

	OnHandleResubscribeToTask(ctx context.Context, id a2a.TaskIDParams) iter.Seq2[a2a.Event, error]

	OnHandleSendMessageStream(ctx context.Context, message a2a.MessageSendParams) iter.Seq2[a2a.Event, error]

	OnHandleGetTaskPushConfig(ctx context.Context, params a2a.GetTaskPushConfigParams) (a2a.TaskPushConfig, error)

	OnHandleListTaskPushConfig(ctx context.Context, params a2a.ListTaskPushConfigParams) ([]a2a.TaskPushConfig, error)

	OnHandleSetTaskPushConfig(ctx context.Context, params a2a.TaskPushConfig) (a2a.TaskPushConfig, error)

	OnHandleDeleteTaskPushConfig(ctx context.Context, params a2a.DeleteTaskPushConfigParams) error
}
