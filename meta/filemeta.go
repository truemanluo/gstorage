package meta

import (
	"errors"
)

type FileMetaInfo struct {
	ID         string
	Name       string
	Location   string
	Size       int64
	ModTime    string
	UploadTime string
	IsDir      bool
	Extension  string
}

var fileMetas map[string]*FileMetaInfo

func init() {
	fileMetas = make(map[string]*FileMetaInfo)
}

func GetFileMataByID(ID string) (*FileMetaInfo, error) {
	if val, ok := fileMetas[ID]; ok {
		return val, nil
	}
	return nil, errors.New("ID doesn't exist")
}

func (f *FileMetaInfo) UpdateFileMeta() {
	fileMetas[f.ID] = f
}
