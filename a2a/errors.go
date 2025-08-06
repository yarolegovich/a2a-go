package a2a

import "fmt"

var ErrTaskNotFound = fmt.Errorf("task not found")

var ErrTaskNotCancelable = fmt.Errorf("task cannot be canceled")

var ErrPushNotificationNotSupported = fmt.Errorf("push notification not supported")

var ErrUnsupportedOperation = fmt.Errorf("this operation is not supported")

var ErrUnsupportedContentType = fmt.Errorf("incompatible content types")

var ErrInvalidAgentResponse = fmt.Errorf("invalid agent response")

var ErrInvalidRequest = fmt.Errorf("invalid request")
