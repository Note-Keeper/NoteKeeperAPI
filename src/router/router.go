package router

import (
	cors "github.com/gin-contrib/cors"
	gin "github.com/gin-gonic/gin"

	Controllers "NoteKeeperAPI/src/controllers"
)

func CreateServer() *gin.Engine {
	R := gin.Default()

	R.Use(cors.Default())
	R.SetTrustedProxies([]string{"192.168.1.2"})

	R.GET("/notes", Controllers.GetNotes)

	return R
}
