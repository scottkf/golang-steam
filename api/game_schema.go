package api

import (
	"encoding/json"
	"strconv"
)

type GameSchemaResult struct {
	Game GameSchema `json:"game"`
}

type GameSchema struct {
	GameName       string          `json:"gameName"`
	GameVersion    string          `json:"gameVersion"`
	AvailableStats GameSchemaStats `json:"availableGameStats"`
}

type GameSchemaStats struct {
	Achievements []Achievement `json:"achievements"`
}

type Achievement struct {
	ApiName      string `json:"apiname"`
	Achieved     int    `json:"achieved"`
	Name         string `json:"name"`
	DefaultValue int    `json:"defaultvalue"`
	DisplayName  string `json:"displayName"`
	Hidden       int    `json:"hidden"`
	Description  string `json:"description"`
	IconClosed   string `json:"icon"`
	IconOpen     string `json:"icongray"`
}

func GetGameSchema(key string, appId int) (GameSchema, int) {
	url := "http://api.steampowered.com/ISteamUserStats/GetSchemaForGame/v2?key=" + key + "&appid=" + strconv.Itoa(appId) + "&format=json"

	body, status := ApiRequest(url)

	var data GameSchemaResult
	json.Unmarshal(body, &data)
	return data.Game, status
}
