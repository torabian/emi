package golang

import (
	"bytes"
	"text/template"

	"github.com/torabian/emi/lib/core"
)

func GoActionRenderReactive(
	action core.EmiRpcAction,
	ctx core.MicroGenContext,
	complexes []RecognizedComplex,
) (*core.CodeChunkCompiled, error) {

	realms, deps, err := GoActionReactiveRealms(action, ctx, complexes)
	if err != nil {
		return nil, err
	}

	// For socket, we need some extra dependencies
	deps = append(deps, core.CodeChunkDependency{
		Location: "github.com/gorilla/websocket",
	},

		core.CodeChunkDependency{
			Location: "unicode/utf8",
		})

	res := &core.CodeChunkCompiled{
		Tokens: []core.GeneratedScriptToken{
			{
				Name:  core.TOKEN_ORIGINAL_NAME,
				Value: realms.ActionName,
			},
		},
	}

	const tmpl = `/**
* Action to communicate with the action {{ .realms.ActionName }}
*/

{{ define "upgradeSequence" }}

	{{ if .realms.PathParameter }}
	pathParams := {{ .realms.ActionName }}PathParameterFromGin(c)
	{{ end }}

	ws, err := upgrader{{ .realms.ActionName }}.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot upgrade websocket"})
		return
	}
{{ end }}

func {{ .realms.ActionName }}Meta() struct {
    Name   string
    URL    string
    Method string
	CliName string
	Description string
} {
    return struct {
        Name   string
        URL    string
        Method string
		CliName string
		Description string
    }{
        Name:   "{{ .realms.ActionName }}",
        URL:    "{{ .realms.SafeUrl }}",
        Method: "{{ UPPER .action.Method }}",
		CliName: "{{ .action.CliName }}",
		Description: "{{ .action.Description }}",
    }
}

{{ if .realms.PathParameter }}
	{{ b2s .realms.PathParameter.ActualScript }}
{{ end }}

{{ if .realms.QueryParams }}
	{{ b2s .realms.QueryParams.ActualScript }}
{{ end }}


// WebSocket upgrader
var upgrader{{ .realms.ActionName }} = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}


type {{ .realms.ActionName }}Message struct {
	Raw []byte
	Conn *websocket.Conn	
	MessageType int
	Error error
	{{ if .realms.PathParameter }}
	PathParams {{ .realms.ActionName }}PathParameter
	{{ end}}
}

// Developer handler type
type {{ .realms.ActionName }}Handler func(msg {{ .realms.ActionName }}Message ) error

// Generated handler
func {{ .realms.ActionName }}(r *gin.Engine, handler {{ .realms.ActionName }}Handler) {
	meta := {{ .realms.ActionName }}Meta()
	r.GET(meta.URL, func(c *gin.Context) {
		{{ template "upgradeSequence" . }}

		defer ws.Close()

		for {
			mt, raw, err := ws.ReadMessage()
			msg := {{ .realms.ActionName }}Message{
				Conn: ws,
				Raw: raw,
				Error: err,
				MessageType: mt,
			}

			{{ if .realms.PathParameter }}
			msg.PathParams = pathParams
			{{ end}}

			
			// Provide raw message to developer handler
			if err := handler(msg); err != nil {
				errMsg := fmt.Sprintf("handler error: %v", err)
				if writeErr := ws.WriteMessage(mt, []byte(errMsg)); writeErr != nil {
					break
				}
			}
		}
	})
}

type {{ .realms.ActionName }}Session struct {
	Ctx    *gin.Context
	Socket *websocket.Conn
	Done   chan bool
	Read   chan {{ .realms.ActionName }}ReadChan
	QueryParams {{ .realms.ActionName }}Query
}

type {{ .realms.ActionName }}HandlerDuplex func(*{{ .realms.ActionName }}Session)


type {{ .realms.ActionName }}ReadChan struct {
	Data  []byte
	Error error
}


func {{ .realms.ActionName }}ReactiveHandler(factory func(
	session {{ .realms.ActionName }}Session,
) (chan []byte, error)) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		read := make(chan {{ .realms.ActionName }}ReadChan)
		done := make(chan bool)

		c, err := Upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
		if err != nil {
			c.WriteMessage(websocket.TextMessage, []byte(err.Error()))

			c.Close()
			return
		}

		session := {{ .realms.ActionName }}Session{
			Ctx:    ctx,
			Socket: c,
			Done:   done,
			Read:   read,
		}
		session.QueryParams = {{ .realms.ActionName }}QueryFromGin(ctx)

		write, err := factory(session)

		if err != nil {
			c.WriteMessage(websocket.TextMessage, []byte(err.Error()))
		}

		go func() {
			for {
				_, data, err := c.ReadMessage()
				read <- {{ .realms.ActionName }}ReadChan{
					Data:  data,
					Error: err,
				}

				if err != nil {
					return
				}
			}
		}()

		go func() {
			for {
				select {
				case msg, ok := <-write:
					if !ok {
						// Channel closed; shutdown
						c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
						done <- true
						return
					}
					msgType := websocket.TextMessage
					if !utf8.Valid(msg) {
						msgType = websocket.BinaryMessage
					}
					err := c.WriteMessage(msgType, msg)

					if err != nil {
						// Optionally log the error or send to a logger
						done <- true
						return
					}
				case <-done:
					c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
					return
				}
			}
		}()
	}
}


`

	t := template.Must(template.New("action").Funcs(core.CommonMap).Parse(tmpl))

	var buf bytes.Buffer
	if err := t.Execute(&buf, core.H{
		"action":       action,
		"realms":       realms,
		"shouldExport": true,
	}); err != nil {
		return nil, err
	}

	res.ActualScript = buf.Bytes()
	res.SuggestedFileName = realms.ActionName
	res.SuggestedExtension = ".go"
	res.CodeChunkDependensies = deps

	return res, nil
}
