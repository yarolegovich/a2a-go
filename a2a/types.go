package a2a

// Defines a security scheme using an API key.
type APIKeySecurityScheme struct {
	// An optional description for the security scheme.
	Description *string

	// The location of the API key.
	In APIKeySecuritySchemeIn

	// The name of the header, query, or cookie parameter to be used.
	Name string

	// The type of the security scheme. Must be 'apiKey'.
	Type string
}

type APIKeySecuritySchemeIn string

const (
	APIKeySecuritySchemeInCookie APIKeySecuritySchemeIn = "cookie"
	APIKeySecuritySchemeInHeader APIKeySecuritySchemeIn = "header"
	APIKeySecuritySchemeInQuery  APIKeySecuritySchemeIn = "query"
)

// Defines optional capabilities supported by an agent.
type AgentCapabilities struct {
	// A list of protocol extensions supported by the agent.
	Extensions []AgentExtension

	// Indicates if the agent supports sending push notifications for asynchronous
	// task updates.
	PushNotifications *bool

	// Indicates if the agent provides a history of state transitions for a task.
	StateTransitionHistory *bool

	// Indicates if the agent supports Server-Sent Events (SSE) for streaming
	// responses.
	Streaming *bool
}

// The AgentCard is a self-describing manifest for an agent. It provides essential
// metadata including the agent's identity, capabilities, skills, supported
// communication methods, and security requirements.
type AgentCard struct {
	// A list of additional supported interfaces (transport and URL combinations).
	// This allows agents to expose multiple transports, potentially at different
	// URLs.
	//
	// Best practices:
	// - SHOULD include all supported transports for completeness
	// - SHOULD include an entry matching the main 'url' and 'preferredTransport'
	// - MAY reuse URLs if multiple transports are available at the same endpoint
	// - MUST accurately declare the transport available at each URL
	//
	// Clients can select any interface from this list based on their transport
	// capabilities
	// and preferences. This enables transport negotiation and fallback scenarios.
	AdditionalInterfaces []AgentInterface

	// A declaration of optional capabilities supported by the agent.
	Capabilities AgentCapabilities

	// Default set of supported input MIME types for all skills, which can be
	// overridden on a per-skill basis.
	DefaultInputModes []string

	// Default set of supported output MIME types for all skills, which can be
	// overridden on a per-skill basis.
	DefaultOutputModes []string

	// A human-readable description of the agent, assisting users and other agents
	// in understanding its purpose.
	Description string

	// An optional URL to the agent's documentation.
	DocumentationURL *string

	// An optional URL to an icon for the agent.
	IconURL *string

	// A human-readable name for the agent.
	Name string

	// The transport protocol for the preferred endpoint (the main 'url' field).
	// If not specified, defaults to 'JSONRPC'.
	//
	// IMPORTANT: The transport specified here MUST be available at the main 'url'.
	// This creates a binding between the main URL and its supported transport
	// protocol.
	// Clients should prefer this transport and URL combination when both are
	// supported.
	PreferredTransport string

	// The version of the A2A protocol this agent supports.
	ProtocolVersion string

	// Information about the agent's service provider.
	Provider *AgentProvider

	// A list of security requirement objects that apply to all agent interactions.
	// Each object
	// lists security schemes that can be used. Follows the OpenAPI 3.0 Security
	// Requirement Object.
	// This list can be seen as an OR of ANDs. Each object in the list describes one
	// possible
	// set of security requirements that must be present on a request. This allows
	// specifying,
	// for example, "callers must either use OAuth OR an API Key AND mTLS."
	Security []map[string][]string

	// A declaration of the security schemes available to authorize requests. The key
	// is the
	// scheme name. Follows the OpenAPI 3.0 Security Scheme Object.
	SecuritySchemes map[string]any

	// JSON Web Signatures computed for this AgentCard.
	Signatures []AgentCardSignature

	// The set of skills, or distinct capabilities, that the agent can perform.
	Skills []AgentSkill

	// If true, the agent can provide an extended agent card with additional details
	// to authenticated users. Defaults to false.
	SupportsAuthenticatedExtendedCard *bool

	// The preferred endpoint URL for interacting with the agent.
	// This URL MUST support the transport specified by 'preferredTransport'.
	URL string

	// The agent's own version number. The format is defined by the provider.
	Version string
}

// AgentCardSignature represents a JWS signature of an AgentCard.
// This follows the JSON format of an RFC 7515 JSON Web Signature (JWS).
type AgentCardSignature struct {
	// The unprotected JWS header values.
	Header map[string]any

	// The protected JWS header for the signature. This is a Base64url-encoded
	// JSON object, as per RFC 7515.
	Protected string

	// The computed signature, Base64url-encoded.
	Signature string
}

// A declaration of a protocol extension supported by an Agent.
type AgentExtension struct {
	// A human-readable description of how this agent uses the extension.
	Description *string

	// Optional, extension-specific configuration parameters.
	Params map[string]any

	// If true, the client must understand and comply with the extension's
	// requirements
	// to interact with the agent.
	Required *bool

	// The unique URI identifying the extension.
	Uri string
}

// Declares a combination of a target URL and a transport protocol for interacting
// with the agent.
// This allows agents to expose the same functionality over multiple transport
// mechanisms.
type AgentInterface struct {
	// The transport protocol supported at this URL.
	Transport string

	// The URL where this interface is available. Must be a valid absolute HTTPS URL
	// in production.
	URL string
}

// Represents the service provider of an agent.
type AgentProvider struct {
	// The name of the agent provider's organization.
	Org string

	// A URL for the agent provider's website or relevant documentation.
	URL string
}

// Represents a distinct capability or function that an agent can perform.
type AgentSkill struct {
	// A detailed description of the skill, intended to help clients or users
	// understand its purpose and functionality.
	Description string

	// Example prompts or scenarios that this skill can handle. Provides a hint to
	// the client on how to use the skill.
	Examples []string

	// A unique identifier for the agent's skill.
	ID string

	// The set of supported input MIME types for this skill, overriding the agent's
	// defaults.
	InputModes []string

	// A human-readable name for the skill.
	Name string

	// The set of supported output MIME types for this skill, overriding the agent's
	// defaults.
	OutputModes []string

	// Security schemes necessary for the agent to leverage this skill.
	// As in the overall AgentCard.security, this list represents a logical OR of
	// security
	// requirement objects. Each object is a set of security schemes that must be used
	// together
	// (a logical AND).
	Security []map[string][]string

	// A set of keywords describing the skill's capabilities.
	Tags []string
}

// Represents a file, data structure, or other resource generated by an agent
// during a task.
type Artifact struct {
	// A unique identifier for the artifact within the scope of the task.
	ArtifactID string

	// An optional, human-readable description of the artifact.
	Description *string

	// The URIs of extensions that are relevant to this artifact.
	Extensions []string

	// Optional metadata for extensions. The key is an extension-specific identifier.
	Metadata map[string]any

	// An optional, human-readable name for the artifact.
	Name *string

	// An array of content parts that make up the artifact.
	Parts []ArtifactPartsElem
}
type ArtifactPartsElem interface{}

// An A2A-specific error indicating that the agent does not have an Authenticated
// Extended Card configured
type AuthenticatedExtendedCardNotConfiguredError struct {
	// The error code for when an authenticated extended card is not configured.
	Code int

	// A primitive or structured value containing additional information about the
	// error.
	// This may be omitted.
	Data any

	// The error message.
	Message string
}

// Defines configuration details for the OAuth 2.0 Authorization Code flow.
type AuthzCodeOAuthFlow struct {
	// The authorization URL to be used for this flow.
	// This MUST be a URL and use TLS.
	AuthzURL string

	// The URL to be used for obtaining refresh tokens.
	// This MUST be a URL and use TLS.
	RefreshURL *string

	// The available scopes for the OAuth2 security scheme. A map between the scope
	// name and a short description for it.
	Scopes map[string]string

	// The token URL to be used for this flow.
	// This MUST be a URL and use TLS.
	TokenURL string
}

// Represents a request for the `tasks/cancel` method.
type CancelTaskRequest struct {
	// The parameters identifying the task to cancel.
	Params TaskIDParams
}

// Represents a successful response for the `tasks/cancel` method.
type CancelTaskResponse struct {
	// The result, containing the final state of the canceled Task object.
	Result Task
}

// Defines configuration details for the OAuth 2.0 Client Credentials flow.
type ClientCredentialsOAuthFlow struct {
	// The URL to be used for obtaining refresh tokens. This MUST be a URL.
	RefreshURL *string

	// The available scopes for the OAuth2 security scheme. A map between the scope
	// name and a short description for it.
	Scopes map[string]string

	// The token URL to be used for this flow. This MUST be a URL.
	TokenURL string
}

// Represents a structured data segment (e.g., JSON) within a message or artifact.
type DataPart struct {
	// The structured data content.
	Data map[string]any

	// The type of this part, used as a discriminator. Always 'data'.
	Kind string

	// Optional metadata associated with this part.
	Metadata map[string]any
}

// Defines parameters for deleting a specific push notification configuration for a
// task.
type DeleteTaskPushNotificationConfigParams struct {
	// The unique identifier of the task.
	ID string

	// Optional metadata associated with the request.
	Metadata map[string]any

	// The ID of the push notification configuration to delete.
	PushNotificationConfigID string
}

// Represents a request for the `tasks/pushNotificationConfig/delete`
// method.
type DeleteTaskPushNotificationConfigRequest struct {
	// The parameters identifying the push notification configuration to delete.
	Params DeleteTaskPushNotificationConfigParams
}

// Represents a successful response for the `tasks/pushNotificationConfig/delete` method.
type DeleteTaskPushNotificationConfigResponse struct {
	// The result is null on successful deletion.
	Result any
}

// Defines base properties for a file.
type FileBase struct {
	// The MIME type of the file (e.g., "application/pdf").
	MimeType *string

	// An optional name for the file (e.g., "document.pdf").
	Name *string
}

// Represents a file segment within a message or artifact. The file content can be
// provided either directly as bytes or as a URI.
type FilePart struct {
	// The file content, represented as either a URI or as base64-encoded bytes.
	File FilePartFile

	// The type of this part, used as a discriminator. Always 'file'.
	Kind string

	// Optional metadata associated with this part.
	Metadata map[string]any
}

// Represents a file with its content provided directly as a base64-encoded string.
type FilePartFile struct {
	// The base64-encoded content of the file.
	Bytes string

	// The MIME type of the file (e.g., "application/pdf").
	MimeType *string

	// An optional name for the file (e.g., "document.pdf").
	Name *string

	// A URL pointing to the file's content.
	Uri string
}

// Represents a file with its content provided directly as a base64-encoded string.
type FileWithBytes struct {
	// The base64-encoded content of the file.
	Bytes string

	// The MIME type of the file (e.g., "application/pdf").
	MimeType *string

	// An optional name for the file (e.g., "document.pdf").
	Name *string
}

// Represents a file with its content located at a specific URI.
type FileWithUri struct {
	// The MIME type of the file (e.g., "application/pdf").
	MimeType *string

	// An optional name for the file (e.g., "document.pdf").
	Name *string

	// A URL pointing to the file's content.
	Uri string
}

// Represents a successful response for the `agent/getAuthenticatedExtendedCard` method.
type GetAuthenticatedExtendedCardResponse struct {
	// The result is an Agent Card object.
	Result AgentCard
}

// Defines parameters for fetching a specific push notification configuration for a task.
type GetTaskPushNotificationConfigParams struct {
	// The unique identifier of the task.
	ID string

	// Optional metadata associated with the request.
	Metadata map[string]any

	// The ID of the push notification configuration to retrieve.
	PushNotificationConfigID *string
}

// Represents a request for the `tasks/pushNotificationConfig/get` method.
type GetTaskPushNotificationConfigRequest struct {
	// The parameters for getting a push notification configuration.
	Params GetTaskPushNotificationConfigRequestParams
}

// Defines parameters containing a task ID, used for simple task operations.
type GetTaskPushNotificationConfigRequestParams struct {
	// The unique identifier of the task.
	ID string

	// Optional metadata associated with the request.
	Metadata map[string]any

	// The ID of the push notification configuration to retrieve.
	PushNotificationConfigID *string
}

// Represents a response for the `tasks/pushNotificationConfig/get` method.
type GetTaskPushNotificationConfigResponse struct {
	// The result, containing the requested push notification configuration.
	Result TaskPushConfig
}

// Represents a request for the `tasks/get` method.
type GetTaskRequest struct {
	// The parameters for querying a task.
	Params TaskQueryParams
}

// Represents a successful response for the `tasks/get` method.
type GetTaskResponse struct {
	// The result, containing the requested Task object.
	Result Task
}

// Defines a security scheme using HTTP authentication.
type HTTPAuthSecurityScheme struct {
	// A hint to the client to identify how the bearer token is formatted (e.g.,
	// "JWT").
	// This is primarily for documentation purposes.
	BearerFormat *string

	// An optional description for the security scheme.
	Description *string

	// The name of the HTTP Authentication scheme to be used in the Authorization
	// header,
	// as defined in RFC7235 (e.g., "Bearer").
	// This value should be registered in the IANA Authentication Scheme registry.
	Scheme string

	// The type of the security scheme. Must be 'http'.
	Type string
}

// Defines configuration details for the OAuth 2.0 Implicit flow.
type ImplicitOAuthFlow struct {
	// The authorization URL to be used for this flow. This MUST be a URL.
	AuthzURL string

	// The URL to be used for obtaining refresh tokens. This MUST be a URL.
	RefreshURL *string

	// The available scopes for the OAuth2 security scheme. A map between the scope
	// name and a short description for it.
	Scopes map[string]string
}

// An error indicating an internal error on the server.
type InternalError struct {
	// The error code for an internal server error.
	Code int

	// A primitive or structured value containing additional information about the
	// error.
	// This may be omitted.
	Data any

	// The error message.
	Message string
}

// Defines parameters for listing all push notification configurations associated
// with a task.
type ListTaskPushNotificationConfigParams struct {
	// The unique identifier of the task.
	ID string

	// Optional metadata associated with the request.
	Metadata map[string]any
}

// Represents a request for the `tasks/pushNotificationConfig/list`
// method.
type ListTaskPushNotificationConfigRequest struct {
	// The parameters identifying the task whose configurations are to be listed.
	Params ListTaskPushNotificationConfigParams
}

// Represents a response for the `tasks/pushNotificationConfig/list` method.
type ListTaskPushNotificationConfigResponse struct {
	// The result, containing an array of all push notification configurations for the
	// task.
	Result []TaskPushConfig
}

// Represents a single message in the conversation between a user and an agent.
type Message struct {
	// The context identifier for this message, used to group related interactions.
	ContextID *string

	// The URIs of extensions that are relevant to this message.
	Extensions []string

	// The type of this object, used as a discriminator. Always 'message' for a
	// Message.
	Kind string

	// A unique identifier for the message, typically a UUID, generated by the sender.
	MessageID string

	// Optional metadata for extensions. The key is an extension-specific identifier.
	Metadata map[string]any

	// An array of content parts that form the message body. A message can be
	// composed of multiple parts of different types (e.g., text and files).
	Parts []MessagePartsElem

	// A list of other task IDs that this message references for additional context.
	ReferenceTasks []TaskID

	// Identifies the sender of the message.
	// service.
	Role MessageRole

	// The identifier of the task this message is part of. Can be omitted for the
	// first message of a new task.
	TaskID *TaskID
}
type MessagePartsElem interface{}

type MessageRole string

const (
	MessageRoleAgent MessageRole = "agent"
	MessageRoleUser  MessageRole = "user"
)

// Defines configuration options for a `message/send` or `message/stream` request.
type MessageSendConfig struct {
	// A list of output MIME types the client is prepared to accept in the response.
	AcceptedOutputModes []string

	// If true, the client will wait for the task to complete. The server may reject
	// this if the task is long-running.
	Blocking *bool

	// The number of most recent messages from the task's history to retrieve in the
	// response.
	HistoryLength *int

	// Configuration for the agent to send push notifications for updates after the
	// initial response.
	PushConfig *PushConfig
}

// Defines the parameters for a request to send a message to an agent. This can be
// used
// to create a new task, continue an existing one, or restart a task.
type MessageSendParams struct {
	// Optional configuration for the send request.
	Configuration *MessageSendConfig

	// The message object being sent to the agent.
	Message Message

	// Optional metadata for extensions.
	Metadata map[string]any
}

// Defines a security scheme using mTLS authentication.
type MutualTLSSecurityScheme struct {
	// An optional description for the security scheme.
	Description *string

	// The type of the security scheme. Must be 'mutualTLS'.
	Type string
}

// Defines a security scheme using OAuth 2.0.
type OAuth2SecurityScheme struct {
	// An optional description for the security scheme.
	Description *string

	// An object containing configuration information for the supported OAuth 2.0
	// flows.
	Flows OAuthFlows

	// URL to the oauth2 authorization server metadata
	// [RFC8414](https://datatracker.ietf.org/doc/html/rfc8414). TLS is required.
	Oauth2MetadataURL *string

	// The type of the security scheme. Must be 'oauth2'.
	Type string
}

// Defines the configuration for the supported OAuth 2.0 flows.
type OAuthFlows struct {
	// Configuration for the OAuth Authorization Code flow. Previously called
	// accessCode in OpenAPI 2.0.
	AuthzCode *AuthzCodeOAuthFlow

	// Configuration for the OAuth Client Credentials flow. Previously called
	// application in OpenAPI 2.0.
	ClientCredentials *ClientCredentialsOAuthFlow

	// Configuration for the OAuth Implicit flow.
	Implicit *ImplicitOAuthFlow

	// Configuration for the OAuth Resource Owner Password flow.
	Password *PasswordOAuthFlow
}

// Defines a security scheme using OpenID Connect.
type OpenIDConnectSecurityScheme struct {
	// An optional description for the security scheme.
	Description *string

	// The OpenID Connect Discovery URL for the OIDC provider's metadata.
	OpenIDConnectURL string

	// The type of the security scheme. Must be 'openIDConnect'.
	Type string
}

// A discriminated union representing a part of a message or artifact, which can
// be text, a file, or structured data.
type Part interface{}

// Defines base properties common to all message or artifact parts.
type PartBase struct {
	// Optional metadata associated with this part.
	Metadata map[string]any
}

// Defines configuration details for the OAuth 2.0 Resource Owner Password flow.
type PasswordOAuthFlow struct {
	// The URL to be used for obtaining refresh tokens. This MUST be a URL.
	RefreshURL *string

	// The available scopes for the OAuth2 security scheme. A map between the scope
	// name and a short description for it.
	Scopes map[string]string

	// The token URL to be used for this flow. This MUST be a URL.
	TokenURL string
}

// Defines authentication details for a push notification endpoint.
type PushAuthInfo struct {
	// Optional credentials required by the push notification endpoint.
	Credentials *string

	// A list of supported authentication schemes (e.g., 'Basic', 'Bearer').
	Schemes []string
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

// Defines a security scheme that can be used to secure an agent's endpoints.
// This is a discriminated union type based on the OpenAPI 3.0 Security Scheme
// Object.
type SecurityScheme interface{}

// Defines base properties shared by all security scheme objects.
type SecuritySchemeBase struct {
	// An optional description for the security scheme.
	Description *string
}

// Represents a request for the `message/send` method.
type SendMessageRequest struct {
	// The parameters for sending a message.
	Params MessageSendParams
}

// Represents a response for the `message/send` method.
type SendMessageResponse interface {
	isSendMessageResponse()
}

func (Task) isSendMessageResponse()    {}
func (Message) isSendMessageResponse() {}

// Represents a request for the `message/stream` method.
type SendStreamingMessageRequest struct {
	// The parameters for sending a message.
	Params MessageSendParams
}

// Represents a request for the `tasks/pushNotificationConfig/set` method.
type SetTaskPushNotificationConfigParams struct {
	// The parameters for setting the push notification configuration.
	Params TaskPushConfig
}

// Represents a response for the `tasks/pushNotificationConfig/set` method.
type SetTaskPushNotificationConfigResponse struct {
	// The result, containing the configured push notification settings.
	Result TaskPushConfig
}

type TaskID string

// Represents a single, stateful operation or conversation between a client and an
// agent.
type Task struct {
	// A collection of artifacts generated by the agent during the execution of the
	// task.
	Artifacts []Artifact

	// A server-generated identifier for maintaining context across multiple related
	// tasks or interactions.
	ContextID string

	// An array of messages exchanged during the task, representing the conversation
	// history.
	History []Message

	// A unique identifier for the task, generated by the server for a new task.
	ID TaskID

	// The type of this object, used as a discriminator. Always 'task' for a Task.
	Kind string

	// Optional metadata for extensions. The key is an extension-specific identifier.
	Metadata map[string]any

	// The current status of the task, including its state and a descriptive message.
	Status TaskStatus
}

// An event sent by the agent to notify the client that an artifact has been
// generated or updated. This is typically used in streaming models.
type TaskArtifactUpdateEvent struct {
	// If true, the content of this artifact should be appended to a previously sent
	// artifact with the same ID.
	Append *bool

	// The artifact that was generated or updated.
	Artifact Artifact

	// The context ID associated with the task.
	ContextID string

	// The type of this event, used as a discriminator. Always 'artifact-update'.
	Kind string

	// If true, this is the final chunk of the artifact.
	LastChunk *bool

	// Optional metadata for extensions.
	Metadata map[string]any

	// The ID of the task this artifact belongs to.
	TaskID TaskID
}

// Defines parameters containing a task ID, used for simple task operations.
type TaskIDParams struct {
	// The unique identifier of the task.
	ID TaskID

	// Optional metadata associated with the request.
	Metadata map[string]any
}

// A container associating a push notification configuration with a specific task.
type TaskPushConfig struct {
	// The push notification configuration for this task.
	PushConfig PushConfig

	// The ID of the task.
	TaskID TaskID
}

// Defines parameters for querying a task, with an option to limit history length.
type TaskQueryParams struct {
	// The number of most recent messages from the task's history to retrieve.
	HistoryLength *int

	// The unique identifier of the task.
	ID TaskID

	// Optional metadata associated with the request.
	Metadata map[string]any
}

// Represents a request for the `tasks/resubscribe` method, used to resume
// a streaming connection.
type TaskResubscriptionRequest struct {
	// The parameters identifying the task to resubscribe to.
	Params TaskIDParams
}

type TaskState string

const (
	TaskStateAuthRequired  TaskState = "auth-required"
	TaskStateCanceled      TaskState = "canceled"
	TaskStateCompleted     TaskState = "completed"
	TaskStateFailed        TaskState = "failed"
	TaskStateInputRequired TaskState = "input-required"
	TaskStateRejected      TaskState = "rejected"
	TaskStateSubmitted     TaskState = "submitted"
	TaskStateUnknown       TaskState = "unknown"
	TaskStateWorking       TaskState = "working"
)

// Represents the status of a task at a specific point in time.
type TaskStatus struct {
	// An optional, human-readable message providing more details about the current
	// status.
	Message *Message

	// The current state of the task's lifecycle.
	State TaskState

	// An ISO 8601 datetime string indicating when this status was recorded.
	Timestamp *string
}

// Event interface is used to represent types that can be sent over a streaming connection.
type Event interface {
	isEvent()
}

func (t Task) isEvent()                    {}
func (t Message) isEvent()                 {}
func (t TaskStatusUpdateEvent) isEvent()   {}
func (t TaskArtifactUpdateEvent) isEvent() {}

// An event sent by the agent to notify the client of a change in a task's status.
// This is typically used in streaming or subscription models.
type TaskStatusUpdateEvent struct {
	// The context ID associated with the task.
	ContextID string

	// If true, this is the final event in the stream for this interaction.
	Final bool

	// The type of this event, used as a discriminator. Always 'status-update'.
	Kind string

	// Optional metadata for extensions.
	Metadata map[string]any

	// The new status of the task.
	Status TaskStatus

	// The ID of the task that was updated.
	TaskID TaskID
}

// Represents a text segment within a message or artifact.
type TextPart struct {
	// The type of this part, used as a discriminator. Always 'text'.
	Kind string

	// Optional metadata associated with this part.
	Metadata map[string]any

	// The string content of the text part.
	Text string
}

type TransportProtocol string

const (
	TransportProtocolJSONRPC  TransportProtocol = "JSONRPC"
	TransportProtocolGRPC     TransportProtocol = "GRPC"
	TransportProtocolHTTPJSON TransportProtocol = "HTTP+JSON"
)
