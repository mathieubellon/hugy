package main

import (
	"fmt"
	"github.com/go-cmd/cmd"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
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

func (hs *HugoServer) start() {

	log.Println("Run called")
	cmdOptions := cmd.Options{
		Buffered:  false,
		Streaming: true,
	}
	hs.cmd = cmd.NewCmdOptions(cmdOptions, "hugo", strings.Split("server /users/matthieu/apps/www.hby.io", " ")...)

	// Print STDOUT and STDERR lines streaming from Cmd
	doneChan := make(chan struct{})
	go func() {
		defer close(doneChan)
		// Done when both channels have been closed
		// https://dave.cheney.net/2013/04/30/curious-channels
		for hs.cmd.Stdout != nil || hs.cmd.Stderr != nil {
			select {
			case line, open := <-hs.cmd.Stdout:
				if !open {
					hs.cmd.Stdout = nil
					continue
				}
				fmt.Println(line)
			case line, open := <-hs.cmd.Stderr:
				if !open {
					hs.cmd.Stderr = nil
					continue
				}
				fmt.Fprintln(os.Stderr, line)
			}
		}
	}()

	// Run Hugo server
	fmt.Printf("error 1 starting %s\n", hs.cmd.Status().Error)
	// Start a long-running process, capture stdout and stderr
	<-hs.cmd.Start() // non-blocking
	fmt.Printf("error starting %s\n", hs.cmd.Status().Error)
	// Block waiting for command to exit, be stopped, or be killed
	<-doneChan
}

func (hs *HugoServer) status() string {
	log.Println("Status called")
	return fmt.Sprintf("%#v", hs.cmd.Status())
}

func (hs *HugoServer) version() string {
	log.Println("Status called")
	serverversion, err := exec.Command("hugo", "version").Output()
	if err != nil {
		log.Printf("Command finished with error: %v", err)
	}
	return fmt.Sprintf("%v", string(serverversion))
}

func (hs *HugoServer) stop() string {
	log.Println("Stop called")
	err := hs.cmd.Stop()
	if err != nil {
		log.Fatalf("Erro trying to stop server : %s", err)
	}

	return "Server stopped"
}
