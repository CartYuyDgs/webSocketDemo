package handler

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"webSocketDemo/model"
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
		mes := &model.Message{}
		//new(model.Message)

		err = json.Unmarshal(message, mes)
		if err != nil {
			log.Printf("Error %s", err)
			continue
		}

		log.Printf("message: %s", message)
		switch mes.Type {
		case model.LOGIN:
			break
		case model.LOGOUT:
			break
		case model.SAY:
			break

		}
	}
}
