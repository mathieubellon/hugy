package main

import (
	"bufio"
	"log"
	"os/exec"
	"runtime"
	"strings"
	"syscall"
)

func runHugoServer() {

	arguments := "version"
	args := strings.Split(arguments, " ")
	// The command you want to run along with the argument
	if runtime.GOOS == "windows" {
		cmd = exec.Command("./resources/darwin/hugo", args...)
	} else {
		cmd = exec.Command("./resources/darwin/hugo", args...)
	}

	// Get a pipe to read from standard out
	r, _ := cmd.StdoutPipe()

	// Use the same pipe for standard error
	cmd.Stderr = cmd.Stdout

	// Make a new channel which will be used to ensure we get all output
	done := make(chan struct{})

	// Create a scanner which scans r in a line-by-line fashion
	scanner := bufio.NewScanner(r)

	// Use the scanner to scan the output line by line and log it
	// It's running in a goroutine so that it doesn't block
	go func() {

		// Read line by line and process it
		for scanner.Scan() {
			line := scanner.Text()
			log.Print(line)
		}

		// We're all done, unblock the channel
		done <- struct{}{}

	}()

	// Start the command and check for errors
	err := cmd.Start()
	log.Print(err)

	// Wait for all output to be processed
	<-done

	// Wait for the command to finish
	err = cmd.Wait()
	log.Print(err)
}

func stopHugoServer(cmd *exec.Cmd) {
	err := cmd.Process.Signal(syscall.SIGTERM)
	if err != nil {
		log.Printf("Could not gracefully stop app %s", err)
	}
}

func killHugoServer(cmd *exec.Cmd) {
	log.Printf("Try to kill")
	err := cmd.Process.Kill()
	if err != nil {
		log.Fatalf("Could not force kill app %s", err)
	}
}
