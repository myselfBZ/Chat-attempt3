package ws

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/myselfBZ/Chat2/errs"
)


type Handler struct{
    Upgrader    *websocket.Upgrader
    Clients     map[Client]bool
}

func (h *Handler) CreateRoom(w http.ResponseWriter, r *http.Request){
    var room Room
    if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
        errs.JSONError(w)
    }
    room.ID = len(Rooms) + 1
    Rooms = append(Rooms, room)
    //Broadcasting the newly created room
    for c := range h.Clients{
        err := c.conn.WriteJSON(room)
        if err != nil {
            delete(h.Clients, c)
            errs.ConnError(w)
        }
    }
}




func (h *Handler) JoinRoom(w http.ResponseWriter, r *http.Request)  {
    roomId :=  r.PathValue("roomId")
    id, err := strconv.Atoi(roomId)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return 
    }
    
    conn, err := h.Upgrader.Upgrade(w, r, nil)
    client := Client{
        conn: conn,
    }
    h.Clients[client] = true 

    log.Println("Client has connected: ", conn.RemoteAddr())
    if err != nil{
        errs.ConnError(w)
    }
    
    

    client.RoomId = id 
    

    go h.readMessage(&client)
    go h.writeMesages()

}


func (h *Handler) writeMesages(){
    msg := <-MsgChan
    for {
        for cl := range h.Clients{
        
            if cl.RoomId == msg.RoomId{
                err := cl.conn.WriteJSON(msg)
                if err != nil{
                    delete(h.Clients, cl)
                }
            }
        } 
    }
}


func (h *Handler) readMessage(c *Client) {
    if c.RoomId != 0 {
        for {
            var msg Message 
            err := c.conn.ReadJSON(&msg)
            if err != nil {
                delete(h.Clients, *c)
            }
            msg.RoomId = c.RoomId
            MsgChan <- msg

        }
    }
}



