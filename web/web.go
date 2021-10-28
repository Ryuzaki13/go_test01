package web

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Run() {
	router := gin.Default()

	router.LoadHTMLGlob("html/*.html")
	router.Static("/assets/", "resources")

	router.GET("/", handlerIndex)
	router.PUT("/user", handlerCreateUser)
	router.GET("/user", handlerGetUser)

	router.Run("127.0.0.1:8080")
}

func handlerIndex(c *gin.Context) {
	c.HTML(200, "index", nil)
}

func handlerGetUser(c *gin.Context) {
	users := [3]string{"Вася", "Петя", "Григорий"}
	c.JSON(200, users)
}

func handlerCreateUser(c *gin.Context) {
	type User struct {
		Login string `json:"name"`
	}
	var user User

	e := c.BindJSON(&user)
	if e != nil {
		fmt.Println(e)
	}

	c.JSON(200, nil)
}
