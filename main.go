package main

import (
	"crowfunding/auth"
	"crowfunding/handler"
	"crowfunding/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	
  		dsn := "root:root@tcp(127.0.0.1:3306)/campaign?charset=utf8mb4&parseTime=True&loc=Local"
  		db , err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

		if err != nil {
			log.Fatal(err.Error())
		}

		userRepository := user.NewRepository(db)
		userService := user.NewService(userRepository)
		authService := auth.NewService()

		fmt.Println(authService.GenerateToken(1001))

		userHandler := handler.NewUserHandler(userService, authService)
		router := gin.Default()

		api := router.Group("/api/v1")

		api.POST("/users", userHandler.RegisterUser)
		api.POST("/login", userHandler.Login)
		api.POST("/email_checkers", userHandler.CheckEmailAvailable)
		api.POST("/avatars", userHandler.UploadAvatar)

		router.Run()


}

