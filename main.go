package main

import (
	"curso5/models"
	"curso5/routes"
	"fmt"
)

func main() {
	models.Alunos = []models.Aluno{
		{Id: 10,
			Nome: "Joao",
			Cpf:  "16323827760"},
		{Id: 2,
			Nome: "André",
		},
	}

	fmt.Println("Hello")
	routes.HandleRequest()
}
