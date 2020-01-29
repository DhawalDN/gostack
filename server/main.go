package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/tidwall/gjson"

	"github.com/tidwall/sjson"

	"github.com/DhawalDN/gostack/server/dao"
	"github.com/DhawalDN/gostack/server/handlers"
	"github.com/DhawalDN/gostack/server/middleware"
	"github.com/DhawalDN/gostack/server/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// initConfig(os.Args[1])
	initConfig("/home/dhost/server/gostack/projectConfig.json")
	// fmt.Println(ProjectID)
	dao.InitDB("localhost", 27017)
	dao.InitDAO()
	// fmt.Println("Auction Routes Initiated")
	startServer()
}
func initConfig(filename string) {
	// fp, err1 := filepath.Abs(filepath.Dir(filename))
	// if err1 != nil {
	// log.Fatal(err1)
	// }

	// fmt.Println(fp)
	file, _ := ioutil.ReadFile(filename)
	var data interface{}
	err := json.Unmarshal([]byte(file), &data)
	if err != nil {
		log.Fatal("Error in Parsing the File")
	}
	projectConfigTemp := `{}`
	projectConfigTemp, _ = sjson.Set(projectConfigTemp, "projectConfig", data)
	projectConfig := gjson.Parse(projectConfigTemp)
	models.ProjectID = projectConfig.Get("projectConfig.projectId").String()
}

func startServer() {
	r := gin.Default()
	md := cors.DefaultConfig()
	md.AllowAllOrigins = true
	md.AllowHeaders = []string{"*"}
	md.AllowMethods = []string{"*"}
	r.Use(cors.New(md))
	r.Static("/"+models.ProjectID+"/files/", handlers.UploadPath)

	middleware.InitMiddleware(r)
	// http.HandleFunc("/upload", uploadFileHandler())
	// s := &http.Server{
	// 	Addr:    ":9000",
	// 	Handler: r,
	// }
	// s.ListenAndServe()
	r.Run(":9000")
}

// To build: use this command go build -ldflags="-s -w" main.go
