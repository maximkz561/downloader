package youtube

import (
	"downloader/utils"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
)

func DownloadFile(fileId, formatId string) (string, error) {
	var files []File
	stdout, stderr, err := utils.Shellout(fmt.Sprintf("yt-dlp -j %s", fileId))

	if err != nil {
		utils.Logger.Error(err, stderr)
		return "", err
	}

	formats := gjson.Get(stdout, "formats")
	title := gjson.Get(stdout, "title")

	err = json.Unmarshal([]byte(formats.String()), &files)

	var fileToDownload File
	for _, file := range files {
		if file.FormatId == formatId {
			fileToDownload = file
			break
		}
	}
	if fileToDownload == (File{}) {
		return "No file with such formatId", err
	}

	filename := fmt.Sprintf("%s-%s-%s.%s", title.String(), fileId, formatId, fileToDownload.Ext)
	fmt.Println(filename)
	go func() {
		_, err = Download(ConvertFile(fileToDownload), fileId, filename)
		if err != nil {
			utils.Logger.Error(err)
		}
	}()

	return filename, nil

}
