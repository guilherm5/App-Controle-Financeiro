package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/guilherm5/crudComplete/database"
	"github.com/guilherm5/crudComplete/models"
	"golang.org/x/crypto/bcrypt"
)

var DB = database.DBConnection()

func GetUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var getUser []models.User

		rows, err := DB.Query(`SELECT * FROM users`)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Erro ao realizar select na tabela users ": err,
			})
			log.Println("Erro ao realizar select na tabela users ", err)
			return
		}

		for rows.Next() {
			var getAll models.User
			if err := rows.Scan(&getAll.IDUser, &getAll.NameUser, &getAll.EmailUser, &getAll.SecretUser); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"Erro ao scanear tabela user para realizar select ": err,
				})
				log.Println("Erro ao scanear tabela user para realizar select ", err)
				return
			}
			getUser = append(getUser, getAll)
		}
		c.JSON(http.StatusOK, getUser)

	}
}

func GetUsersID() gin.HandlerFunc {
	return func(c *gin.Context) {
		var getOne models.User

		err := c.ShouldBindJSON(&getOne)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Erro ao decodificar body para realizar select ID na tabela users ": err,
			})
			log.Println("Erro ao decodificar body para realizar select ID na tabela users ", err)
			return
		}

		row := DB.QueryRow(`SELECT * FROM users WHERE id_user = $1`, getOne.IDUser)

		if err := row.Scan(&getOne.IDUser, &getOne.NameUser, &getOne.EmailUser, &getOne.SecretUser); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Erro ao scaear tabela para realizar select ID na tabela users ": err,
			})
			log.Println("Erro ao scaear tabela para realizar select ID na tabela users ", err)
			return
		}
		c.JSON(http.StatusOK, getOne)
	}
}

func PostUsers() gin.HandlerFunc {
	return func(c *gin.Context) {
		var PostUsers []models.User

		err := c.ShouldBindJSON(&PostUsers)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Erro ao decodificar body para realizar insert na tabela users": err,
			})
			log.Println("Erro ao decodificar body para realizar insert na tabela users", err)
			return
		}

		preparing, err := DB.Prepare(`INSERT INTO users (name_user, email_user, secret_user) VALUES ($1, $2, $3)`)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Erro ao preparar banco de dados para realizar insert ": err,
			})
			log.Println("Erro ao preparar banco de dados para realizar insert ", err)
			return
		}

		for _, add := range PostUsers {
			senha, err := bcrypt.GenerateFromPassword([]byte(add.SecretUser), 14)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"Erro ao encripitar senha ": err,
				})
				log.Println("Erro ao encripitar senha ", err)
				return
			}

			_, err = preparing.Exec(add.NameUser, add.EmailUser, senha)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"Erro ao realizar insert ": err,
				})
				log.Println("Erro ao realizar insert ", err)
				return
			}
		}
		c.JSON(http.StatusOK, PostUsers)
	}
}

func UpdateUsers() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func DeleteUsers() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
