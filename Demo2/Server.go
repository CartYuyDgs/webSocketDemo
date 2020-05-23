package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
	"webSocketDemo/Demo2/impl"
)

var (
	upgrade = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func wsHandler(w http.ResponseWriter, r *http.Request) {

	var (
		wsConn *websocket.Conn
		err    error
		conn   *impl.Connection
		data   []byte
	)

	if wsConn, err = upgrade.Upgrade(w, r, nil); err != nil {
		log.Fatalf("error %v", err)
	}

	if conn, err = impl.InitConnect(wsConn); err != nil {
		log.Fatalf("error %v", err)
	}

	go func(conn *impl.Connection) {
		var err error

		for {
			if err = conn.WriteMessage([]byte("heart beat")); err != nil {
				log.Fatalf("error %v", err)
				conn.Close()
			}
			time.Sleep(5 * time.Second)
		}
	}(conn)

	for {
		if data, err = conn.ReadMessage(); err != nil {
			log.Fatalf("error %v", err)
			conn.Close()
		}

		log.Printf("message: %s\n", data)

		if err = conn.WriteMessage(data); err != nil {
			log.Fatalf("error %v", err)
			conn.Close()
		}
	}
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe(":8888", nil)
}
