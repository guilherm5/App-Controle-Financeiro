package controllers

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/guilherm5/crudComplete/models"
	"github.com/joho/godotenv"
)

func GetSalario() gin.HandlerFunc {
	return func(c *gin.Context) {
		var getSalario []models.Salario

		rows, err := DB.Query(`SELECT * FROM salario`)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Erro ao realizar select na tabela salario ": err,
			})
			log.Println("Erro ao realizar select na tabela salario ", err)
		}

		for rows.Next() {
			var getAll models.Salario

			if err := rows.Scan(&getAll.IDSalario, &getAll.MesSalario, &getAll.Salario, &getAll.RendaExtra, &getAll.NameUser, &getAll.IDUser); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"Erro ao realizar scan para select na tabela salario ": err,
				})
				log.Println("Erro ao realizar scan para select na tabela salario ", err)
			}
			getSalario = append(getSalario, getAll)
		}
		c.JSON(http.StatusOK, getSalario)
	}
}

func GetSalarioID() gin.HandlerFunc {
	return func(c *gin.Context) {
		var getOne models.Salario

		err := c.ShouldBindJSON(&getOne)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Erro ao decodificar body para select ID na tabela Salario ": err,
			})
			log.Println("Erro ao decodificar body para select ID na tabela Salario ", err)
		}

		row := DB.QueryRow(`SELECT * FROM salario WHERE id_salario = $1`, getOne.IDSalario)

		if err := row.Scan(&getOne.IDSalario, &getOne.MesSalario, &getOne.Salario, &getOne.RendaExtra, &getOne.IDUser); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Erro ao realizar scan ta tabela salario para select ID ": err,
			})
			log.Println("Erro ao realizar scan ta tabela salario para select ID ", err)
		}

		c.JSON(http.StatusOK, getOne)
	}
}

func PostSalario() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := godotenv.Load("./.env")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Erro ao acessar arquivo .env para carregar secret key": err,
			})
			log.Println("Erro ao acessar arquivo .env para carregar secret key ", err)
			return
		}

		secret := os.Getenv("SECRET")

		tokenString := c.GetHeader("Authorization")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{
				"Message": "Token JWT inválido",
			})
			log.Println("Token JWT inválido ", err)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Erro ao obter claims do token JWT",
			})
			log.Println("Erro ao obter claims do token JWT")
			return
		}

		name, ok := claims["name"].(string)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Erro ao obter nome do usuario a partir do token JWT",
			})
			log.Println("Erro ao obter nome do usuario a partir do token JWT")
			return
		}

		sub, ok := claims["sub"].(float64)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Message": "Erro ao obter ID do usuario a partir do token JWT",
			})
			log.Println("Erro ao obter ID do usuario a partir do token JWT")
			return
		}
		userNameJwt := (name)
		userIDInt := int(sub)

		var postSalario []models.Salario

		err = c.ShouldBindJSON(&postSalario)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Erro ao decodificar body para realizar insert na tabela Salario ": err,
			})
			log.Println("Erro ao decodificar body para realizar insert na tabela Salario ", err)
			return
		}

		preparing, err := DB.Prepare(`INSERT INTO salario (mes_salario, salario, renda_extra, name_user, id_user) VALUES ($1, $2, $3, $4, $5)`)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Erro ao preparar banco de dados para realizar insert na tabela Salario ": err,
			})
			log.Println("Erro ao preparar banco de dados para realizar insert na tabela Salario ", err)
			return
		}

		for _, add := range postSalario {
			_, err = preparing.Exec(time.Now().Format("01/02/2006"), add.Salario, add.RendaExtra, userNameJwt, userIDInt)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"Erro ao executar insert na tabela Salario ": err,
				})
				log.Println("Erro ao executar insert na tabela Salario ", err)
				return
			}
			c.JSON(http.StatusOK, postSalario)
		}

	}
}

func UpdateSalario() gin.HandlerFunc {
	return func(c *gin.Context) {
		var updateSalario []models.Salario
		err := c.ShouldBindJSON(&updateSalario)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"Erro ao decodificar body para realizar update na tabela salario ": err,
			})
			log.Println("Erro ao decodificar body para realizar update na tabela salario ", err)
			return
		}

		preparing, err := DB.Prepare(`UPDATE salario SET valor =$1, renda_extra =$2`)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"Erro ao preparar banco de dados para realizar update na tabela salario ": err,
			})
			log.Println("Erro ao preparar banco de dados para realizar update na tabela salario ", err)
			return
		}

		for _, up := range updateSalario {
			_, err := preparing.Exec(up.Salario, up.RendaExtra)
			if err != nil {
				c.JSON(http.StatusOK, gin.H{
					"Erro ao executar update na tabela salario ": err,
				})
				log.Println("Erro ao executar update na tabela salario ", err)
				return
			}
		}
		c.JSON(http.StatusOK, gin.H{
			"Success": updateSalario,
		})
	}
}

func DeleteSalario() gin.HandlerFunc {
	return func(c *gin.Context) {
		var deleteSalario []models.Salario

		err := c.ShouldBindJSON(&deleteSalario)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Erro ao decodificar body para realizar delete na tabela de salario": err,
			})
			log.Println("Erro ao decodificar body para realizar delete na tabela de salario", err)
		}

		preparing, err := DB.Prepare(`DELETE FROM salario WHERE id_salario = $1`)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Erro ao preparar banco de dados para realizar delete na tabela de salario": err,
			})
			log.Println("Erro ao preparar banco de dados para realizar delete na tabela de salario", err)
		}

		for _, del := range deleteSalario {
			_, err := preparing.Exec(del.IDSalario)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"Erro ao executar delete na tabela de salario": err,
				})
				log.Println("Erro ao executar delete na tabela de salario", err)
			}
		}
		c.JSON(http.StatusOK, deleteSalario)

	}
}

/*{
	USER 1
    "id":"paulo@gmail.com",
    "secret":"123"

},
{
USER 2
    "id":"guilherme@gmail.com",
    "secret":"123"
}*/
