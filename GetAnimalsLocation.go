package sdtd_stats

import "encoding/json"

func (s *SdtdStats) GetAnimalsLocation() []Entity {
	var entities []Entity
	resp, _ := s.makeCall("getanimalslocation")
	json.Unmarshal(resp, &entities)
	return entities
}
