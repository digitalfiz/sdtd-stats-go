package sdtd_stats

import "encoding/json"

func (s *SdtdStats) GetLandClaims() Claims {
	var claims Claims
	resp, _ := s.makeCall("getlandclaims")
	json.Unmarshal(resp, &claims)
	return claims
}
