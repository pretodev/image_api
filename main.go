package main

import (
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"main/docs"
	"main/server"
)

// @title           Photo API
// @version         1.0
// @description     Esta api salvar imagem de usuário
// @termsOfService  http://swagger.io/terms/

// @contact.name   Silas Ribeiro
// @contact.email  santorsilas@gmail.com

// @BasePath  /api/v1

func main() {
	controllers := server.Controllers{}

	router := gin.Default()

	router.Static("/static", "./static")

	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := router.Group("/api/v1")
	{
		user := v1.Group("/users")
		{
			user.PATCH("/:id/photo", controllers.UpsertUserPhoto)
			user.GET("/:id/photo", controllers.GetUserPhoto)

		}
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
