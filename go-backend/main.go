package main

import (
	"net/http"

	docs "backend/go/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	DatabaseConnection()

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, "Hello World!")
	})
	router.GET("/users", GetUsersAPI)
	router.GET("/user/:id")
	router.POST("/adduser", PostUserAPI)

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := router.Group("/api/v1")
	{
		eg := v1.Group("/example")
		{
			eg.GET("/users", GetUsersAPI)
		}
	}

	router.GET("/swagger/*any", func(c *gin.Context) {
		ginSwagger.WrapHandler(swaggerFiles.Handler)(c)
	})

	router.Run(":8080")
}

// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func GetUsersAPI(context *gin.Context) {

	users, _ := GetUsers()

	context.IndentedJSON(http.StatusOK, users)
}

func PostUserAPI(context *gin.Context) {
	var usr UserData

	if err := context.BindJSON(&usr); err != nil {
		return
	}

	AddUser(usr)

	context.IndentedJSON(http.StatusCreated, usr)
}
