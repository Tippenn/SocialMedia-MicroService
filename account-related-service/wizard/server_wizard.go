package wizard

import (
	"github.com/gin-gonic/gin"
)

func RegisterServer(router *gin.Engine) {

	api := router.Group("/api/v1")
	{
		//User
		api.POST("/users/login", UserHttp.Login)
		api.POST("/users/register", UserHttp.Register)
	}

}
