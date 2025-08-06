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

	ErrAuthenticatedExtendedCardNotConfigured = errors.New("")
)
