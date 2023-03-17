package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherm5/crudComplete/controllers"
	"github.com/guilherm5/crudComplete/middleware"
)

func Despesas(c *gin.Engine) {
	api := c.Group("api")
	api.Use(middleware.MiddlewareGO())

	api.GET("despesa", controllers.GetDespesas())
	api.GET("despesaID", controllers.GetDespesasID())
	api.POST("despesa", controllers.PostDespesas())
	api.PUT("despesa", controllers.UpdateDespesas())
	api.DELETE("despesa", controllers.DeleteDespesa())
}
