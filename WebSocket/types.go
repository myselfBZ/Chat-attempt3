package pkg

import "github.com/gorilla/websocket"

type Room struct{
    ID      string 
    Client  []Client 
}


type Client struct{
    conn    *websocket.Conn
    RoomId  string 
}


type Message struct{
    RoomId string
    Text   string
}
