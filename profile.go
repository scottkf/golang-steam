package steam

import "github.com/scottkf/steam/api"

type Profile struct {
	Id      SteamId
	Friends []api.Friend
	Games   []api.Game
	Info    api.Player
}

func (p *Profile) GetGames(key string, id int) {
	data := api.GetGames(key, id)
	p.Games = data.Games
}

func (p *Profile) GetFriends(key string, id int) {
	data := api.GetFriends(key, id)
	p.Friends = data.Friends
}

func (p *Profile) GetPlayerSummary(key string, id int) {
	data := api.GetPlayerSummary(key, id)
	if len(data.Players) > 0 {
		p.Info = data.Players[0]
	}
}
