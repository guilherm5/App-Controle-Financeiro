package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherm5/crudComplete/controllers"
	"github.com/guilherm5/crudComplete/middleware"
)

func Total(c *gin.Engine) {
	api := c.Group("api")
	api.Use(middleware.MiddlewareGO())

	api.GET("total", controllers.GetTotalDespesas())
}
