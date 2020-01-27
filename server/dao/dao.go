package dao

var (

	// UsersDAO - For DB Operation
	UsersDAO *MongoDAO
	// LoginDAO - For DB Operation
	LoginDAO *MongoDAO
	// TeamsDAO *dao.MongoDAO
)

//InitDAO - Associate DAOs with respective collections
//Create new DAO here for new collection
func InitDAO() {
	// PlayersDAO = dao.GetMongoDAO("players")
	// TeamsDAO = dao.GetMongoDAO("teams")
	// BiddingDAO = dao.GetMongoDAO("bidding")
	UsersDAO = GetMongoDAO("users")
	LoginDAO = GetMongoDAO("login")
}
