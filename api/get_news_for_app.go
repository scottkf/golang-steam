package api

import (
	"encoding/json"
	"strconv"
)

type GameNewsResult struct {
	AppNews GameNews `json:"appnews"`
}

type GameNews struct {
	AppId     int        `json:"appid"`
	NewsItems []NewsItem `json:"newsitems"`
}

type NewsItem struct {
	Gid           string `json:"gid"`
	Title         string `json:"title"`
	Url           string `json:"url"`
	IsExternalUrl string `json:"is_external_url"`
	Author        string `json:"author"`
	Contents      string `json:"contents"`
	Date          int    `json:"date"`
}

func GetNewsForApp(appId int, count int, length int) (GameNews, int) {
	url := "http://api.steampowered.com/ISteamNews/GetNewsForApp/v0002/?appid=" + strconv.Itoa(appId) + "&count=" + strconv.Itoa(count) + "&maxlength=" + strconv.Itoa(length) + "&format=json"
	body, status := ApiRequest(url)

	var data GameNewsResult
	json.Unmarshal(body, &data)
	return data.AppNews, status
}
