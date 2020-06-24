package handlers

import (
	"encoding/json"
	socketio "github.com/googollee/go-socket.io"
	"net/http"
)

type Device struct {
	ID                  string
	Conn                socketio.Conn
	DeviceDetails       string
	DeviceRemoteAddress string
}

var DeviceList []Device

func DeviceHandler(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal(DeviceList)
	sendResponse(w, http.StatusOK, res)
	return
}
