package pkg

import "golang.org/x/net/websocket"


type Room struct{
    ID      int
    Client  []Client 
}


type Client struct{
    conn    *websocket.Conn
    RoomId  int 
}
