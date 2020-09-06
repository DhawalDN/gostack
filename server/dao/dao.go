package dao

/**
 * @author Dhawal Dyavanpalli <dhawalhost@gmail.com>
 * @desc Created on 2020-08-31 4:28:11 pm
 * @copyright Crearosoft
 */

var (

	// UsersDAO - For DB Operation
	UsersDAO *MongoDAO
	// LoginDAO - For DB Operation
	LoginDAO *MongoDAO
	// CdnDAO *dao.MongoDAO
	CdnDAO *MongoDAO
)

//InitDAO - Associate DAOs with respective collections
//Create new DAO here for new collection
func InitDAO() {
	// PlayersDAO = dao.GetMongoDAO("players")
	// TeamsDAO = dao.GetMongoDAO("teams")
	// BiddingDAO = dao.GetMongoDAO("bidding")
	UsersDAO = GetMongoDAO("users")
	LoginDAO = GetMongoDAO("login")
	CdnDAO = GetMongoDAO("cdnmetalist")
}
