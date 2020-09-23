package main

import (
	"github.com/go-cmd/cmd"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type HugoServer struct {
	arguments string
	cmd       *cmd.Cmd
	ws        *websocket.Conn
}

func (hs *HugoServer) openWs(w http.ResponseWriter, r *http.Request) {

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error while upgrading to WS in Hugo server : %s", err)
		return
	}

	hs.ws = ws

	for {
		// Read new message
		_, msg, err := hs.ws.ReadMessage()
		if err != nil {
			log.Printf("connection lost when read %v", err)
		}
		log.Printf("MESSAGE incoming : %s\n", msg)
	}
}
