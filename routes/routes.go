package routes

import (
	"github.com/J-Eugenio/api-go-gin/controllers"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	r.GET("/alunos", controllers.ExibiTodosAlunos)
	r.GET("/:nome", controllers.Saudacao)
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)

	r.POST("/alunos", controllers.CriaNovoAluno)

	r.DELETE("/alunos/:id", controllers.DeletaAluno)

	r.PATCH("/alunos/:id", controllers.EditaAluno)

	r.Run()
}
