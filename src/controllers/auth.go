package controllers

import (
	"fmt"
	http "net/http"

	gin "github.com/gin-gonic/gin"

	Requests "NoteKeeperAPI/src/controllers/requests"
	Services "NoteKeeperAPI/src/services"
)

type AuthController struct {
	Service Services.AuthService
}

func (v AuthController) Register(c *gin.Context) {
	var req Requests.Register

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Println(req)

	// User := v.Service.Register(UserData)
	c.JSON(http.StatusOK, req)
}
