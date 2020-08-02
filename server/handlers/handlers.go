package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// CommonHandler Generic Handler
func CommonHandler() gin.HandlerFunc {
	return gin.HandlerFunc(func(c *gin.Context) {
		serviceToCall := c.GetHeader("service-header")
		fmt.Println(serviceToCall)
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
		case "UploadFileHandler":
			UploadFileHandler(c)
			break
			// case "Register":
			// 	GetUploadedFileData(c)
			// }
		}
	})
}
