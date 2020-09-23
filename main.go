//go:generate go run -tags generate gen.go

package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/zserge/lorca"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"os/signal"
	"strings"

	"github.com/go-cmd/cmd"
)

var upgrader = websocket.Upgrader{}

const SERVER_URL string = "127.0.0.1:8090"

func serveHome(w http.ResponseWriter, r *http.Request) {
	index, err := ioutil.ReadFile("www/index.html")
	if err != nil {
		log.Fatalf("Error opening index file : %s", err)
	}
	fmt.Fprintf(w, "%s", index)
}

func main() {
	// Enable line numbers in logging
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	//log.SetFlags(0)
	//debug := true
	//w := webview.New(debug)
	//defer w.Destroy()
	//w.SetTitle("Minimal webview example")
	//w.SetSize(800, 600, webview.HintNone)
	//w.Navigate("https://en.m.wikipedia.org/wiki/Main_Page")
	//w.Run()

	ui, err := lorca.New("", "", 700, 700)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()
	cmdOptions := cmd.Options{
		Buffered:  false,
		Streaming: true,
	}
	hugoServer := HugoServer{
		cmd: cmd.NewCmdOptions(cmdOptions, "hugo", strings.Split("server /users/matthieu/apps/www.hby.io", " ")...),
	}

	http.HandleFunc("/", serveHome)
	go http.ListenAndServe(SERVER_URL, nil)

	// A simple way to know when UI is ready (uses body.onload event in JS)
	ui.Bind("startserver", func() {
		log.Println("Run called")

		// Print STDOUT and STDERR lines streaming from Cmd
		doneChan := make(chan struct{})
		go func() {
			defer close(doneChan)
			// Done when both channels have been closed
			// https://dave.cheney.net/2013/04/30/curious-channels
			for hugoServer.cmd.Stdout != nil || hugoServer.cmd.Stderr != nil {
				select {
				case line, open := <-hugoServer.cmd.Stdout:
					if !open {
						hugoServer.cmd.Stdout = nil
						continue
					}
					fmt.Println(line)
				case line, open := <-hugoServer.cmd.Stderr:
					if !open {
						hugoServer.cmd.Stderr = nil
						continue
					}
					fmt.Fprintln(os.Stderr, line)
				}
			}
		}()

		// Run Hugo server
		fmt.Printf("error 1 starting %s\n", hugoServer.cmd.Status().Error)
		// Start a long-running process, capture stdout and stderr
		<-hugoServer.cmd.Start() // non-blocking
		fmt.Printf("error starting %s\n", hugoServer.cmd.Status().Error)
		// Block waiting for command to exit, be stopped, or be killed
		<-doneChan
	})

	ui.Bind("serverstatus", func() string {
		log.Println("Status called")
		return fmt.Sprintf("%#v", hugoServer.cmd.Status())
	})

	ui.Bind("serverversion", func() string {
		log.Println("Status called")
		serverversion, err := exec.Command("hugo", "version").Output()
		if err != nil {
			log.Printf("Command finished with error: %v", err)
		}
		return fmt.Sprintf("%v", string(serverversion))
	})

	// A simple way to know when UI is ready (uses body.onload event in JS)
	ui.Bind("stopserver", func() string {
		log.Println("Stop called")
		err := hugoServer.cmd.Stop()
		if err != nil {
			log.Fatalf("Erro trying to stop server : %s", err)
		}

		return "Server stopped"
	})

	// Load HTML.
	ui.Load("http://" + SERVER_URL)

	// You may use console.log to debug your JS code, it will be printed via
	// log.Println(). Also exceptions are printed in a similar manner.
	ui.Eval(`
		console.log("Hello, world!");
		console.log('Multiple values:', [1, false, {"x":5}]);
	`)

	// Wait until the interrupt signal arrives or browser window is closed
	sigc := make(chan os.Signal)
	signal.Notify(sigc, os.Interrupt)
	select {
	case <-sigc:
	case <-ui.Done():
	}

	// Do not let this process hang
	defer func() {
		err := hugoServer.cmd.Stop()
		if err != nil {
			log.Fatalf("Error killing hugo server cmd : %s", err)
		}
	}()

	log.Println("exiting...")
}
