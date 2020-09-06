package routes

/**
 * @author Dhawal Dyavanpalli <dhawalhost@gmail.com>
 * @desc Created on 2020-08-31 4:28:45 pm
 * @copyright Crearosoft
 */

import (
	"github.com/DhawalDN/gostack/server/handlers"

	"github.com/gin-gonic/gin"
)

// InitLoginRoute :
func InitLoginRoute(o *gin.RouterGroup) {

	o.POST("/api", handlers.CommonHandler())
	// o.POST("/login/registerteam", handlers.RegisterTeam())

}
