package server

import (
	"downloader/server/controllers/youtube"
)

func (s *Server) InitRouter() {
	s.engine.GET("/video/info", youtubeController.GetYoutubeVideoFilesInfo)
	s.engine.GET("/video/init_download", youtubeController.Download)
	s.engine.GET("/video/download", youtubeController.GetDownloadedFile)
	s.engine.GET("/tmp", s.TestController())
}
