package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	unknownpackage "test.com/emi-go-header-demo/generalgen"
)

func main() {
	r := gin.Default()

	unknownpackage.ComputeApiAction(r,
		func(req unknownpackage.ComputeApiActionRequest) (*unknownpackage.ComputeApiActionResponse, error) {

			req.C.JSON(200, "hi")
			return nil, nil
		},
	)

	fmt.Println("Server running on :8080")
	r.Run(":8080")
}
