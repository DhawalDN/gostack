package routes

import (
	"github.com/DhawalDN/gostack/server/handlers"

	"github.com/gin-gonic/gin"
)

// InitLoginRoute :
func InitLoginRoute(o *gin.RouterGroup) {

	o.POST("/api", handlers.CommonHandler())
	// o.POST("/login/registerteam", handlers.RegisterTeam())

}
