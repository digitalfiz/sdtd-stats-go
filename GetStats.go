package sdtd_stats

import "encoding/json"

func (s *SdtdStats) GetStats() Stats {
	var update Stats
	resp, _ := s.makeCall("getstats")
	json.Unmarshal(resp, &update)
	return update
}
