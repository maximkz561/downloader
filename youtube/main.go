package youtube

import (
	"downloader/utils"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"net/url"
	"strings"
)

type Video struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Files []File `json:"files"`
}

func NewVideo(link string) (*Video, error) {
	videoUrl, err := url.Parse(link)
	if err != nil {
		return nil, err
	}

	urlQuery := videoUrl.Query()
	// if link like https://www.youtube.com/watch?v=KJFUEN7fnds
	videoID := urlQuery.Get("v")
	if videoID == "" {
		// if link like https://youtu.be/KJFUEN7fnds
		videoID = strings.TrimPrefix(videoUrl.Path, "/")
	}
	return &Video{Id: videoID, Files: []File{}}, nil
}

func (v *Video) CollectInfo() error {
	var files []File
	stdout, stderr, err := utils.Shellout(fmt.Sprintf("yt-dlp -j -- %s", v.Id))

	if err != nil {
		utils.Logger.Error(err, stderr)
		return err
	}

	formats := gjson.Get(stdout, "formats")
	title := gjson.Get(stdout, "title")

	err = json.Unmarshal([]byte(formats.String()), &files)
	v.Files = files
	v.Title = title.String()
	return nil
}
