package handler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/truemanluo/gstorage/meta"
	"github.com/truemanluo/gstorage/utils"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	// Retrieve the file from the form data
	file, handler, err := r.FormFile("file")
	if err != nil {
		fmt.Println("error retrieving the file:", err)
		return
	}
	defer file.Close()

	// update file meta info
	fID, err := utils.FileMD5(file)
	if err != nil {
		log.Fatalf("get file md5 err %v", err)
	}
	fmeta := &meta.FileMetaInfo{
		ID:         fID,
		Location:   "testdata/" + handler.Filename,
		Name:       handler.Filename,
		Size:       handler.Size,
		IsDir:      false,
		Extension:  utils.FileExt(handler.Filename),
		UploadTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	fmeta.UpdateFileMeta()

	log.Println(fmeta.ID)

	// Create a new file on the server
	dst, err := os.Create(fmeta.Location)
	if err != nil {
		log.Println("error creating the file on the server:", err)
		return
	}
	defer dst.Close()

	// Copy the contents of the uploaded file to the newly created file
	_, err = io.Copy(dst, file)
	if err != nil {
		log.Println("error copying file contents:", err)
		return
	}

	fmt.Fprintf(w, "File uploaded successfully!")
}
