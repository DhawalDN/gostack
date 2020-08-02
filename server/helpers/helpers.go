package helpers

import (
	"errors"

	authmanager "github.com/Crearosoft/corelib/authmanager"
	"github.com/DhawalDN/gostack/server/models"
	"github.com/gin-gonic/gin"
)

func init() {
	authmanager.GlobalJWTKey = "Crearosoft"
}

// GetUserNameFromToken login object from JWT Token
func GetUserNameFromToken(c *gin.Context) (models.Login, error) {
	login := models.Login{}
	authToken := c.GetHeader("Authorization")
	decodedToken, err := authmanager.DecodeJWTToken(authToken)
	// decodedToken, err :=  DecodeToken(c.GetHeader("Authorization"), authmanager.GlobalJWTKey)

	if err != nil {
		return login, errors.New("GetLoginFromToken - unable to decode token")
	}
	// login ID is the compulsary field, so haven't added check for nil
	if decodedToken["username"] == nil || decodedToken["username"] == "" {
		return login, errors.New("GetLoginFromToken - login id not found")
	}
	login.UserName = decodedToken["username"].(string)
	login.Password = decodedToken["username"].(string)
	return login, nil
}
