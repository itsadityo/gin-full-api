package main

import (
	"github.com/gin-gonic/gin"
	"github.com/itsadityo/gin-full-api/config"
	"github.com/itsadityo/gin-full-api/routes"
	"github.com/subosito/gotenv"
)

func main() {
	//set up database
	config.InitDB()
	defer config.DB.Close()
	gotenv.Load()

	//set up routing/router
	router := gin.Default()
	//endopint
	v1 := router.Group("api/v1/")
	{
		v1.GET("/auth/:provider", routes.RedirectHandler)
		v1.GET("/auth/:provider/callback", routes.CallbackHandler)

		articles := v1.Group("/article")

		{
			articles.GET("/", routes.GetHome)
			articles.GET("/:slug", routes.GetArticle)
			articles.POST("/", routes.PostArticle)
		}
	}

	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
