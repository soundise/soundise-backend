package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
)

type WebsocketGateway struct {
	upgrader websocket.Upgrader
}

func (gateway WebsocketGateway) New(router *mux.Router) *WebsocketGateway {
	router.HandleFunc("/ws", gateway.HandleConnection)

	gateway.upgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true }, // TODO: Remove this in the future (security reason)
	}

	return &gateway
}

func (gateway WebsocketGateway) HandleConnection(w http.ResponseWriter, r *http.Request) {

	conn, err := gateway.upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println("There was an error upgrading the websocket connection", err)
	}
	log.Println("Gateway connection established")
	defer conn.Close()

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Message", messageType, message)
			break
		}
	}

	log.Println("Connection closed")

}
