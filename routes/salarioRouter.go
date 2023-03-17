package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherm5/crudComplete/controllers"
	"github.com/guilherm5/crudComplete/middleware"
)

func Salario(c *gin.Engine) {
	api := c.Group("api")
	api.Use(middleware.MiddlewareGO())

	api.GET("salario", controllers.GetSalario())
	api.GET("salarioID", controllers.GetSalarioID())
	api.POST("salario", controllers.PostSalario())
	api.PUT("salario", controllers.UpdateSalario())
	api.DELETE("salario", controllers.DeleteSalario())
}
