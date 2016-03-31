package sdtd_stats

type Claims struct {
	ClaimSize   int          `json:"claimsize"`
	ClaimOwners []ClaimOwner `json:"claimowners"`
}

type ClaimOwner struct {
	SteamId     string     `json:"steamid"`
	ClaimActive bool       `json:"claimactive"`
	PlayerName  string     `json:"playername"`
	Claims      []Position `json:"claims"`
}
