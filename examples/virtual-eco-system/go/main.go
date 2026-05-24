package main

import (
	"fmt"
	unknownpackage "test/emi/sdkgen"
)

func main() {

	query, err := unknownpackage.PrepareInsertUserVsql(unknownpackage.InsertUserVsqlParams{
		Name: "Ali",
	})

	fmt.Println(query, err)

	// r := gin.Default()
	// Register the endpoint
	// unknownpackage.PatchGiantAction(r, func(c unknownpackage.PatchGiantActionRequest, gin *gin.Context) (*unknownpackage.PatchGiantActionResponse, error) {
	// 	return &unknownpackage.PatchGiantActionResponse{
	// 		Payload: &unknownpackage.GiantDto{
	// 			FirstName: "asdasd",
	// 		},
	// 	}, nil
	// })

	// unknownpackage.GetAsGiantsActionGin(r, func(c unknownpackage.GetAsGiantsActionRequest, gin *gin.Context) (*unknownpackage.GetAsGiantsActionResponse, error) {
	// 	return &unknownpackage.GetAsGiantsActionResponse{
	// 		Payload: c.Params.Id,
	// 	}, nil
	// })

	// fmt.Println("Server running on :8080")
	// r.Run(":8080") // blocks here
}
