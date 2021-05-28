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
	Time     string `json:"time"`
}

type ControlMessage struct {
	Message  string
	Username string
	RoomId   int
	Time     string
}

type UserChannel struct {
	Conn   *websocket.Conn
	RoomId int
}
