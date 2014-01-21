package api

import "strconv"

type SteamId struct {
	Id64        int
	CommunityId string
}

func (id *SteamId) ValidSteamId() bool {
	part_two := id.Id64 - 76561197960265728

	if part_two > 0 {
		return true
	} else {
		return false
	}
}

func ConvertId64ToSteamId(id string) SteamId {
	steam_id, err := strconv.Atoi(id)
	if err != nil {
		steam_id = 0
	}
	return SteamId{Id64: steam_id}
}

func (id *SteamId) ToString() string {
	return strconv.Itoa(id.Id64)
}

func (id *SteamId) ConvertSteamIdToCommunityId() {
	part_one := id.Id64 % 2
	part_two := id.Id64 - 76561197960265728

	if part_two > 0 {
		part_two = (part_two - part_one) / 2
		id.CommunityId = "STEAM_0:" + strconv.Itoa(part_one) + ":" + strconv.Itoa(part_two)
	} else {
		id.CommunityId = ""
	}
}
