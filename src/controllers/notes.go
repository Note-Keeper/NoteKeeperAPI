package controllers

import (
	http "net/http"

	gin "github.com/gin-gonic/gin"

	Services "NoteKeeperAPI/src/services"
)

func GetNotes(c *gin.Context) {
	c.JSON(http.StatusOK, Services.GetNotes())
}
