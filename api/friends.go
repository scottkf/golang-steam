package api

import "encoding/json"

type FriendsResult struct {
	FriendsList FriendsList `json:"friendslist"`
}

type FriendsList struct {
	Friends []Friend `json:"friends"`
}

type Friend struct {
	SteamId      string `json:"steamid"`
	Relationship string `json:"relationship"`
	FriendSince  int    `json:"friend_since"`
}

func GetFriends(key string, id SteamId) (FriendsList, int) {
	url := "http://api.steampowered.com/ISteamUser/GetFriendList/v1?key=" + key + "&steamid=" + id.ToString() + "&format=json"

	body, status := ApiRequest(url)

	var data FriendsResult
	json.Unmarshal(body, &data)
	return data.FriendsList, status
}
