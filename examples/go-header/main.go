package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	fmt.Println("Server running on :8080")
	r.Run(":8080")
}
