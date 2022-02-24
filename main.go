package main

import (
	"github.com/J-Eugenio/api-go-gin/database"
	"github.com/J-Eugenio/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()

	routes.HandleRequests()
}
