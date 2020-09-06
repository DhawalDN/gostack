package routes

/**
 * @author Dhawal Dyavanpalli <dhawalhost@gmail.com>
 * @desc Created on 2020-08-31 4:28:38 pm
 * @copyright Crearosoft
 */

import (
	"github.com/DhawalDN/gostack/server/handlers"
	"github.com/gin-gonic/gin"
)

// InitCdnRoutes :
func InitCdnRoutes(o, r *gin.RouterGroup) {
	o.POST("/api", handlers.CommonHandler())
	r.POST("/api", handlers.CommonHandler())

	// o.StaticFS("/files/", gin.Dir(handlers.UploadPath, true))
	// o.Static("/files/", handlers.UploadPath)
	// fs := http.FileServer(http.Dir(handlers.UploadPath))
	// http.Handle("/files/", http.StripPrefix("/files", fs))
	// o.POST("/login/registerteam", handlers.RegisterTeam())

}
