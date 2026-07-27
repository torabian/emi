package emigo

/**
*	Each result from an action, either can directly access to Gin or Cli
* Context and handle things over there, or can return an EmiAction Result
** Which is standard for a quick result.
**/
type EmiActionResult interface {
	GetStatusCode() int
	GetRespHeaders() map[string]string
	GetPayload() interface{}
}

// Each emi request will implement this
// so user of the generated code can kinda predict
// what are the context features
type EmiRequestContexts interface {
	GetGinCtx() interface{}
	GetCliCtx() interface{}
}
