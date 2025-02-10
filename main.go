package main

import (
	"curso5/database"
	"curso5/routes"
	"fmt"
)

func main() {
	

	fmt.Println("Hello")
	database.ConectaDB()
	routes.HandleRequest()
}
