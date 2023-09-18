package main

import (
	"log"
	"net/http"

	"github.com/truemanluo/gstorage/handler"
)

func main() {
	http.HandleFunc("/upload", handler.UploadFile)
	http.HandleFunc("/info", handler.FileInfo)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
