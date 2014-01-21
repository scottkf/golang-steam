package api

import "encoding/json"

type GameResult struct {
	Response GameResponse `json:"response"`
}

type GameResponse struct {
	GameCount *int32 `json:"game_count,omitempty"`
	Games     []Game `json:"games"`
}

type Game struct {
	AppId                    int    `json:"appid"`
	Name                     string `json:"name"`
	ImgIconUrl               string `json:"img_icon_url"`
	ImgLogoUrl               string `json:"img_logo_url"`
	HasCommunityVisibleStats bool   `json:"has_community_visible_stats"`
	PlaytimeForever          int    `json:"playtime_forever"`
	Playtime2Weeks           int    `json:"playtime_2weeks"`
}

func GetGames(key string, id SteamId) (GameResponse, int) {
	url := "http://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key=" + key + "&steamid=" + id.ToString() + "&format=json&include_played_free_games=1&include_appinfo=1"

	body, status := ApiRequest(url)

	var data GameResult
	json.Unmarshal(body, &data)
	return data.Response, status
}
