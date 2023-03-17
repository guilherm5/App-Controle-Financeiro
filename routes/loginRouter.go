package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherm5/crudComplete/controllers"
)

func Login(c *gin.Engine) {
	c.POST("login", controllers.LoginUser())
}
