package routes

import (
	"database/sql"

	"go_mysql2/controllers"
	"go_mysql2/repositories"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	r := gin.Default()

	userRepo := repositories.NewUserRepository(db)
	userController := controllers.NewUserController(*userRepo)

	v1 := r.Group("/api/v1")
	{
		v1.GET("/users", userController.GetAllUsers)
		v1.POST("/users", userController.CreateUser)
	}

	return r
}
