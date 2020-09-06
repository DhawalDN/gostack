package models

/**
 * @author Dhawal Dyavanpalli <dhawalhost@gmail.com>
 * @desc Created on 2020-08-31 4:26:41 pm
 * @copyright Crearosoft
 */
import (
	"github.com/crearosoft/corelib/cachemanager"
)

// ProjectID that should be used for routing before Request
var ProjectID string

//User : User
type User struct {
	Name         string `json:"name"`
	MobileNumber string `json:"mobileno"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Username     string `json:"username"`
}

//Login : Login user
type Login struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

// FC is fastCache instance that can be initialized during start of server
var FC *cachemanager.CacheHelper
