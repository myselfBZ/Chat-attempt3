package ws 

import "github.com/gorilla/websocket"

type Room struct{
    Name    string  `json:"name"`
    ID      int  `json:"id"`
}


type Client struct{
    conn    *websocket.Conn
    RoomId  int 
}


type Message struct{
    RoomId  int 
    Text   string   `json:"text"`
}
