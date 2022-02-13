package controllers

import (
	http "net/http"

	gin "github.com/gin-gonic/gin"

	Services "NoteKeeperAPI/src/services"
)

type NotesController struct {
	Service Services.NotesService
}

func (v NotesController) GetNotes(c *gin.Context) {
	c.JSON(http.StatusOK, v.Service.GetNotes())
}
