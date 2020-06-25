package handlers

import (
	"encoding/json"
	socketio "github.com/googollee/go-socket.io"
	"net/http"
)

type Device struct {
	ID                  string `json:"id"`
	EmailAddresses		[]string `json:"email_addresses"`
	Conn                socketio.Conn `json:"conn"`
	DeviceDetails       string `json:"device_details"`
	DeviceRemoteAddress string `json:"device_remote_address"`
}

type DeviceArray struct {
	Devices []Device `json:"devices"`
}

var DevicesList DeviceArray

func DeviceHandler(w http.ResponseWriter, r *http.Request) {
	res, _ := json.Marshal(DevicesList)
	sendResponse(w, http.StatusOK, res)
	return
}
