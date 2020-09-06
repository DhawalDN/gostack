package services

/**
 * @author Dhawal Dyavanpalli <dhawalhost@gmail.com>
 * @desc Created on 2020-08-31 4:26:57 pm
 * @copyright Crearosoft
 */

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/crearosoft/corelib/authmanager"

	"github.com/tidwall/gjson"

	"github.com/tidwall/sjson"

	"github.com/dhawalhost/gostack/server/dao"
	"github.com/dhawalhost/gostack/server/models"
	"go.mongodb.org/mongo-driver/bson"
)

// AddUser :
func AddUser(user models.User) {
	passwd := getPwd(user.Password)
	hashedPasswd := hashAndSalt(passwd)
	userCredentialsTemp := `{}`
	username := ""
	exists := false
	if user.Username != "" {
		username = user.Username
		existinguser := GetUser(username)
		if existinguser != "" {
			log.Fatal("User Already Exists")
		}
	} else {
		username, exists = CreateUsername(user.Name)
		if exists {
			log.Fatal("user already exists")
		}
		user.Username = username
	}

	userCredentialsTemp, _ = sjson.Set(userCredentialsTemp, "username", username)
	userCredentialsTemp, _ = sjson.Set(userCredentialsTemp, "password", hashedPasswd)
	userCredentials := gjson.Parse(userCredentialsTemp)
	res := dao.UsersDAO.Insert(user)
	// loginres := dao.LoginDAO.Insert(userCredentials.Value())
	if res != nil {
		fmt.Println("Error in registration")
	} else {

		credentials := CreateCredentials(userCredentials.Value())
		if credentials != nil {
			log.Fatal("Problem while creating Credentials")
		}
		fmt.Println("Successfully Registered")
	}
}

// CreateUsername : create new username for logging in to the system
func CreateUsername(name string) (string, bool) {
	tempname := strings.Split(name, " ")
	username := ""
	if len(tempname) == 2 {
		username = strings.ToLower(tempname[0]) + strings.ToLower(string(tempname[1][0]))
	} else if len(tempname) == 3 {
		username = strings.ToLower(tempname[0]) + strings.ToLower(string(tempname[2][0]))
	} else {
		username = strings.ToLower(tempname[0])
	}
	existinguser := GetUser(username)
	fmt.Println("existing user", existinguser)
	if existinguser != "" {
		return username, true
	}
	return username, false
}

// CreateCredentials : create credentials for registered user.
func CreateCredentials(user interface{}) error {
	loginres := dao.LoginDAO.Insert(user)
	if loginres != nil {
		fmt.Println("Error in registration")
	} else {
		fmt.Println("Successfully Registered")
	}
	return loginres
}

// func AddTeams(team interface{}) {
// 	res := datamodel.TeamsDAO.Insert(team)
// 	if res != nil {
// 		fmt.Println("Error in registering team")
// 	} else {
// 		fmt.Println("Team is registered successfully")
// 	}

// }

//LoginUser : Login Buyer to the Auction
func LoginUser(login models.Login) (interface{}, string) {
	getUserDBResult, err := dao.LoginDAO.FindData(bson.M{"username": login.UserName})
	if err != nil {
		log.Fatal("LoginUser : UNABLE_TO_FETCH_BUYER")
		return nil, ""
	}
	user := getUserDBResult.Get("0")
	// hashedPasswd := hashAndSalt(getPwd(login.Password))
	// login.Password = hashedPasswd
	matched := comparePasswords(user.Get("password").String(), getPwd(login.Password))
	if !matched {
		fmt.Println("Password does not match")
		return bson.M{"token": false}, ""
	}
	// if hashedPasswd != user.Get("password").String() {
	// 	fmt.Println("Password does not match")
	// 	return bson.M{"token": false}, ""
	// }
	expiresAt := time.Now().Add(24 * time.Hour).Unix()
	token, tokenError := authmanager.GenerateToken(login.UserName, expiresAt)
	if tokenError != nil {
		log.Fatal("Unable to generate token")
		return nil, ""
	}
	return bson.M{"token": true, "name": user.Get("username").String()}, token
}

//ValidatedUser : Check for valid user
func ValidatedUser(userRequest models.Login) (string, bool, error) {
	user := GetValidatedUser(userRequest)
	if user != "" {
		return user, true, nil
	}
	return user, false, nil
}

//GetUser : Return requested user
func GetUser(name string) string {

	dbResult, err := dao.LoginDAO.FindData(bson.M{"username": name})
	if err != nil {
		log.Fatal("Failed to get user")
	}
	fmt.Println(dbResult)
	userData := dbResult.Get("0").Get("username").String()
	return userData
}

// GetValidatedUser :
func GetValidatedUser(user models.Login) string {

	dbResult, err := dao.UsersDAO.FindData(bson.M{"username": user.UserName, "password": user.Password})
	if err != nil {
		log.Fatal("Failed to get user")
	}
	userData := dbResult.Get("0").Get("name").String()
	return userData
}
