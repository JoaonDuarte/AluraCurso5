package routes

import (
	"curso5/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/home", controllers.Home)
	r.GET("/:nome", controllers.BuscaAluno)
	r.POST("/criar", controllers.CriarAluno)
	r.Run(":8001")
	// log.Fatal(http.ListenAndServe(":8001", r))
}
