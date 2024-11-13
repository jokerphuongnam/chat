package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (handler *Handler) FindUsersByNameHandler(c *gin.Context) {
	// Get name from the query parameters.
	name := c.Query("name")
	if name == "" {
		c.JSON(http.StatusBadRequest, ResponseMessage{
			Code:    http.StatusBadRequest,
			Message: "name parameter is required",
		})
		return
	}

	// Call FindUsersByName function
	users, err := handler.Database.FindUsersByNameHandler(name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseMessage{
			Code:    http.StatusInternalServerError,
			Message: "failed to find users",
		})
		return
	}

	// Success response with users data
	c.JSON(http.StatusOK, ResponseMessage{
		Code:    http.StatusOK,
		Message: "success",
		Data:    users,
	})
}
