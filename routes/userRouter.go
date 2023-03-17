package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/guilherm5/crudComplete/controllers"
)

func User(c *gin.Engine) {
	c.GET("user", controllers.GetUsers())
	c.GET("userID", controllers.GetUsersID())
	c.POST("user", controllers.PostUsers())
	c.PUT("user", controllers.UpdateUsers())
	c.DELETE("user", controllers.DeleteUsers())
}
