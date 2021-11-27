package server

import (
	"downloader/config"
	"downloader/storage"
	"downloader/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	store  *storage.Store
}

func RunServer() {
	conf := config.GetConfig()
	store, err := storage.New()
	if err != nil {
		utils.Logger.Error("Create storage error")
		utils.Logger.Fatal(err)
	}
	server := Server{engine: gin.Default(), store: store}
	server.InitRouter()
	err = server.engine.Run(fmt.Sprintf(":%s", conf.App.Port))
	if err != nil {
		utils.Logger.Fatal(err)
	}
}
