package main

import (
	"fmt"
	"log"

	users "github.com/Tippenn/Social-Media-Project/domains/users/entities"
	"github.com/Tippenn/Social-Media-Project/wizard"
	"github.com/gin-gonic/gin"
)

func main() {
	err := wizard.PostgresDatabase.GetInstance().AutoMigrate(
		&users.User{},
		&users.Session{},
	)

	if err != nil {
		log.Fatal("migration failed: ", err)
	}
	
	router := gin.Default()

	wizard.RegisterServer(router)
	router.Run(fmt.Sprintf(":%d", wizard.Config.Server.Port))
}
