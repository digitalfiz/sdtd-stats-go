package sdtd_stats

import (
	"encoding/json"
	"fmt"
)

func (s *SdtdStats) GetPlayerInventory(steamid string) PlayerInventory {
	var inventory PlayerInventory
	resp, _ := s.makeCallWithParams("getplayerinventory", fmt.Sprintf("steamid=%v", steamid))
	json.Unmarshal(resp, &inventory)
	return inventory
}
