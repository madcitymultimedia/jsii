package kernel

import (
	"runtime/debug"
	"github.com/aws/jsii-runtime-go/internal/api"
)

type InvokeProps struct {
	Method     string        `json:"method"`
	Arguments  []interface{} `json:"args"`
	ObjRef     api.ObjectRef `json:"objref"`
	StackTrace string				 `json:"stacktrace"`
}

type StaticInvokeProps struct {
	FQN        api.FQN       `json:"fqn"`
	Method     string        `json:"method"`
	Arguments  []interface{} `json:"args"`
	StackTrace string				 `json:"stacktrace"`
}

type InvokeResponse struct {
	kernelResponse
	Result interface{} `json:"result"`
}

func (c *Client) Invoke(props InvokeProps) (response InvokeResponse, err error) {
	type request struct {
		kernelRequest
		InvokeProps
	}
	props.StackTrace = captureStack()
	err = c.request(request{kernelRequest{"invoke"}, props}, &response)
	return
}

func (c *Client) SInvoke(props StaticInvokeProps) (response InvokeResponse, err error) {
	type request struct {
		kernelRequest
		StaticInvokeProps
	}
	props.StackTrace = captureStack()
	err = c.request(request{kernelRequest{"sinvoke"}, props}, &response)
	return
}

// UnmarshalJSON provides custom unmarshalling implementation for response
// structs. Creating new types is required in order to avoid infinite recursion.
func (r *InvokeResponse) UnmarshalJSON(data []byte) error {
	type response InvokeResponse
	return unmarshalKernelResponse(data, (*response)(r), r)
}

func captureStack() string {
	return string(debug.Stack());
}
