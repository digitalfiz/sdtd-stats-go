package sdtd_stats

import "encoding/json"

func (s *SdtdStats) GetHostileLocation() []Entity {
	var entities []Entity
	resp, _ := s.makeCall("gethostilelocation")
	json.Unmarshal(resp, &entities)
	return entities
}
