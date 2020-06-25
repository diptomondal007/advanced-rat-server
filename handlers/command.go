package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

func CommandHandler(w http.ResponseWriter, r *http.Request) {
	keys, ok := r.URL.Query()["name"]
	if !ok && len(keys) < 1{
		_ = json.NewEncoder(w).Encode("Url param missing!")
		return
	}
	key := keys[0]
	log.Println(key)
	for _, v := range DevicesList.Devices{
		switch key {
		case "record":
			v.Conn.Emit("command", key, 20)
		case "fm":
			subCommand, ok := r.URL.Query()["subcommand"]
			if !ok && len(subCommand) <0{
				_ = json.NewEncoder(w).Encode("Url param missing!")
				return
			}
			path, ok:= r.URL.Query()["path"]
			if !ok && len(subCommand) <0{
				_ = json.NewEncoder(w).Encode("Url param missing!")
				return
			}
			v.Conn.Emit("command", key, subCommand[0], path[0])

		default:
			v.Conn.Emit("command", key)
		}
	}
	res , _:= json.Marshal(SuccessMessage{Message:"command successful"})
	sendResponse(w, http.StatusOK, res)
	return
}
