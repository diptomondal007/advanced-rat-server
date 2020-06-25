package handlers

import "net/http"

type SuccessMessage struct {
	Message string `json:"message"`
}

func sendResponse(w http.ResponseWriter,statusCode int, data []byte){
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(statusCode)
	w.Write(data)
}


