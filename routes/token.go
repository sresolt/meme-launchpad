package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/janction/meme-launchpad/controllers"
)

func RegisterTokenRoutes(r *gin.Engine) {
	r.POST("/token", controllers.CreateToken)
}
