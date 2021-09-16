package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func logStdout(w http.ResponseWriter, req *http.Request) {
	msg := "Log into stdout"
	stdoutLogger := log.New(os.Stdout, "stdout", log.LstdFlags)
	stdoutLogger.Println(msg)
	_, _ = w.Write([]byte(msg))
}

func logStderr(w http.ResponseWriter, req *http.Request) {
	msg := "Log into stderr"
	stdoutLogger := log.New(os.Stderr, "stderr", log.LstdFlags)
	stdoutLogger.Println(msg)
	_, _ = w.Write([]byte(msg))
}

func main() {

	http.HandleFunc("/stdout", logStdout)
	http.HandleFunc("/stderr", logStderr)

	serverPort := 8080
	if port, err := strconv.Atoi(os.Getenv("SERVER_PORT")); err == nil {
		serverPort = port
	}

	err := http.ListenAndServe(fmt.Sprintf(":%d", serverPort), nil)
	if err != nil {
		panic(err)
	}
}
