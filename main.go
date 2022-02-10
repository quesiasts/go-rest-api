package main

import (
	"fmt"

	"github.com/quesiasts/go-rest-api/database"
	"github.com/quesiasts/go-rest-api/models"
	"github.com/quesiasts/go-rest-api/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome1", Historia: "Historia1"},
		{Id: 2, Nome: "Nome2", Historia: "Historia2"},
	}
	database.ConectaBancoDeDados()
	fmt.Println("Iniciando o servidor Rest com Go")
	routes.HandleRequest()
}
