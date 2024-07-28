package main

import (
	"log"
	"net/http"

	ws "github.com/myselfBZ/Chat2/WebSocket"
)


func main()  {
    h := ws.NewHandler()
    mux := http.NewServeMux()
    mux.HandleFunc("/rooms/{roomId}", h.JoinRoom)
    mux.HandleFunc("POST /rooms", h.CreateRoom)
    log.Println("Listening...")
    log.Fatal(http.ListenAndServe(":8080", mux))
}
