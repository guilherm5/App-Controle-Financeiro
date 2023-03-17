package controllers

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/guilherm5/crudComplete/models"
	"github.com/joho/godotenv"
)

func GetTotalDespesas() gin.HandlerFunc {
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

		sub, ok := claims["sub"].(float64)
		if !ok {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Erro ao obter ID do usuario a partir do token JWT": err,
			})
			log.Println("Erro ao obter ID do usuario a partir do token JWT", err)
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

		//Definindo parametros para despesas e saldo, para chamar despesa e salario DAQUELE usuario

		//Saldo
		var responseSaldo models.TotalSalarioResponse
		responseSaldo.NameUserSaldo = name
		responseSaldo.IDUserSaldo = int(sub)

		//Despesa
		var responseDespesas models.TotalDespesasResponse
		responseDespesas.NameUser = name
		responseDespesas.IDUser = int(sub)

		rows, err := DB.Query(`SELECT SUM(valor) as total_despesas FROM public.despesas WHERE id_user = $1 AND name_user = $2`, responseDespesas.IDUser, responseDespesas.NameUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Erro ao realizar select do total de despesas": err,
			})
			log.Println("Erro ao realizar select do total de despesas", err)
			return
		}

		var totalDespesas *float64
		for rows.Next() {
			if err := rows.Scan(&totalDespesas); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"Erro ao realizar scan para select do total de despesas": err,
				})
				log.Println("Erro ao realizar scan para select do total de despesas", err)
				return
			}
		}

		if totalDespesas == nil {
			responseDespesas.TotalDespesas = 0
		} else {
			responseDespesas.TotalDespesas = *totalDespesas
		}

		saldo, err := DB.Query(`SELECT SUM(salario + renda_extra) as total_salario FROM public.salario WHERE id_user = $1 AND name_user = $2`, responseSaldo.IDUserSaldo, responseSaldo.NameUserSaldo)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"Erro ao realizar select do total salario": err,
			})
			log.Println("Erro ao realizar select do total salario", err)
			return
		}

		var totalSalario *float64
		for saldo.Next() {
			if err := saldo.Scan(&totalSalario); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"Erro ao realizar scan para select do total de salario": err,
				})
				log.Println("Erro ao realizar scan para select do total de salario", err)
				return
			}
			if totalSalario == nil {
				responseSaldo.TotalSaldo = 0
				responseSaldo.TotalSaldo = *totalSalario
			} else {
				responseSaldo.TotalSaldo = *totalSalario
			}
		}

		percentualDespesas := (float64(responseDespesas.TotalDespesas) / float64(responseSaldo.TotalSaldo)) * 100

		var resultadoPercentual = fmt.Sprintf("O percentual de despesas em relação ao salário é de %.2f%%.", percentualDespesas)
		totalMes := responseSaldo.TotalSaldo - responseDespesas.TotalDespesas

		c.JSON(http.StatusOK, gin.H{
			"Despesas": responseDespesas,
			"Saldo":    responseSaldo,
		})

		c.JSON(http.StatusOK, gin.H{
			"Total salario - despesas": totalMes,
		})
		c.JSON(http.StatusOK, gin.H{
			"Percentual:": resultadoPercentual,
		})

	}
}
