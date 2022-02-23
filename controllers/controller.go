package controllers

import (
	"github.com/J-Eugenio/api-go-gin/models"
	"github.com/gin-gonic/gin"
)

func ExibiTodosAlunos(c *gin.Context) {
	c.JSON(200, models.Alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz:": "E ai " + nome + ", tudo beleza",
	})
}
