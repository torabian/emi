package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/urfave/cli/v3"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/torabian/emi/examples/fullstack/emigo"
	unk "github.com/torabian/emi/examples/fullstack/sdk"
	"github.com/torabian/emi/lib/core"
)

func main() {

	app := &cli.Command{
		Name:  "compute-api",
		Usage: "vector compute server",
		Commands: []*cli.Command{
			{
				Name:  "start",
				Usage: "start http server",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "addr",
						Value: ":8080",
						Usage: "server listen address",
					},
				},
				Action: func(ctx context.Context, cmd *cli.Command) error {
					return runServer(cmd.String("addr"))
				},
			},
			{
				Name:  "cast-cvc",
				Usage: "Casts common vector compute dto from cli",
				Flags: CastEmiFlagToUrfave(unk.GetCommonVectorComputeDtoCliFlags("")),
				Action: func(ctx context.Context, cmd *cli.Command) error {
					data := unk.CastCommonVectorComputeDtoFromCli(cmd)
					fmt.Println("Nullable value:", data.Json())
					return nil
				},
			},
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		panic(err)
	}
}

func firstNonEmpty(a, b string) string {
	if a != "" {
		return a
	}
	return b
}

func CastEmiFlagToUrfave(flags []emigo.CliFlag) []cli.Flag {
	var out []cli.Flag

	for _, f := range flags {
		// Recursively flatten children
		if len(f.Children) > 0 {
			out = append(out, CastEmiFlagToUrfave(f.Children)...)
			continue
		}

		usage := firstNonEmpty(f.Description, f.Usage)
		req := f.Required

		switch {
		case strings.Contains(f.Type, "int") && !core.IsNullable(f.Type):
			out = append(out, &cli.Int64Flag{Name: f.Name, Usage: usage, Required: req})
		case strings.Contains(f.Type, "bool") && !core.IsNullable(f.Type):
			out = append(out, &cli.BoolFlag{Name: f.Name, Usage: usage, Required: req})
		case strings.Contains(f.Type, "float") && !core.IsNullable(f.Type):
			out = append(out, &cli.Float64Flag{Name: f.Name, Usage: usage, Required: req})
		default:
			out = append(out, &cli.StringFlag{Name: f.Name, Usage: usage, Required: req})
		}
	}

	return out
}
func runServer(addr string) error {
	r := gin.Default()

	// ----------- HTTP -----------
	unk.ComputeApiAction(r, func(req unk.ComputeApiActionRequest, c *gin.Context) (*unk.ComputeApiActionResponse, error) {
		output := sumVectors(req.Body.InitialVector1, req.Body.InitialVector2)

		return &unk.ComputeApiActionResponse{
			StatusCode: http.StatusOK,
			Payload: unk.ComputeApiActionRes{
				OutputVector: output,
			},
		}, nil
	})

	// ----------- SSE -----------
	unk.ComputeApiSseAction(r, func(req unk.ComputeApiSseActionRequest, g *gin.Context) (*unk.ComputeApiSseActionResponse, error) {
		output := sumVectors(req.Body.InitialVector1, req.Body.InitialVector2)

		g.Writer.Header().Set("Content-Type", "text/event-stream")
		g.Writer.Header().Set("Cache-Control", "no-cache")
		g.Writer.Header().Set("Connection", "keep-alive")

		for i := 0; i < 10; i++ {
			fmt.Fprintf(g.Writer, "data: %v\n\n", output)
			g.Writer.Flush()
			time.Sleep(500 * time.Millisecond)
		}
		return nil, nil
	})

	unk.ComputeApiSseChannelAction(r, func(req unk.ComputeApiSseChannelActionRequest, g *gin.Context) (*unk.ComputeApiSseChannelActionResponse, error) {
		ch := computeViaChannel(req.Body.InitialVector1, req.Body.InitialVector2)
		SSEStream(g, ch)
		return nil, nil
	})

	// ----------- WebSocket (simple) -----------
	unk.ComputeReactiveNoPathAction(r, func(msg unk.ComputeReactiveNoPathActionMessage) error {
		var req unk.CommonVectorComputeDto
		_ = json.Unmarshal(msg.Raw, &req)

		output := sumVectors(req.InitialVector1, req.InitialVector2)

		resBytes, _ := json.Marshal(unk.CommonVectorResponseDto{
			OutputVector: output,
		})

		return msg.Conn.WriteMessage(websocket.TextMessage, resBytes)
	})

	// ----------- WebSocket (duplex channel) -----------
	unk.ComputeReactiveActionDuplex(r, func(ctx *unk.ComputeReactiveActionSession) {

		fmt.Println(ctx.PathParams.Age + ctx.PathParams.Id)

		ctx.Out <- unk.ComputeReactiveActionMessage{
			MessageType: websocket.TextMessage,
			Raw:         []byte(fmt.Sprintf("Query Param 1: %v", ctx.QueryParams.QueryParam1)),
		}

		for {
			select {
			case msg, ok := <-ctx.In:
				if !ok {
					return
				}
				ctx.Out <- unk.ComputeReactiveActionMessage{
					MessageType: websocket.TextMessage,
					Raw:         msg.Raw,
				}
			case <-ctx.Done:
				return
			}
		}
	})

	fmt.Println("Server running on", addr)
	return r.Run(addr)
}

// small reusable core logic (cleaner long-term)
func sumVectors(v1, v2 []int) []int {
	minLen := len(v1)
	if len(v2) < minLen {
		minLen = len(v2)
	}

	out := make([]int, minLen)
	for i := 0; i < minLen; i++ {
		out[i] = v1[i] + v2[i]
	}
	return out
}

func computeViaChannel(vec1 []int, vec2 []int) chan interface{} {

	minLen := len(vec1)
	if len(vec2) < minLen {
		minLen = len(vec2)
	}

	output := make([]int, minLen)
	for i := 0; i < minLen; i++ {
		output[i] = vec1[i] + vec2[i]
	}

	ch := make(chan interface{})
	go func() {
		defer close(ch)
		for i := 0; i < 10; i++ {
			ch <- output
			time.Sleep(500 * time.Millisecond)
		}
	}()

	return ch
}

func SSEStream(c *gin.Context, ch <-chan interface{}) {
	c.Writer.Header().Set("Content-Type", "text/event-stream")
	c.Writer.Header().Set("Cache-Control", "no-cache")
	c.Writer.Header().Set("Connection", "keep-alive")

	for val := range ch {
		fmt.Fprintf(c.Writer, "data: %v\n\n", val)
		c.Writer.Flush()
	}
}
