package main

// test a change
import (
	"golang.org/x/net/websocket"
	"log"
	"fmt"
)

var origin = "https://plug.dj"
var url = "wss://echo.websocket.org"

func main () {
	ws, err := websocket.Dial(url, "", origin)

	if err != nil {
		log.Fatal(err)
	}

	message := []byte("hello, world!")
	_, err = ws.Write(message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Send: %s\n", message)

	var msg = make([]byte, 512)
	_, err = ws.Read(msg)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Receive: %s\n", msg)
}