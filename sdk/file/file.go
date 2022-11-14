package file

import (
	"mime/multipart"
	"strings"

	"github.com/gin-gonic/gin"
)

type File struct {
	Content multipart.File
	Meta    *multipart.FileHeader
}

func Init(ctx *gin.Context, key string) (*File, error) {
	content, meta, err := ctx.Request.FormFile(key)
	if err != nil {
		return nil, err
	}

	file := &File{
		Content: content,
		Meta:    meta,
	}

	return file, nil
}

func (f *File) SetFileName(newName string) {
	fileName := strings.Split(f.Meta.Filename, ".")
	fileName[0] = newName
	f.Meta.Filename = strings.Join(fileName, ".")
}

func (f *File) IsImage() bool {
	return strings.Contains(f.Meta.Header.Get("Content-Type"), "image")
}

func GetFileNameFromURL(url string) string {
	fileName := strings.Split(url, "/")[len(strings.Split(url, "/"))-1]

	return fileName
}
