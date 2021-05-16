package main

import (
	"crowfunding/user"
	"log"

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

		userInput := user.RegisterUserInput{}
		userInput.Name = "Tes simpan dari service"
		userInput.Email = "email@gnas.co"
		userInput.Password = "123456"
		userInput.Occupation = "Engineer"

		userService.RegisterUser(userInput)

}

