package main

import (
	"log"
	"net/http"

	gateway "github.com/soundise/soundise/api"
)

func main() {
	var gateway gateway.WebsocketGateway = gateway.Init()

	staticServer := http.FileServer(http.Dir("static/"))
	http.Handle("/", staticServer)
	http.HandleFunc("/api", testFunc)
	http.HandleFunc("/ws", gateway.HandleConnection)

	log.Println("Listening...")
	log.Fatal(http.ListenAndServe("0.0.0.0:8000", nil))
}

func testFunc(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"message\": \"Hello World\"}"))
}
