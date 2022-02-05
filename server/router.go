package server

func (s *Server) InitRouter() {
	s.engine.GET("/video/download", s.DownloadFile())
	s.engine.POST("/video/init_download", s.DownloadNew())
	s.engine.GET("/tmp", s.TestController())
	s.engine.GET("/video/info", s.GetYoutubeVideoFilesInfoNew())
}
