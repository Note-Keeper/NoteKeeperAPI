package router

import (
	cors "github.com/gin-contrib/cors"
	gin "github.com/gin-gonic/gin"

	Controllers "NoteKeeperAPI/src/controllers"
	Database "NoteKeeperAPI/src/database"
	Middlewares "NoteKeeperAPI/src/middlewares"
)

func CreateServer() *gin.Engine {
	R := gin.Default()

	Database.ConnectDB()

	R.Use(cors.Default())
	R.SetTrustedProxies([]string{"192.168.1.2"})

	R.Use(Middlewares.ParseBody)

	AuthRoutes := R.Group("/auth")
	{
		var Controller Controllers.AuthController
		AuthRoutes.POST("/register", Controller.Register)
		AuthRoutes.POST("login", Controller.Login)
		AuthRoutes.POST("logout", Controller.Logout)
	}

	NoteRoutes := R.Group("/notes")
	{
		var Controller Controllers.NotesController
		NoteRoutes.GET("/", Controller.GetNotes)
	}

	return R
}
