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

package a2a

import "errors"

var (
	ErrTaskNotFound = errors.New("task not found")

	ErrTaskNotCancelable = errors.New("task cannot be canceled")

	ErrPushNotificationNotSupported = errors.New("push notification not supported")

	ErrUnsupportedOperation = errors.New("this operation is not supported")

	ErrUnsupportedContentType = errors.New("incompatible content types")

	ErrInvalidAgentResponse = errors.New("invalid agent response")

	ErrInvalidRequest = errors.New("invalid request")

	ErrAuthenticatedExtendedCardNotConfigured = errors.New("extended card not configured")
)
