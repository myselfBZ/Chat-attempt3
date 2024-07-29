package main

import (
	"log"
	"net/http"

	ws "github.com/myselfBZ/Chat2/WebSocket"
)

func init()  {
    ws.Rooms = make([]ws.Room, 128)
}

func main()  {
    h := ws.NewHandler()
    mux := http.NewServeMux()
    mux.HandleFunc("/rooms", h.ListRooms)
    mux.HandleFunc("/rooms/{roomId}", h.JoinRoom)
    mux.HandleFunc("POST /rooms", h.CreateRoom)
    go h.WriteMesages()
    log.Println("Listening...")
    log.Fatal(http.ListenAndServe(":8080", mux))
}
