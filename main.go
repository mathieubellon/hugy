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
	"os/signal"
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

	ui, err := lorca.New("", "", 700, 1400)
	if err != nil {
		log.Fatal(err)
	}
	defer ui.Close()

	hugoServer := HugoServer{
		isRunning: false,
	}

	http.HandleFunc("/ws", hugoServer.openWs)
	http.HandleFunc("/", serveHome)
	go http.ListenAndServe(SERVER_URL, nil)

	// A simple way to know when UI is ready (uses body.onload event in JS)
	ui.Bind("startserver", func() string {
		if hugoServer.isRunning {
			log.Print("Server already running")
			return fmt.Sprint("Server already running")
		} else {
			hugoServer.start()
			return fmt.Sprint("Server exit")
		}
	})

	ui.Bind("serverstatus", func() string {
		return hugoServer.status()
	})

	ui.Bind("serverversion", func() string {
		return hugoServer.version()
	})

	// A simple way to know when UI is ready (uses body.onload event in JS)
	ui.Bind("stopserver", func() string {
		return hugoServer.stop()
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
