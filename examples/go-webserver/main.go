package main

import (
	"fmt"
	"time"

	unknownpackage "example.com/go-webserver/lib"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	unknownpackage.SampleSseAction(r, func(ssar unknownpackage.SampleSseActionRequest) (*unknownpackage.SampleSseActionResponse, error) {
		c := ssar.C

		return &unknownpackage.SampleSseActionResponse{
			Payload: make(chan string),
		}, nil

		// Set SSE headers
		c.Writer.Header().Set("Content-Type", "text/event-stream")
		c.Writer.Header().Set("Cache-Control", "no-cache")
		c.Writer.Header().Set("Connection", "keep-alive")

		// Send 10 SSE messages
		for i := 1; i <= 10; i++ {
			fmt.Fprintf(c.Writer, "data: Message %d\n\n", i)
			c.Writer.Flush() // Important: flush after each message
			time.Sleep(time.Second)
		}

		// Return nil to skip JSON handling
		return nil, nil

	})
	r.Run(":8080")
}
