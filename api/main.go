package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/users/:name", getUsersHandler)
	r.POST("/users", addUserHandler)
	r.Run()

}

func addUserHandler(context *gin.Context) {
	context.JSON(201, gin.H{
		"name": "john",
		"lastname": "mat",
	})
}

func getUsersHandler(context *gin.Context) {
	name := context.Param("name")
	lastname := context.Query("lastname")
	context.JSON(200, gin.H{
		"name": name,
		"lastname": lastname,
	})
}