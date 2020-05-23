package handler

import (
	"container/list"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
	"webSocketDemo/Demo1/model"
)

type ImHandler struct {
	sync.RWMutex
	userLinks map[string]*list.List
	linkNum   int
}

func NewImHandler(userLinks map[string]*list.List, numLink int) *ImHandler {
	return &ImHandler{userLinks: userLinks, linkNum: numLink}
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
			link.Close()
			return
		}
		mes := &model.Message{}
		//new(model.Message)

		err = json.Unmarshal(message, mes)
		if err != nil {
			log.Printf("Error %s", err)
			continue
		}
		log.Printf("message: %s", message)
		log.Printf("message: %s", mes)

		h.Pong(mes)

		switch mes.Type {
		case model.LOGIN:
			h.Login(link, mes)
			break
		case model.LOGOUT:
			h.LogOut(link, mes)
			break
		case model.SAY:
			h.Say(link, mes)
			break

		}

		link.WriteMessage(websocket.TextMessage, message)
	}
}

func (h *ImHandler) Login(link *websocket.Conn, msg *model.Message) {
	h.Lock()
	defer h.Unlock()
	if nil == h.userLinks[msg.FromUid] {
		h.userLinks[msg.FromUid] = list.New()
	}

	if h.userLinks[msg.FromUid].Len() < h.linkNum {
		h.userLinks[msg.FromUid].PushBack(link)
	}

}

func (h *ImHandler) Say(link *websocket.Conn, msg *model.Message) {
	h.Lock()
	defer h.Unlock()
	msgByte, _ := json.Marshal(msg)
	fmt.Printf("msg:%s\n", msgByte)
	if nil != h.userLinks[msg.ToUid] {
		fmt.Printf("www:%s\n", h.userLinks[msg.ToUid])
		for e := h.userLinks[msg.ToUid].Front(); nil != e; e = e.Next() {
			conn := e.Value.(*websocket.Conn)
			if err := conn.WriteMessage(websocket.TextMessage, msgByte); err != nil {
				log.Println(err)
				h.userLinks[msg.ToUid].Remove(e)
			}
		}
	}

}

func (h *ImHandler) LogOut(link *websocket.Conn, msg *model.Message) {

}

func (h *ImHandler) Pong(msg *model.Message) {
	m := &model.Message{
		ToUid: msg.FromUid,
		Type:  model.PONG,
	}

	msgByte, _ := json.Marshal(m)
	if nil != h.userLinks[msg.FromUid] {
		for e := h.userLinks[msg.FromUid].Front(); nil != e; e = e.Next() {
			conn := e.Value.(*websocket.Conn)
			if err := conn.WriteMessage(websocket.TextMessage, msgByte); err != nil {
				h.userLinks[msg.FromUid].Remove(e)
			}
		}
	}
}
