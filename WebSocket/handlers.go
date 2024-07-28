package pkg

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/myselfBZ/Chat2/errs"
)


type Handler struct{
    Upgrader    *websocket.Upgrader
    Clients     map[Client]bool
}


func(h *Handler) HandleConn(w http.ResponseWriter, r *http.Request) {
    conn, err := h.Upgrader.Upgrade(w, r, nil) 
    if err != nil {
        errs.ConnError(w)
    }
    client := Client{
        conn: conn,
    }
    h.Clients[client] = true 

}

func (h *Handler) JoinRoom(w http.ResponseWriter, r *http.Request)  {
    roomId :=  r.PathValue("roomId")
    
    conn, err := h.Upgrader.Upgrade(w, r, nil)
    if err != nil{
        errs.ConnError(w)
    }

    var client *Client

    for cl := range h.Clients{
        if  cl.conn == conn{
            client = &cl
            break 
        } 
    }
    
    if client == nil {
        json.NewEncoder(w).Encode(map[string]string{ "error" : "connection not found"})
    }

    client.RoomId = roomId

    h.Clients[*client] = true

    w.WriteHeader(http.StatusOK)
}


func HandleMesages(){
    msg := <-MsgChan
    for {
        for cl := range H.Clients{
        
            if cl.RoomId == msg.RoomId{
                err := cl.conn.WriteJSON(msg)
                if err != nil{
                    delete(H.Clients, cl)
                }
            }
        } 
    }
}


func readMessage(c *Client) {
    if c.RoomId != "" {
        for {
            var msg Message 
            err := c.conn.ReadJSON(&msg)
            if err != nil {
                delete(H.Clients, *c)
            }
            msg.RoomId = c.RoomId
            MsgChan <- msg

        }
    }
}



