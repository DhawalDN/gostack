package main

import (
	"github.com/DhawalDN/gostack/server/dao"
	"github.com/DhawalDN/gostack/server/handlers"
	"github.com/DhawalDN/gostack/server/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	dao.InitDB("localhost", 27017)
	dao.InitDAO()

	// fmt.Println("Auction Routes Initiated")
	startServer()
}

func startServer() {
	r := gin.Default()
	md := cors.DefaultConfig()
	md.AllowAllOrigins = true
	md.AllowHeaders = []string{"*"}
	md.AllowMethods = []string{"*"}
	r.Use(cors.New(md))
	r.Static("/files/", handlers.UploadPath)
	middleware.InitMiddleware(r)
	// http.HandleFunc("/upload", uploadFileHandler())
	// s := &http.Server{
	// 	Addr:    ":9000",
	// 	Handler: r,
	// }
	// s.ListenAndServe()
	r.Run(":9000")
}
