package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"

	"github.com/guilherm5/crudComplete/models"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func LoginUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var login models.Login
		var usuario = models.User{}

		err := c.ShouldBindJSON(&login)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Erro ao decodificar body para preencher informações de login ": err,
			})
			log.Println("Erro ao decodificar body para preencher informações de login  ", err)
			return
		}

		verificator := fmt.Sprintf(`SELECT secret_user, email_user, id_user, name_user FROM users WHERE email_user = '%s'`, login.ID)
		emailRows, err := DB.Query(verificator)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Email inexistente ": err,
			})
			log.Println("Email inexistente ", err)
			return
		}

		emailFound := false
		for emailRows.Next() {
			emailFound = true
			err := emailRows.Scan(&usuario.SecretUser, &usuario.EmailUser, &usuario.IDUser, &usuario.NameUser)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"Erro ao realizar scan na tabela users para verificação de email ": err,
				})
				log.Println("Erro ao realizar scan na tabela users para verificação de email", err)
				return
			}
		}

		Inexistente := fmt.Sprintf("Usuário com email [%s] não encontrado", login.ID)

		// Se nenhum registro foi encontrado, retornar erro de email ou conta inexistente
		if !emailFound {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Erro": Inexistente,
			})
			log.Println(Inexistente, err)
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(usuario.SecretUser), []byte(login.Secret))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Senha incorreta": err,
			})
			log.Println("Senha incorreta", err)
			return
		}

		Token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email": usuario.EmailUser,
			"sub":   usuario.IDUser,
			"name":  usuario.NameUser,
			"exp":   time.Now().Add(time.Hour).Unix(),
		})

		err = godotenv.Load("./.env")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Erro ao carregar arquivo .env para gerar token JWT ",
			})
			log.Println("Erro ao carregar arquivo .env para gerar token JWT ", err)
			return
		}

		secret := os.Getenv("SECRET")

		TokenString, err := Token.SignedString([]byte(secret))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Erro ao gerar token JWT": err,
			})
			log.Println("Erro ao gerar token JWT ", err.Error())
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"Token JWT": TokenString,
		})

	}
}
