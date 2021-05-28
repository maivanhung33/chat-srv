package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

var ServiceAccountUri string

func validateToken(token string) (*UserInfo, error) {
	req, _ := http.NewRequest("GET", ServiceAccountUri+"/me/userinfo", nil)
	req.Header.Set("Authorization", "Bearer "+token)

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var userInfo UserInfo

	defer resp.Body.Close()

	if resp == nil || resp.StatusCode != 200 {
		return nil, errors.New("UNAUTHORIZED")
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("A", err)
		return nil, err
	}
	err = json.Unmarshal(body, &userInfo)
	if err != nil {
		fmt.Println("B", err)
		return nil, err
	}
	return &userInfo, nil
}

func ParseInfo(r *http.Request) (string, int, error) {
	token := r.URL.Query().Get("token")
	roomIdString := r.URL.Query().Get("roomId")

	username := ""
	if token == "" || roomIdString == "" {
		return "", 0, errors.New("TOKEN AND ROOM_ID NEED FILL")
	}

	user, err := validateToken(token)
	if err != nil {
		return "", 0, err
	}

	username = user.Username
	roomId, _ := strconv.Atoi(roomIdString)

	if roomId < LOTTERY || roomId > SIGBO {
		return "", 0, errors.New("ROOM_ID INVALID")
	}

	return username, roomId, nil
}
