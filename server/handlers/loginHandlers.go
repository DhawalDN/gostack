package handlers

/**
 * @author Dhawal Dyavanpalli <dhawalhost@gmail.com>
 * @desc Created on 2020-08-31 4:27:48 pm
 * @copyright Crearosoft
 */

import (
	"fmt"
	"net/http"

	"os"

	"github.com/dhawalhost/gostack/server/models"
	"github.com/dhawalhost/gostack/server/services"
	"github.com/gin-gonic/gin"
)

//Login : Login
func Login(c *gin.Context) {
	loginDetails := models.Login{}
	if c.Bind(&loginDetails) != nil {
		fmt.Println("Coulnt Bind")
	}
	data, token := services.LoginUser(loginDetails)
	fmt.Println(token)
	c.Header("Authorization", token)
	c.JSON(http.StatusOK, data)
	fmt.Println(os.Getwd())
	// c.String(http.StatusOK, )

}

// Register func
func Register(c *gin.Context) {
	buyer := models.User{}
	err := c.Bind(&buyer)
	if err != nil {
		fmt.Print(err)
		c.JSON(http.StatusExpectationFailed, err)
	}
	services.AddUser(buyer)
	c.String(http.StatusOK, "Customer registered successfully! \n ")

}

// func RegisterTeam() gin.HandlerFunc {
// 	return func(c *gin.Context) {
// 		team := models.Team{}
// 		err := c.Bind(&team)
// 		if err != nil {
// 			fmt.Print(err)
// 			c.JSON(http.StatusExpectationFailed, err)
// 		}
// 		services.AddTeams(team)
// 		c.String(http.StatusOK, "Team registered successfully! \n ")
// 	}
// }
