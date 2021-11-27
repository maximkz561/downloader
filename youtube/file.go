package youtube

import (
	"downloader/utils"
	"fmt"
)

type File struct {
	Filesize   int         `json:"filesize,omitempty"` // 134 - 12324442
	FormatId   string      `json:"format_id"`          // 249
	Fps        interface{} `json:"fps,omitempty"`      // 25, might be null
	Url        string      `json:"url,omitempty"`
	Ext        string      `json:"ext,omitempty"`        // mp4
	Vcodec     string      `json:"vcodec,omitempty"`     // avc1.4d4015, might be "none" as string
	Acodec     string      `json:"acodec,omitempty"`     // vp9, might be "none" as string
	AudioExt   string      `json:"audio_ext,omitempty"`  // webm, might be "none" as string
	VideoExt   string      `json:"video_ext,omitempty"`  // mp4, might be "none" as string
	Format     string      `json:"format,omitempty"`     // 134 - 640x360 (360p)
	Resolution string      `json:"resolution,omitempty"` // 1920x1080 or "audio only"
}

func (f File) getFileType() FileTypeEnum {
	if f.Resolution == "audio only" {
		return FileType.Audio
	}
	return FileType.Video
}

type YoutubeFile interface {
	getFormatId() string
}

func Download(file YoutubeFile, videoId string, title string) (string, error) {
	_, stderr, err := utils.Shellout(fmt.Sprintf("yt-dlp -f %s %s -o '%s'", file.getFormatId(), videoId, title))

	if err != nil {
		return stderr, err
	}
	utils.Logger.Infoln(fmt.Sprintf("File %s %s stars downloading", file.getFormatId(), videoId))
	return "start download", err
}

type VideoFile struct {
	Filesize   int
	FormatId   string
	Ext        string
	Resolution string
	Fps        int
}

func (f VideoFile) getFormatId() string {
	return f.FormatId
}

type AudioFile struct {
	Filesize   int
	FormatId   string
	Ext        string
	Resolution string
}

func (f AudioFile) getFormatId() string {
	return f.FormatId
}

func ConvertFile(file File) YoutubeFile {
	switch file.getFileType() {
	case FileType.Audio:
		return AudioFile{
			Filesize:   file.Filesize,
			FormatId:   file.FormatId,
			Ext:        file.Ext,
			Resolution: file.Resolution,
		}
	case FileType.Video:
		return VideoFile{
			Filesize:   file.Filesize,
			FormatId:   file.FormatId,
			Ext:        file.Ext,
			Resolution: file.Resolution,
			Fps:        int(file.Fps.(float64)),
		}
	}
	return nil
}
