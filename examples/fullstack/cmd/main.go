package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	unk "github.com/torabian/emi/examples/fullstack/sdk"
)

func main() {

	r := gin.Default()

	// Register ComputeApiAction
	unk.ComputeApiAction(r, func(req unk.ComputeApiActionRequest, c *gin.Context) (*unk.ComputeApiActionResponse, error) {
		// Example business logic: sum corresponding elements of InitialVector1 and InitialVector2
		vec1 := req.Body.InitialVector1
		vec2 := req.Body.InitialVector2

		minLen := len(vec1)
		if len(vec2) < minLen {
			minLen = len(vec2)
		}

		output := make([]int, minLen)
		for i := 0; i < minLen; i++ {
			output[i] = vec1[i] + vec2[i]
		}

		return &unk.ComputeApiActionResponse{
			StatusCode: http.StatusOK,
			Payload: unk.ComputeApiActionRes{
				OutputVector: output,
			},
		}, nil
	})

	unk.ComputeApiSseAction(r, func(req unk.ComputeApiSseActionRequest, gin *gin.Context) (*unk.ComputeApiSseActionResponse, error) {
		vec1 := req.Body.InitialVector1
		vec2 := req.Body.InitialVector2

		minLen := len(vec1)
		if len(vec2) < minLen {
			minLen = len(vec2)
		}

		output := make([]int, minLen)
		for i := 0; i < minLen; i++ {
			output[i] = vec1[i] + vec2[i]
		}

		/// This is the very simple implementation of it, for the gin only. No extra code.
		gin.Writer.Header().Set("Content-Type", "text/event-stream")
		gin.Writer.Header().Set("Cache-Control", "no-cache")
		gin.Writer.Header().Set("Connection", "keep-alive")

		for i := 0; i < 10; i++ {
			fmt.Fprintf(gin.Writer, "data: %v\n\n", output)
			gin.Writer.Flush()
			time.Sleep(500 * time.Millisecond)
		}

		return nil, nil
	})

	unk.ComputeApiSseChannelAction(r, func(req unk.ComputeApiSseChannelActionRequest, gin *gin.Context) (*unk.ComputeApiSseChannelActionResponse, error) {
		ch := computeViaChannel(req.Body.InitialVector1, req.Body.InitialVector2)
		SSEStream(gin, ch)
		return nil, nil
	})

	// Using gin loop instead of a channel, per each message, the func will be called.
	unk.ComputeReactiveNoPathAction(r, func(msg unk.ComputeReactiveNoPathActionMessage) error {
		var req unk.CommonVectorComputeDto
		json.Unmarshal(msg.Raw, &req)

		// Simple computation example
		minLen := len(req.InitialVector1)
		if len(req.InitialVector2) < minLen {
			minLen = len(req.InitialVector2)
		}
		output := make([]int, minLen)
		for i := 0; i < minLen; i++ {
			output[i] = req.InitialVector1[i] + req.InitialVector2[i]
		}

		resBytes, _ := json.MarshalIndent(unk.CommonVectorResponseDto{OutputVector: output}, "", "  ")
		return msg.Conn.WriteMessage(websocket.TextMessage, resBytes)
	})

	// Channel oriented action
	unk.ComputeReactiveActionDuplex(r, func(ctx *unk.ComputeReactiveActionSession) {

		fmt.Println(ctx.PathParams.Age + ctx.PathParams.Id)
		ctx.Out <- unk.ComputeReactiveActionMessage{
			MessageType: websocket.TextMessage,
			Raw:         []byte("Query Param 1:" + fmt.Sprintf("%v", ctx.QueryParams.QueryParam1)),
		}
		for {
			select {
			case msg, ok := <-ctx.In:
				fmt.Println("Message, ok", msg, ok)
				if !ok {
					return // client disconnected
				}

				ctx.Out <- unk.ComputeReactiveActionMessage{
					MessageType: websocket.TextMessage,
					Raw:         msg.Raw,
				}

			case <-ctx.Done:
				fmt.Println("Completed")
				return
			}
		}
	})

	r.Run(":8080") // listen and serve on 0.0.0.0:8080
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
