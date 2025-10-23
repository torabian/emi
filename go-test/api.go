package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// --- Request / Response structs ---
type GetUserRequest struct {
	ID            string `uri:"id" binding:"required"`
	Verbose       bool   `form:"verbose"`
	Authorization string `header:"Authorization"`
}

type GetUserResponse struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// --- Service interface ---
type GetUserService interface {
	Handle(req GetUserRequest) (GetUserResponse, error)
}

// Optional raw handling
type GetUserRawService interface {
	HandleRaw(req GetUserRequest, c *gin.Context) bool
}

func GetUserHandler(service interface{}) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req GetUserRequest
		if err := c.ShouldBindUri(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := c.ShouldBindQuery(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := c.ShouldBindHeader(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Check if raw service is implemented
		if rawService, ok := service.(GetUserRawService); ok {
			if handled := rawService.HandleRaw(req, c); handled {
				return
			}
		}

		// Check if typed service is implemented
		if typedService, ok := service.(GetUserService); ok {
			resp, err := typedService.Handle(req)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, resp)
			return
		}

		// If neither implemented, return 501
		c.JSON(http.StatusNotImplemented, gin.H{"error": "handler not implemented"})
	}
}
