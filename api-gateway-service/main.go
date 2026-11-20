package main

import (
	"fmt"

	"github.com/Tippenn/Social-Media-Project/wizard"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	wizard.RegisterServer(router)
	router.Run(fmt.Sprintf(":%d", wizard.Config.Server.Port))

}
