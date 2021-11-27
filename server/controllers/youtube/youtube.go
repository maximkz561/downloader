package youtubeController

import (
	"downloader/utils"
	"downloader/youtube"
	"github.com/gin-gonic/gin"
	"path/filepath"
)

func GetYoutubeVideoFilesInfo(c *gin.Context) {
	params := c.Request.URL.Query()
	link := params.Get("link")
	if link == "" {
		c.JSON(422, gin.H{"success": "false", "error": "link param is required"})
		return
	}
	video, err := youtube.NewVideo(link)
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

func Download(c *gin.Context) {
	params := c.Request.URL.Query()
	fileId := params.Get("file_id")
	if fileId == "" {
		c.JSON(422, gin.H{"success": "false", "error": "link param is required"})
		return
	}
	formatId := params.Get("format_id")
	if formatId == "" {
		c.JSON(422, gin.H{"success": "false", "error": "link param is required"})
		return
	}

	filename, err := youtube.DownloadFile(fileId, formatId)
	if err != nil {
		utils.Logger.Error()
		return
	}

	c.JSON(200, gin.H{"success": "true", "filename": filename})
}

func GetDownloadedFile(c *gin.Context) {
	DOWNLOADS_PATH := "."

	params := c.Request.URL.Query()
	filename := params.Get("filename")
	if filename == "" {
		c.JSON(422, gin.H{"success": "false", "error": "filename param is required"})
		return
	}
	targetPath := filepath.Join(DOWNLOADS_PATH, filename)

	//Seems this headers needed for some browsers (for example without this headers Chrome will download files as txt)
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", "application/octet-stream")
	c.File(targetPath)
}
