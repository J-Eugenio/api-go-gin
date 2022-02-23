package main

import (
	"github.com/J-Eugenio/api-go-gin/models"
	"github.com/J-Eugenio/api-go-gin/routes"
)

func main() {
	models.Alunos = []models.Aluno{
		{
			Nome: "Teste1",
			CPF:  "00000000000",
			RG:   "000000-00",
		},
		{
			Nome: "Teste2",
			CPF:  "11111111111",
			RG:   "111111-11",
		},
	}
	routes.HandleRequests()
}
