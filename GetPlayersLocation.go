package sdtd_stats

import "encoding/json"

func (s *SdtdStats) GetPlayersLocation() []Player {
	var players []Player
	resp, _ := s.makeCall("getplayerslocation")

	json.Unmarshal(resp, &players)

	return players
}
