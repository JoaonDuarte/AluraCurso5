package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.JSON(200, gin.H{
		"id":   "1",
		"nome": "João Duarte",
	})
}

func main() {
	fmt.Println("Hello")
	r := gin.Default()
	r.GET("/home", Home)
	r.Run(":8001")
	// log.Fatal(http.ListenAndServe(":8001", r))
}
