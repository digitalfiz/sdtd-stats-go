package sdtd_stats

import "encoding/json"

func (s *SdtdStats) GetWebUIUpdates() Stats {
	var update Stats
	resp, _ := s.makeCall("getwebuiupdates")
	json.Unmarshal(resp, &update)
	return update
}
