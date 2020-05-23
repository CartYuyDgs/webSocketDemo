package impl

import (
	"errors"
	"github.com/gorilla/websocket"
	"log"
	"sync"
)

type Connection struct {
	wsConnect *websocket.Conn
	inCHan    chan []byte
	outChan   chan []byte
	closeChan chan byte

	mutex   sync.Mutex
	isClose bool
}

func InitConnect(wsConn *websocket.Conn) (conn *Connection, err error) {
	conn = &Connection{
		wsConnect: wsConn,
		inCHan:    make(chan []byte, 100),
		outChan:   make(chan []byte, 100),
		closeChan: make(chan byte, 1),
	}
	// 启动读协程
	go conn.readLoop()
	// 启动写协程
	go conn.writeLoop()
	return
}

func (conn *Connection) Close() {
	// 线程安全，可多次调用
	conn.wsConnect.Close()
	// 利用标记，让closeChan只关闭一次
	conn.mutex.Lock()
	if !conn.isClose {
		close(conn.closeChan)
		conn.isClose = true
	}
	conn.mutex.Unlock()
}

func (conn *Connection) readLoop() {
	var (
		data []byte
		err  error
	)

	for {
		if _, data, err = conn.wsConnect.ReadMessage(); err != nil {
			goto ERR
		}

		select {
		case conn.inCHan <- data:
			log.Printf("inChan data %s\n", data)
		case <-conn.closeChan:
			goto ERR

		}
	}
ERR:
	conn.Close()
}

func (conn *Connection) ReadMessage() (data []byte, err error) {
	select {
	case data = <-conn.inCHan:
		log.Printf("read data:%s\n", data)
	case <-conn.closeChan:
		err = errors.New("connect is close")
	}
	return
}

func (conn *Connection) WriteMessage(date []byte) (err error) {
	select {
	case conn.outChan <- date:
		log.Printf("write1 data:%s\n", date)
	case <-conn.closeChan:
		err = errors.New("connect is close")
	}
	return
}

func (conn *Connection) writeLoop() {
	var (
		data []byte
		err  error
	)

	for {
		select {
		case data = <-conn.outChan:
		case <-conn.closeChan:
			goto ERR
		}
		if err = conn.wsConnect.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}

ERR:
	conn.Close()

}
