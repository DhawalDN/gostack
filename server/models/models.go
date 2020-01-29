package models

// ProjectID that should be used for routing before Request
var ProjectID string

//Player : struct
type Player struct {
	PlayerID  string
	BasePrice int64
}

//Bidder : struct
type Bidder struct {
	Bid       int    `json:"bid" bson:"bid"`
	BuyerName string `json:"buyerName" bson:"buyerName"`
	TeamName  string `json:"teamName" bson:"teamName"`
}

//User : User
type User struct {
	Name         string `json:"name"`
	MobileNumber string `json:"mobileno"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Username     string `json:"username"`
}

//BuyHelp : To add player
type BuyHelp struct {
	BuyerName string `json:"buyerName"`
	PlayerID  string `json:"playerId"`
}

//Login : Login user
type Login struct {
	UserName string `json:"username"`
	Password string `json:"password"`
}

//UpdateBid : Update current bid
type UpdateBid struct {
	Bid      int    `json:"bid" bson:"bid"`
	PlayerID string `json:"playerId" bson:"playerId"`
}

// type BuyerName struct {
// 	Name string `json:"name"`
// }
