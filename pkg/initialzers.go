package pkg

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	MsgChan = make(chan Message)
	H       *Handler
)

func HandlerInit() {
    if H == nil{
        H = &Handler{
            Clients:  make(map[Client]bool),
            Upgrader: &websocket.Upgrader{
                CheckOrigin: func(r *http.Request)bool{
                    return true
                },
            },
        }
    }
    return 
}
