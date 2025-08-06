package a2a

// Defines parameters for fetching a specific push notification configuration for a task.
type GetTaskPushConfigParams struct {
	// The unique identifier of the task.
	TaskID TaskID

	// The ID of the push notification configuration to retrieve.
	ConfigID *string

	// Optional metadata associated with the request.
	Metadata map[string]any
}

// Defines parameters for listing all push notification configurations associated
// with a task.
type ListTaskPushConfigParams struct {
	// The unique identifier of the task.
	TaskID TaskID

	// Optional metadata associated with the request.
	Metadata map[string]any
}

// Defines parameters for deleting a specific push notification configuration for a task.
type DeleteTaskPushConfigParams struct {
	// The unique identifier of the task.
	TaskID TaskID

	// Optional metadata associated with the request.
	Metadata map[string]any

	// The ID of the push notification configuration to delete.
	ConfigID string
}

// A container associating a push notification configuration with a specific task.
type TaskPushConfig struct {
	// The push notification configuration for this task.
	Config PushConfig

	// The ID of the task.
	TaskID TaskID
}

// Defines the configuration for setting up push notifications for task updates.
type PushConfig struct {
	// Optional authentication details for the agent to use when calling the
	// notification URL.
	Auth *PushAuthInfo

	// A unique ID for the push notification configuration, set by the client
	// to support multiple notification callbacks.
	ID *string

	// A unique token for this task or session to validate incoming push
	// notifications.
	Token *string

	// The callback URL where the agent should send push notifications.
	URL string
}

// Defines authentication details for a push notification endpoint.
type PushAuthInfo struct {
	// Optional credentials required by the push notification endpoint.
	Credentials *string

	// A list of supported authentication schemes (e.g., 'Basic', 'Bearer').
	Schemes []string
}
