package handlers

import (
	"encoding/json"
	"net/http"
)


var MessageList interface{}
var ContactList interface{}
var CallLogs interface{}

func SmsListView(w http.ResponseWriter, r *http.Request){
	res , _ :=json.Marshal(MessageList)
	sendResponse(w, http.StatusOK, res)
	return
}

func CallLogsView(w http.ResponseWriter, r *http.Request){
	res , _ :=json.Marshal(CallLogs)
	sendResponse(w, http.StatusOK, res)
	return
}

func ContactListView(w http.ResponseWriter, r *http.Request){
	res , _ :=json.Marshal(ContactList)
	sendResponse(w, http.StatusOK, res)
	return
}
