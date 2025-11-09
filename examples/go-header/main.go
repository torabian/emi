package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	unknownpackage "test.com/emi-go-header-demo/generalgen"
	"test.com/emi-go-header-demo/generalgen/emigo"
)

func main() {
	r := gin.Default()

	// Register the endpoint
	unknownpackage.ComputeApiAction(r, func(c unknownpackage.ComputeApiActionRequest, ginCtx *gin.Context) (*unknownpackage.ComputeApiActionResponse, error) {
		// Set a cookie as example
		ginCtx.SetCookie("name", "Hossein", 100, "/", "", true, true)

		return &unknownpackage.ComputeApiActionResponse{
			StatusCode: 200,
			Payload:    map[string]string{"msg": "hello from server"},
		}, nil
	})

	// Start client in a goroutine, after a short delay
	go func() {
		time.Sleep(500 * time.Millisecond) // give server time to start

		res, err := unknownpackage.ComputeApiActionCall(unknownpackage.ComputeApiActionRequest{
			Body: unknownpackage.ComputeApiActionReq{
				InitialVector1: []int{1, 2, 3},
				InitialVector2: []int{4, 5, 6},
			},
		}, &emigo.APIClient{
			BaseURL:    "http://localhost:8080",
			Headers:    http.Header{"Authorization": {"Bearer token"}},
			HTTPClient: &http.Client{},
		})

		if err != nil {
			fmt.Println("Client error:", err)
		} else {
			fmt.Println("Client result:", res.Payload)
		}
	}()

	fmt.Println("Server running on :8080")
	r.Run(":8080") // blocks here
}
