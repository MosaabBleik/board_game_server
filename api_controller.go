package main

import (
	"io"
	"net/http"
	"net/url"
)

func Login(username, password string) {

}

func GetPlayers(roomName string) ([]byte, error) {
	body := url.Values{
		"room": {roomName},
	}
	url := "https://game.wowdigital.sa/api/method/wow_game.wow_game.gameApi.GetPlayers"

	resp, err := http.PostForm(url, body)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func GetStepCards(roomName string, playerNo string) ([]byte, error) {
	body := url.Values{
		"room":      {roomName},
		"player_no": {playerNo},
	}
	url := "https://game.wowdigital.sa/api/method/wow_game.wow_game.gameApi.GetCards"

	resp, err := http.PostForm(url, body)
	if err != nil {
		return nil, err
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return data, nil
}
