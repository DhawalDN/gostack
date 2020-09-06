package helpers

/**
 * @author Dhawal Dyavanpalli <dhawalhost@gmail.com>
 * @desc Created on 2020-08-31 4:27:27 pm
 * @copyright Crearosoft
 */

import (
	"errors"
	"fmt"
	"time"

	"github.com/dhawalhost/gostack/server/models"
	authmanager "github.com/crearosoft/corelib/authmanager"
	"github.com/gin-gonic/gin"
)

func init() {
	authmanager.GlobalJWTKey = "crearosoft"
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

//GetDateForPath date for only creating directory
func GetDateForPath() string {
	d := time.Now()
	year, month, date := d.Date()
	return fmt.Sprintf("%v%v%v", year, int(month), date)
}
