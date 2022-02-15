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
	UserID := c.Param("id")

	Notes := v.Service.GetNotes(UserID)
	c.JSON(http.StatusOK, Notes)
}

func (v NotesController) GetSingleNote(c *gin.Context) {
	NoteID := c.Param("id")
	c.JSON(http.StatusOK, NoteID)
}

func (v NotesController) CreateNote(c *gin.Context) {}

func (v NotesController) UpdateNote(c *gin.Context) {
	NoteID := c.Param("id")
	c.JSON(http.StatusOK, NoteID)
}

func (v NotesController) DeleteNote(c *gin.Context) {
	NoteID := c.Param("id")
	c.JSON(http.StatusOK, NoteID)
}
