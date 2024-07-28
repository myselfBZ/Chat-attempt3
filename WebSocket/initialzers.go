package ws 

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	MsgChan chan Message
    Rooms   []Room 
)

func NewHandler() *Handler {
    return &Handler{
            Clients:  make(map[Client]bool),
            Upgrader: &websocket.Upgrader{
                CheckOrigin: func(r *http.Request)bool{
                    return true
                },
            },
        }

}
