package main

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	// Case 1: JSON missing NullableObject
	json1 := `{
		"min": 1,
		"max": 10,
		"count": 5,
		"staticObject": {"firstName": "Static"}
	}`

	var dto1 SingleDto
	_ = json.Unmarshal([]byte(json1), &dto1)

	fmt.Println(dto1.Max)
	fmt.Println("Case 1 - missing field")
	fmt.Println("NullableObject.IsSet:", dto1.NullableObject.IsSet)
	fmt.Println("NullableObject.Value:", dto1.NullableObject.Value)
	fmt.Println()

	// Case 2: JSON explicitly null
	json2 := `{
		"min": 1,
		"max": 10,
		"count": 5,
		"nullableObject": null,
		"staticObject": {"firstName": "Static"}
	}`

	var dto2 SingleDto
	_ = json.Unmarshal([]byte(json2), &dto2)

	fmt.Println("Case 2 - explicit null")
	fmt.Println("NullableObject.IsSet:", dto2.NullableObject.IsSet)
	fmt.Println("NullableObject.Value:", dto2.NullableObject.Value)
	fmt.Println()

	// Case 3: JSON with actual object
	json3 := `{
		"min": 1,
		"max": 10,
		"count": 5,
		"nullableObject": {"firstName": "John"},
		"staticObject": {"firstName": "Static"}
	}`

	var dto3 SingleDto
	_ = json.Unmarshal([]byte(json3), &dto3)

	fmt.Println("Case 3 - actual object")
	fmt.Println("NullableObject.IsSet:", dto3.NullableObject.IsSet)
	fmt.Println("NullableObject.Value:", dto3.NullableObject.Value)

	r := gin.Default()

	service := &MyGetUserService{}

	// Wire route to generated handler
	r.GET("/users/:id", GetUserHandler(service))

	r.Run(":8080")

}

// Typed response implementation
type MyGetUserService struct{}

// func (s *MyGetUserService) Handle(req GetUserRequest) (GetUserResponse, error) {
// 	fmt.Println("Headers:", req.Authorization)
// 	fmt.Println("Query verbose:", req.Verbose)
// 	fmt.Println("Path ID:", req.ID)

// 	// Example response
// 	return GetUserResponse{
// 		Name: "Alice",
// 		Age:  30,
// 	}, nil
// }

// Optional raw override for special cases
func (s *MyGetUserService) HandleRaw(req GetUserRequest, c *gin.Context) bool {
	if req.ID == "0" {
		// Redirect example
		c.Redirect(302, "/login")
		return true
	}
	return false
}
