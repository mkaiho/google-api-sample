package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mkaiho/google-api-sample/infrastructure"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func main() {
	ctx := context.Background()
	gbpConfig, err := infrastructure.LoadGBPConfigEnv()
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	config := oauth2.Config{
		ClientID:     gbpConfig.ClientID(),
		ClientSecret: gbpConfig.ClientSecret(),
		Endpoint:     google.Endpoint,
		RedirectURL:  gbpConfig.RedirectURL(),
		Scopes: []string{
			"https://www.googleapis.com/auth/business.manage",
		},
	}

	router := gin.Default()

	router.GET("/token/code", func(c *gin.Context) {
		state := c.GetString("state")
		url := config.AuthCodeURL(state, oauth2.AccessTypeOffline, oauth2.ApprovalForce)
		c.PureJSON(http.StatusOK, gin.H{
			"url": url,
		})
	})

	router.GET("/token", func(c *gin.Context) {
		code, _ := c.GetQuery("code")
		token, err := config.Exchange(ctx, code)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": http.StatusText(http.StatusUnauthorized),
				"detail":  err.Error(),
			})
			return
		}
		c.PureJSON(http.StatusOK, gin.H{
			"accessToken":  token.AccessToken,
			"refreshToken": token.RefreshToken,
		})
	})

	router.Run(":3000")
}
