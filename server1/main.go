package main

import (
	"flag"
	"log"
	"net/http"
	"github.com/gorilla/websocket"
)


var addr = flag.String("addr", "localhost:8080", "http service address")
var upgrader = websocket.Upgrader{}

func echo(w http.ResponseWriter, r *http.Request){
	upgrader.CheckOrigin = func (r *http.Request) bool {
		return true
	}
	c, err := upgrader.Upgrade(w, r, nil)


	if err != nil {
		log.Print("upgrade: ", err)
		return
	}

	defer c.Close()

	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read: ", err)
			return
		}

		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write: ", err)
			break
		}
	}
}

func main(){
	http.HandleFunc("/",echo)
	http.ListenAndServe(*addr, nil)
}