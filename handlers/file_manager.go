package handlers

import (
	"encoding/json"
	"net/http"
)

var FmList interface{}

func GetFileList(w http.ResponseWriter, r *http.Request){
	res , _ :=json.Marshal(FmList)
	sendResponse(w, http.StatusOK, res)
}
