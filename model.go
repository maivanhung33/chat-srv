package main

import "github.com/gorilla/websocket"

const (
	LOTTERY = iota
	XOCDIA
	SCRATCH
	SIGBO
)

type UserInfo struct {
	Username string
}

type ChatMessage struct {
	Message string `json:"message"`
}

type SystemMessage struct {
	Message  string `json:"message"`
	Username string `json:"username"`
}

type ControlMessage struct {
	Message  string
	Username string `json:"username"`
	RoomId   int
}

type UserChannel struct {
	Conn   *websocket.Conn
	RoomId int
}
