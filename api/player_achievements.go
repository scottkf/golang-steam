package api

import (
	"encoding/json"
	"strconv"
)

type AchievementResult struct {
	Response AchievementResponse `json:"playerstats"`
}

type AchievementResponse struct {
	SteamId      string `json:"steamID"`
	GameName     string `json:"gameName"`
	Achievements []Achievement
}

func GetPlayerAchievements(key string, id SteamId, appId int) (AchievementResponse, int) {
	url := "http://api.steampowered.com/ISteamUserStats/GetPlayerAchievements/v1?key=" + key + "&appid=" + strconv.Itoa(appId) + "&steamid=" + id.ToString() + "&format=json&include_appinfo=1"

	body, status := ApiRequest(url)

	var data AchievementResult
	json.Unmarshal(body, &data)
	return data.Response, status
}
