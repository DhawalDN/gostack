package routes

import (
	"github.com/DhawalDN/gostack/server/handlers"

	"github.com/gin-gonic/gin"
)

// InitLoginRoute :
func InitLoginRoute(o *gin.RouterGroup) {

	o.POST("/login/register", handlers.Register())
	// o.POST("/login/registerteam", handlers.RegisterTeam())
	o.POST("/login/login", handlers.Login())

}
