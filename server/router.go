package server

import (
	"downloader/server/controllers/youtube"
)

func (s *Server) InitRouter() {
	s.engine.GET("/video/info", youtubeController.GetYoutubeVideoFilesInfo)
	s.engine.GET("/video/download", youtubeController.GetDownloadedFile)
	s.engine.GET("/video/download/new", s.DownloadFile())
	s.engine.GET("/video/init_download/new", s.DownloadNew())
	s.engine.GET("/tmp", s.TestController())
	s.engine.GET("/video/info/new", s.GetYoutubeVideoFilesInfoNew())
}
