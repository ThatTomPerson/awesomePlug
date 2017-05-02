package main

import (
	"golang.org/x/net/websocket"
	"log"
	"fmt"
)

var origin = "https://plug.dj"
var url = "wss://godj.plug.dj:443/socket"

func main () {
	ws, err := websocket.Dial(url, "", origin)

	if err != nil {
		log.Fatal(err)
	}

	var msg = make([]byte, 512)
	_, err = ws.Read(msg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Receive: %s\n", msg)
}