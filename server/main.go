/**
 * @author Dhawal Dyavanpalli <dhawalhost@gmail.com>
 * @desc Created on 2020-08-31 1:03:01 pm
 * @copyright Crearosoft
 */
package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/crearosoft/corelib/loggermanager"

	"github.com/crearosoft/corelib/cachemanager"

	"github.com/tidwall/gjson"

	"github.com/tidwall/sjson"

	"github.com/dhawalhost/gostack/server/dao"
	"github.com/dhawalhost/gostack/server/handlers"
	"github.com/dhawalhost/gostack/server/middleware"
	"github.com/dhawalhost/gostack/server/models"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// initConfig(os.Args[1])
	initConfig("/home/dhost/server/gostack/projectConfig.json")
	// fmt.Println(ProjectID)
	dao.InitDB("localhost", 27017)
	dao.InitDAO()
	c := startCacheServer("./cachelogs")
	models.FC = c
	initAndStartServer(c)

}

// Start Cache Server from file,
func startCacheServer(fname string) *cachemanager.CacheHelper {
	var c cachemanager.CacheHelper
	// c.Setup(100000, 30*time.Second, 10*time.Second)
	c.Setup(100000, 24*7*time.Hour, 12*time.Hour)
	err := c.LoadFile(fname)

	if err != nil {
		loggermanager.LogError("No log file found. Creating new ...")
		c = *cachemanager.SetupCache()
	}
	// c.SetWithExpiration("hello", "world", 5*time.Second)
	// c.SaveFile("./logs")
	// c = *cachemanager.SetupCache()
	return &c
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
		loggermanager.LogError("Error in Parsing the File")
	}
	projectConfigTemp := `{}`
	projectConfigTemp, _ = sjson.Set(projectConfigTemp, "projectConfig", data)
	projectConfig := gjson.Parse(projectConfigTemp)
	models.ProjectID = projectConfig.Get("projectConfig.projectId").String()
}
func initAndStartServer(ci *cachemanager.CacheHelper) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		oscall := <-c
		loggermanager.LogError("system call: %+v", oscall)
		cancel()
	}()

	if err := startServer(ctx, ci); err != nil {
		loggermanager.LogError("failed to serve:+%v\n", err)
	}
}
func startServer(ctx context.Context, ci *cachemanager.CacheHelper) (err error) {
	r := gin.Default()
	md := cors.DefaultConfig()
	md.AllowAllOrigins = true
	md.AllowHeaders = []string{"*"}
	md.AllowMethods = []string{"*"}
	r.Use(cors.New(md))
	r.GET("/"+models.ProjectID+"/images/:directory/:imageId", handlers.DownloadHandler)
	// r.Static("/"+models.ProjectID+"/images/", handlers.UploadPath)

	middleware.InitMiddleware(r)
	// http.HandleFunc("/upload", uploadFileHandler())
	srv := &http.Server{
		Addr:    ":9000",
		Handler: r,
	}
	// s.ListenAndServe()
	go func() {
		// r.Run(":9000")
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%+s\n", err)
		}
	}()
	loggermanager.LogError("server started")

	<-ctx.Done()

	ci.SaveFile("./cachelogs")
	loggermanager.LogError("server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()
	if err = srv.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("server Shutdown Failed:%+s", err)
	}

	loggermanager.LogError("server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}

	return

}

// To build: use this command go build -ldflags="-s -w" main.go
