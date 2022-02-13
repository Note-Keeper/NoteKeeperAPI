package controllers

import (
	"fmt"
	http "net/http"

	gin "github.com/gin-gonic/gin"

	Services "NoteKeeperAPI/src/services"
	Requests "NoteKeeperAPI/src/types/requests"
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

	User := v.Service.Register(req)
	fmt.Println(User)

	c.JSON(http.StatusOK, req)
}
