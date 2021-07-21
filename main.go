package main

import (
	"crowfunding/auth"
	"crowfunding/campaign"
	"crowfunding/handler"
	"crowfunding/helper"
	"crowfunding/user"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
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
		campaignRepository := campaign.NewRepository(db)

		userService := user.NewService(userRepository)
		authService := auth.NewService()
		campaignService := campaign.NewService(campaignRepository)
		
		// campaings, _ := campaignService.GetCampaigns(1)
		// fmt.Println(campaings)

		// for _, campaign := range campaigns

		// dotenv := auth.SecretKeyJwt("SECRET_KEY_JWT")

  		// fmt.Println(dotenv)

		userHandler := handler.NewUserHandler(userService, authService)
		campaignHandler := handler.NewCampaignHandler(campaignService)
		router := gin.Default()
		router.Static("/images", "./images")


		api := router.Group("/api/v1")

		// users
		api.POST("/users", userHandler.RegisterUser)
		api.POST("/login", userHandler.Login)
		api.POST("/email_checkers", userHandler.CheckEmailAvailable)
		api.POST("/avatars", authMiddleware(authService, userService) , userHandler.UploadAvatar)
		
		// campaigns
		api.GET("/campaigns", campaignHandler.GetCampaigns)
		api.GET("/campaign/:id", campaignHandler.GetCampaign)
		api.POST("/campaigns", authMiddleware(authService, userService) ,campaignHandler.CreateCampaign)
		api.PUT("/campaign/:id", authMiddleware(authService, userService), campaignHandler.UpdateCampaign)
		api.POST("/campaign-images", authMiddleware(authService, userService), campaignHandler.UploadImage)
		router.Run()


}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func (c *gin.Context)  {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claims["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}



