package a2asrv

import "github.com/a2aproject/a2a-go/a2a"

type RequestContext struct {
	Request      a2a.MessageSendParams
	Task         *a2a.Task
	RelatedTasks []a2a.Task
	ContextID    string
}
