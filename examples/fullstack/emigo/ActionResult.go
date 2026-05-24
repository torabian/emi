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
