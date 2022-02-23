package routes

import (
	"github.com/J-Eugenio/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/alunos", controllers.ExibiTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)

	r.Run()
}
