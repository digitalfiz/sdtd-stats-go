package sdtd_stats

type Stats struct {
	Gametime GameTime `json:"gametime"`
	Players  int      `json:"players"`
	Hostiles int      `json:"hostiles"`
	Animals  int      `json:"animals"`
	Newlogs  int      `json:"newlogs"`
}

type GameTime struct {
	Days    int `json:"days"`
	Hours   int `json:"hours"`
	Minutes int `json:"minutes"`
}
