package youtube

type FileTypeEnum string

var FileType = struct {
	Video FileTypeEnum
	Audio FileTypeEnum
}{
	Video: "Video",
	Audio: "Audio",
}
