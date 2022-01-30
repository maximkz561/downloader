package server

import (
	"downloader/utils"
	"downloader/youtube"
	"github.com/gin-gonic/gin"
)

func (s *Server) GetYoutubeVideoFilesInfoNew() gin.HandlerFunc {

	type request struct {
		Link string `form:"link" binding:"required"`
	}

	return func (c *gin.Context) {
		req := &request{}
		if err := c.BindQuery(req); err != nil {
			c.JSON(400, gin.H{"success": false})
			return
		}

		video, err := youtube.NewVideo(req.Link)
		if err != nil {
			c.JSON(400, gin.H{"success": "false", "error": "error occurs during parsing video files"})
			return
		}
		err = video.CollectInfo()
		if err != nil {
			c.JSON(400, gin.H{"success": "false", "error": "error occurs during parsing video files"})
			return
		}

		c.JSON(200, video)

	}
}


func (s *Server) DownloadNew() gin.HandlerFunc {

	type request struct {
		FileId string `form:"file_id" binding:"required"`
		FormatId string `form:"format_id" binding:"required"`
	}

	return func (c *gin.Context) {
		req := &request{}
		if err := c.BindQuery(req); err != nil {
			c.JSON(400, gin.H{"success": false})
			return
		}

		file, err := youtube.DownloadFile(s.store, req.FileId, req.FormatId)
		if err != nil {
			utils.Logger.Error()
			return
		}

		c.JSON(200, gin.H{"success": "true", "file_id": file.Id})

	}
}