package a2asrv

import (
	"context"
	"github.com/a2aproject/a2a-go/a2a"
)

type EventWriter interface {
	Write(ctx context.Context, event a2a.Event) error
}

type EventReader interface {
	Read(ctx context.Context) (a2a.Event, error)
}

type EventQueue interface {
	EventWriter
	EventReader

	Close()
}

type EventQueueManager interface {
	GetOrCreate(ctx context.Context, taskId a2a.TaskID) (EventQueue, error)

	Destroy(ctx context.Context, taskId a2a.TaskID) error
}
