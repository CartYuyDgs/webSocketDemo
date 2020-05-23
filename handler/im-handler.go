package handler

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type ImHandler struct {
}

func NewImHandler() *ImHandler {
	return &ImHandler{}
}

func (h *ImHandler) Handler(w http.ResponseWriter, r *http.Request) {
	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool {
		return true
	}}).Upgrade(w, r, nil)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	go h.Read(conn)
	return
}

func (h *ImHandler) Read(link *websocket.Conn) {
	defer link.Close()

	for {
		_, message, err := link.ReadMessage()
		if err != nil {
			log.Printf("Error %s", err)
		}
		log.Printf("message: %s", message)
	}
}
