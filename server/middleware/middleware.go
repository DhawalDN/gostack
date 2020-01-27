package middleware

import (
	"fmt"
	"strings"

	"github.com/DhawalDN/gostack/server/helpers"
	"github.com/DhawalDN/gostack/server/routes"
	"github.com/DhawalDN/gostack/server/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// InitMiddleware -Init
func InitMiddleware(g *gin.Engine) {
	g.Use(cors.Default()) // CORS request
	fmt.Println("InitMiddleware called")
	o := g.Group("/o")
	o.Use(OpenRequestMiddleware())
	r := g.Group("/r")
	r.Use(RestrictedRequestMiddleware())
	c := r.Group("/c")
	c.Use(RoleBasedRequestMiddleware())
	// routes.InitAuctionRoute(r)
	routes.InitLoginRoute(o)
	routes.InitCdnRoutes(o)

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
		userres, isValid, usererr := services.ValidatedUser(login)
		if usererr != nil || !isValid {
			fmt.Println("Failed to validate user", userres)
			c.AbortWithStatusJSON(401, gin.H{"error": "Failed to validate user"})
		}
		if !isValid {
			c.AbortWithStatusJSON(401, gin.H{"error": "User Not Exists"})
		}
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
