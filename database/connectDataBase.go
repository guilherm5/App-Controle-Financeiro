package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func DBConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Erro ao carregar variavel de ambiente para capturar dados de conexão com banco de dados")
	}

	host := os.Getenv("HOST")
	portdb := os.Getenv("PORTDB")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")

	stringConnection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, portdb, user, password, dbname)

	db, err := sql.Open("postgres", stringConnection)
	if err != nil {
		log.Println("Erro ao realizar conexão com banco de dados ", err)
	} else {
		fmt.Println("Sucesso ao realizar conexão com banco de dados")
	}
	return db
}
