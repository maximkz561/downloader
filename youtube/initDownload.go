package youtube

import (
	"downloader/storage"
	"downloader/utils"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"go.mongodb.org/mongo-driver/bson"
)

func DownloadFile(store *storage.Store, fileId, formatId string) (*storage.File, error) {
	var files []File
	stdout, stderr, err := utils.Shellout(fmt.Sprintf("yt-dlp -j -- %s", fileId))

	if err != nil {
		utils.Logger.Error(err, stderr)
		return nil, err
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
		return nil, err
	}

	filename := fmt.Sprintf("%s-%s-%s.%s", title.String(), fileId, formatId, fileToDownload.Ext)
	fmt.Println(filename)
	file := &storage.File{Title: title.String(), Downloaded: false, FileId: fileId, FormatId: formatId, FileName: filename}
	createdFile, err := store.FileRepository.Create(file)
	if err != nil {
		return nil, err
	}
	go func() {

		_, err = Download(ConvertFile(fileToDownload), fileId, filename)
		if err != nil {
			utils.Logger.Error(err)
		}
		err = store.FileRepository.Update(file.Id, bson.D{{"downloaded", true}})
		if err != nil {
			return
		}
	}()

	return createdFile, nil

}
