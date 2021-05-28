package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
)

var clients = make(map[string]*UserChannel)
var broadcaster = make(chan ControlMessage)
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleConnections(w http.ResponseWriter, r *http.Request) {
	username, roomId, err := ParseInfo(r)
	if err != nil {
		log.Println(err)
		return
	}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal(err)
	}

	// ensure connection close when function returns
	defer closeClient(username, ws)
	if clients[username] != nil {
		//username = username + "X"
		closeClient(username, clients[username].Conn)
	}
	clients[username] = &UserChannel{ws, roomId}
	log.Println("[INFO] " + username + " connected")

	for {
		var msg ChatMessage
		// Read in a new message as JSON and map it to a Message object
		err := ws.ReadJSON(&msg)
		if err != nil {
			fmt.Println(err)
			delete(clients, username)
			break
		}
		// send new message to the channel
		broadcaster <- ControlMessage{msg.Message, username, roomId}
	}
}

func closeClient(username string, client *websocket.Conn) {
	_ = client.Close()
	fmt.Println("[STATUS] Closed connect of client " + username)
}

// If a message is sent while a client is closing, ignore the error
func unsafeError(err error) bool {
	return !websocket.IsCloseError(err, websocket.CloseGoingAway) && err != io.EOF
}

func handleMessages() {
	for {
		// grab any next message from channel
		msg := <-broadcaster
		messageClients(msg)
	}
}

func messageClients(msg ControlMessage) {
	// send to every client in room
	for username, client := range clients {
		if client.RoomId == msg.RoomId {
			fmt.Println("Send client " + username)
			messageClient(username, client, msg)
		}
	}
}

func messageClient(username string, channel *UserChannel, msg ControlMessage) {
	err := channel.Conn.WriteJSON(SystemMessage{msg.Message, msg.Username})
	if err != nil && unsafeError(err) {
		log.Printf("[ERROR]: %v", err)
		_ = channel.Conn.Close()
		delete(clients, username)
	}
}
