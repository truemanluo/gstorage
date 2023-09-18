package handler

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/truemanluo/gstorage/meta"
)

func FileInfo(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fileID := r.Form.Get("id")

	fileMeta, err := meta.GetFileMataByID(fileID)
	if err != nil {
		log.Printf("get fileinfo failed: %v", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(fileMeta)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Write(data)
}
