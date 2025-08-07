// Copyright 2025 The A2A Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
