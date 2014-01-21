package api

import (
	"encoding/json"
)

type PlayerSummaryResult struct {
	Response PlayerList `json:"response"`
}

type PlayerList struct {
	Players []Player `json:"players"`
}

type Player struct {
	SteamId                  string `json:"steamid"`
	CommunityVisibilityState int    `json:"communityvisibilitystate"`
	ProfileState             int    `json:"profilestate"`
	PersonaState             int    `json:"personastate"`
	PersonaName              string `json:"personaname"`
	LastLogoff               int    `json:"lastlogoff"`
	ProfileUrl               string `json:"profileurl"`
	Avatar                   string `json:"avatar"`
	AvatarMedium             string `json:"avatarmedium"`
	AvatarFull               string `json:"avatarfull"`
	PrimaryClanId            string `json:"primaryclanid"`
	TimeCreated              int    `json:"timecreated"`
	LocCountryCode           string `json:"loccountrycode"`
}

func GetPlayerSummary(key string, id SteamId) (PlayerList, int) {
	url := "http://api.steampowered.com/ISteamUser/GetPlayerSummaries/v0002/?key=" + key + "&steamids=" + id.ToString() + "&format=json"

	body, status := ApiRequest(url)

	var data PlayerSummaryResult
	json.Unmarshal(body, &data)
	return data.Response, status
}
