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
    broadcast   chan Message
}

func (h *Handler) CreateRoom(w http.ResponseWriter, r *http.Request){
    var room Room
    if err := json.NewDecoder(r.Body).Decode(&room); err != nil {
        errs.JSONError(w)
    }
    room.ID = len(Rooms) + 1
    Rooms = append(Rooms, room)
    log.Println(Rooms)
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

func (h *Handler) WriteMesages(){
    for msg := range h.broadcast{
        log.Println("I am about to be written",msg)
        for cl := range h.Clients{
            log.Println("Client's id:", cl.RoomId)
            log.Println("Message's  id:", msg.RoomId) 
            if cl.RoomId == msg.RoomId{
                err := cl.conn.WriteJSON(msg)
                if err != nil{
                    log.Println("Apperantly we have a problem")
                    delete(h.Clients, cl)
                    return 
                }
                log.Println("We have written that mf")
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
            log.Println("msg has been read",msg)
            msg.RoomId = c.RoomId
            log.Println("id has been set for the message")
            h.broadcast <- msg 
            log.Println("After writing to the channel")

        }
    }
}



