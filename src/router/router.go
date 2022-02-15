package router

import (
	cors "github.com/gin-contrib/cors"
	gin "github.com/gin-gonic/gin"

	Controllers "NoteKeeperAPI/src/controllers"
	Database "NoteKeeperAPI/src/database"
)

func CreateServer() *gin.Engine {
	R := gin.Default()

	Database.ConnectDB()

	R.Use(cors.Default())
	R.SetTrustedProxies([]string{"192.168.1.2"})

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

		NoteRoutes.GET("/:id/all", Controller.GetNotes)  // User
		NoteRoutes.GET("/:id", Controller.GetSingleNote) // Note
		NoteRoutes.POST("/", Controller.CreateNote)      // -- None!
		NoteRoutes.PUT("/:id", Controller.UpdateNote)    // Note
		NoteRoutes.DELETE("/:id", Controller.DeleteNote) // Note
	}

	return R
}
