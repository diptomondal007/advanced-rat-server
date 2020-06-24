package handlers

import "net/http"

func sendResponse(w http.ResponseWriter,statusCode int, data []byte){
	w.Header().Set("Content-Type","application/json")
	w.WriteHeader(statusCode)
	w.Write(data)
}


