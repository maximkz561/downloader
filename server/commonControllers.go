package server

import (
	"downloader/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"path/filepath"
)

func (s *Server) DownloadFile() gin.HandlerFunc {

	type request struct {
		FileId string `form:"file_id" binding:"required"`
	}

	return func(c *gin.Context) {

		req := &request{}
		if err := c.BindQuery(req); err != nil {
			c.JSON(400, gin.H{"success": false})
			return
		}

		objectId, err := primitive.ObjectIDFromHex(req.FileId)
		if err != nil {
			c.JSON(400, gin.H{"success": false, "error": "wrong file_id"})
			return
		}

		DOWNLOADS_PATH := "./"

		file, err := s.store.FileRepository.Find(objectId)
		if err != nil {
			utils.Logger.Error(err)
			//возможно надо ошибку отдавать
			return
		}
		if file == nil {
			c.JSON(404, gin.H{"success": false, "error": "file not found"})
			return
		}

		if !file.Downloaded {
			c.JSON(202, gin.H{"success": true, "message": "file not ready"})
			return
		}

		targetPath := filepath.Join(DOWNLOADS_PATH, file.FileName)
		c.FileAttachment(targetPath, file.Title)

	}
}
