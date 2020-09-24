package main

import (
	"fmt"
	"github.com/go-cmd/cmd"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

type HugoServer struct {
	arguments string
	cmd       *cmd.Cmd
	ws        *websocket.Conn
	isRunning bool
	PID       int
}

func HugoBinary() string {
	binName := "hugo"
	//path, err := exec.LookPath("hugo")
	//if err != nil {
	//	log.Printf(" Error finding binary : %s", err)
	//}

	currentDir, _ := os.Getwd()
	hugoAbsPath := currentDir + "bin/" + binName
	log.Printf("Hugo path %s", hugoAbsPath)
	return hugoAbsPath
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
	wg := new(sync.WaitGroup)
	wg.Add(2)
	cmdOptions := cmd.Options{
		Buffered:  false,
		Streaming: true,
	}
	hs.cmd = cmd.NewCmdOptions(cmdOptions, HugoBinary(), strings.Split("server /users/matthieu/apps/www.hby.io", " ")...)
	go func() {
		log.Println("Run called")
		statusChan := hs.cmd.Start() // non-blocking
		hs.isRunning = true
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
					hs.ws.WriteMessage(websocket.TextMessage, []byte(line))
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
		// Start a long-running process, capture stdout and stderr
		<-statusChan
		// Block waiting for command to exit, be stopped, or be killed
		<-doneChan
		wg.Done()
	}()

	// Print last line of stdout every 2s
	go func() {
		for {
			hs.PID = hs.cmd.Status().PID
			if hs.PID > 0 {
				fmt.Printf("$$$$$$$$$$$$$$$$$$$$$$$\n\n %s", hs.PID)
				hs.ws.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("Server status found : %#v", hs.cmd.Status())))
				break
			}
		}
		wg.Done()
	}()

	wg.Wait()

}

func (hs *HugoServer) status() string {
	log.Println("Status called")
	return fmt.Sprintf("%#v", hs.cmd.Status())
}

func (hs *HugoServer) version() string {
	log.Println("Status called")
	wd, _ := os.Getwd()
	return fmt.Sprintf("%v", wd)
}

func (hs *HugoServer) stop() string {
	log.Println("Stop called")
	err := hs.cmd.Stop()
	if err != nil {
		log.Fatalf("Erro trying to stop server : %s", err)
	}
	hs.isRunning = false
	return "Server stopped"
}
