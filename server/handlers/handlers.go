package handlers

/**
 * @author Dhawal Dyavanpalli <dhawalhost@gmail.com>
 * @desc Created on 2020-08-31 4:27:56 pm
 * @copyright Crearosoft
 */

import (
	"github.com/gin-gonic/gin"
)

// CommonHandler Generic Handler
func CommonHandler() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		serviceToCall := c.GetHeader("service-header")
		switch serviceToCall {
		case "GetUploadedFileData":
			GetUploadedFileData(c)
			break
		case "Login":
			Login(c)
			break
		case "Register":
			Register(c)
			break
		case "UploadFile":
			UploadFileHandlerCommon(c)
			break
		case "Backup":
			CacheBackup(c)
			break
			// case "Register":
			// 	GetUploadedFileData(c)
			// }
		}
	})
}
