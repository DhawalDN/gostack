package middleware

/**
 * @author Dhawal Dyavanpalli <dhawalhost@gmail.com>
 * @desc Created on 2020-08-31 4:27:15 pm
 * @copyright Crearosoft
 */

import (
	"fmt"
	"strings"

	"github.com/DhawalDN/gostack/server/helpers"
	"github.com/DhawalDN/gostack/server/models"
	"github.com/DhawalDN/gostack/server/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// InitMiddleware -Init
func InitMiddleware(g *gin.Engine) {
	g.Use(cors.Default()) // CORS request
	fmt.Println("InitMiddleware called")
	v1 := g.Group("/" + models.ProjectID + "")
	{
		o := v1.Group("/o")
		o.Use(OpenRequestMiddleware())
		r := v1.Group("/r")
		r.Use(RestrictedRequestMiddleware())
		c := v1.Group("/c")
		c.Use(RoleBasedRequestMiddleware())

		// routes.InitLoginRoute(o)
		routes.InitCdnRoutes(o, r)
	}

	// routes.InitAuctionRoute(r)

}

// OpenRequestMiddleware :
func OpenRequestMiddleware() gin.HandlerFunc {

	return func(c *gin.Context) {
		c.Next()
		fmt.Println("OpenRequestMiddleware called")
	}
}

// RestrictedRequestMiddleware :
// Need to check JWT token here
func RestrictedRequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		login, err := helpers.GetUserNameFromToken(c)
		if err != nil {
			fmt.Println("Token not available", err)
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid API token"})
		}
		if strings.Trim(token, "") == "" {
			fmt.Println("Token not available")
			c.AbortWithStatusJSON(401, gin.H{"error": "Invalid API token"})
		}
		if login.UserName == "" {
			c.AbortWithStatusJSON(401, gin.H{"error": "Token is invalid"})
		}
		// userres, isValid, usererr := services.ValidatedUser(login)
		// if usererr != nil || !isValid {
		// 	fmt.Println("Failed to validate user", userres)
		// 	c.AbortWithStatusJSON(401, gin.H{"error": "Failed to validate user"})
		// }
		// if !isValid {
		// 	c.AbortWithStatusJSON(401, gin.H{"error": "User Not Exists"})
		// }
		c.Next()
		fmt.Println("RestrictedRequestMiddleware called")
	}
}

// RoleBasedRequestMiddleware : Need to check JWT token here with group validation
func RoleBasedRequestMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("RoleBasedRequestMiddleware called")

		//
	}
}
