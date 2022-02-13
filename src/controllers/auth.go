package controllers

import (
	http "net/http"

	gin "github.com/gin-gonic/gin"

	Helpers "NoteKeeperAPI/src/helpers"
	Services "NoteKeeperAPI/src/services"
	Requests "NoteKeeperAPI/src/types/requests"
)

type AuthController struct {
	Service Services.AuthService
}

func (v AuthController) Register(c *gin.Context) {
	var req Requests.Register

	err := c.ShouldBindJSON(&req)
	Helpers.PrintError(err)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	if !v.Service.ValidateRegister(req) {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "existing_data",
		})

		return
	}

	User := v.Service.Register(req)
	c.JSON(http.StatusOK, User)

}
