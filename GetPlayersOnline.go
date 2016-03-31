package sdtd_stats

import "encoding/json"

func (s *SdtdStats) GetPlayersOnline() []Player {
	var players []Player
	resp, _ := s.makeCall("getplayersonline")
	json.Unmarshal(resp, &players)
	return players
}
