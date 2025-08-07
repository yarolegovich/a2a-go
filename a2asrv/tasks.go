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

	"github.com/a2aproject/a2a-go/a2a"
)

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
