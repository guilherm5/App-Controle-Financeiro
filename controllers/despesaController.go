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

func GetDespesas() gin.HandlerFunc {
	return func(c *gin.Context) {
		var getDespesas []models.Despesas

		rows, err := DB.Query(`SELECT * FROM despesas`)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Erro ao realizar select na tabela despesas ": err,
			})
			log.Println("Erro ao realizar select na tabela despesas ", err)
			return
		}

		for rows.Next() {
			var getAll models.Despesas

			if err := rows.Scan(&getAll.IDDespesa, &getAll.NomeDespesa, &getAll.DescricaoDespesa, &getAll.MesDespesa, &getAll.Valor, &getAll.NameUser, &getAll.IDUser); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"Erro ao realizar scan para select na tabela despesas": err,
				})
				log.Println("Erro ao realizar scan para select na tabela despesas", err)
				return
			}
			getDespesas = append(getDespesas, getAll)
		}
		c.JSON(http.StatusOK, getDespesas)
	}
}

func GetDespesasID() gin.HandlerFunc {
	return func(c *gin.Context) {
		var GetOne models.Despesas

		err := c.ShouldBindJSON(&GetOne)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"Erro ao decodificar body para realizar select ID na tabela despesa": err,
			})
			log.Println("Erro ao decodificar body para realizar select ID na tabela despesa", err)
			return
		}

		row := DB.QueryRow(`SELECT * FROM despesas WHERE id_despesa = $1`, GetOne.IDDespesa)

		if err := row.Scan(&GetOne.IDDespesa, &GetOne.NomeDespesa, &GetOne.DescricaoDespesa, &GetOne.MesDespesa, &GetOne.Valor, &GetOne.NameUser, &GetOne.IDUser); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Erro ao realizar scan na tabela despesas para select ID": err,
			})
			log.Println("Erro ao realizar scan na tabela despesas para select ID", err)
			return
		}
		c.JSON(http.StatusOK, GetOne)
	}
}

func PostDespesas() gin.HandlerFunc {
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

		var postDespesa []models.Despesas
		err = c.ShouldBindJSON(&postDespesa)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Erro ao decodificar body para realizar post na tabela despesa": err,
			})
			log.Println("Erro ao decodificar body para realizar post na tabela despesa", err)
			return
		}

		preparing, err := DB.Prepare(`INSERT INTO despesas (nome_despesa, descricao_despesa, mes_despesa, valor, name_user, id_user) VALUES ($1, $2, $3, $4, $5, $6)`)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Erro ao preparar banco de dados para realizar insert na tabela despesas": err,
			})
			log.Println("Erro ao preparar banco de dados para realizar insert na tabela despesas", err)
			return
		}

		for _, add := range postDespesa {

			_, err := preparing.Exec(add.NomeDespesa, add.DescricaoDespesa, time.Now().Format("01/02/2006"), add.Valor, userNameJwt, userIDInt)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"Erro ao executar insert na tabela despesas": err.Error(),
				})
				log.Println("Erro ao executar insert na tabela despesas", err.Error())
				return
			}
		}
	}
}

func UpdateDespesas() gin.HandlerFunc {
	return func(c *gin.Context) {
		var putDespesas []models.Despesas

		err := c.ShouldBindJSON(&putDespesas)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Erro ao decodificar body para realizar update na despesa selecionada ": err,
			})
			log.Println("Erro ao decodificar body para realizar update na despesa selecionada ", err)
			return
		}

		preparing, err := DB.Prepare(`UPDATE despesas SET nome_despesa =$1, descricao_despesa =$2, valor =$3 WHERE id_despesa = $4`)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Erro ao preparar banco de dados para realizar update na despesa selecionada ": err,
			})
			log.Println("Erro ao preparar banco de dados para realizar update na despesa selecionada ", err)
			return
		}

		for _, up := range putDespesas {
			_, err := preparing.Exec(&up.NomeDespesa, &up.DescricaoDespesa, &up.Valor, &up.IDDespesa)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"Erro ao executar update na despesa selecionada ": err,
				})
				log.Println("Erro ao executar update na despesa selecionada ", err)
				return
			}
			c.JSON(http.StatusOK, up)
		}

	}
}

func DeleteDespesa() gin.HandlerFunc {
	return func(c *gin.Context) {
		var deleteDespesa []models.Despesas

		err := c.ShouldBindJSON(&deleteDespesa)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Erro ao decodificar body para realizar operação de delete na tabela despesas ": err,
			})
			log.Println("Erro ao decodificar body para realizar operação de delete na tabela despesas ", err)
			return
		}

		preparing, err := DB.Prepare(`DELETE FROM despesas WHERE id_despesa = $1`)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"Erro ao preparar banco de dados para realizar operação de delete na tabela despesas ": err,
			})
			log.Println("Erro ao preparar banco de dados para realizar operação de delete na tabela despesas ", err)
			return
		}

		for _, del := range deleteDespesa {
			_, err := preparing.Exec(&del.IDDespesa)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"Erro ao executar delete na tabela despesas ": err,
				})
				log.Println("Erro ao executar delete na tabela despesas ", err)
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"Success": err,
			})
		}
	}
}
