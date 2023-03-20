package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/guilherm5/crudComplete/database"
	"github.com/guilherm5/crudComplete/routes"
	"github.com/joho/godotenv"
)

func main() {
	database.DBConnection()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Erro ao carregar variavel de ambiente para definir porta HTTP")
	}
	port := os.Getenv("PORT")

	Api := gin.New()
	Api.Use(gin.Logger())

	//Cadastro e login de usuario
	routes.User(Api)
	routes.Login(Api)

	//Despesas
	routes.Despesas(Api)

	//Salario
	routes.Salario(Api)

	//Total
	routes.Total(Api)

	Api.Run(port)

}
