package ws

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/websocket"
	"github.com/myselfBZ/Chat2/errs"
)


type Handler struct{
    Upgrader    *websocket.Upgrader
    Clients     map[Client]bool
    broadcast   chan Message
}

func (h *Handler) CreateRoom(w http.ResponseWriter, r *http.Request){
    var room Room
    if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
        errs.JSONError(w)
    }
    room.ID = len(Rooms) + 1
    Rooms = append(Rooms, room)
}




func (h *Handler) JoinRoom(w http.ResponseWriter, r *http.Request)  {
    roomId :=  r.PathValue("roomId")
    id, err := strconv.Atoi(roomId)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return 
    }
    
    conn, err := h.Upgrader.Upgrade(w, r, nil)
    if err != nil {
        w.WriteHeader(http.StatusBadRequest)
        return 
    }
    client := Client{
        conn: conn,
        RoomId: id,
    }
    h.Clients[client] = true 

    h.readMessage(&client)

}

func ListRooms(w http.ResponseWriter, r *http.Request) {
    
    json.NewEncoder(w).Encode(Rooms) 

}

func (h *Handler) WriteMesages(){
    for msg := range h.broadcast{
        for cl := range h.Clients{
            if cl.RoomId == msg.RoomId{
                err := cl.conn.WriteJSON(msg)
                if err != nil{
                    delete(h.Clients, cl)
                    return 
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
                return 
            }
            msg.RoomId = c.RoomId
            h.broadcast <- msg 

        }
    }
}



