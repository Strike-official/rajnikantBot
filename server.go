package main

import (
	"io"
	"log"
	"os"

	"github.com/Strike-official/rajnikantBot/configmanager"
	"github.com/Strike-official/rajnikantBot/internal/model"
	pkg "github.com/Strike-official/rajnikantBot/pkg/mongodb"
	"github.com/gin-gonic/gin"
)

func main() {
	// Read Config
	err := configmanager.InitAppConfig("configs/config.json")
	if err != nil {
		log.Fatal("[startAPIs] Failed to start APIs. Error: ", err)
	}
	model.Conf = configmanager.GetAppConfig()

	// Init LogFile
	logFile := initLogger(model.Conf.LogFilePath)

	// Init Routes
	gin.DefaultWriter = io.MultiWriter(logFile, os.Stdout)
	router := gin.Default()
	initializeRoutes(router)

	// Initialize mongodb
	pkg.Init()

	// Start serving the application
	err = router.Run(model.Conf.Port)
	if err != nil {
		log.Fatal("[startAPIs] Failed to start APIs. Error: ", err)
	}
}

func initLogger(filePath string) *os.File {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	return file
}
