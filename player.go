package sdtd_stats

type Player struct {
	SteamId      string   `json:"steamid"`
	EntityId     int      `json:"entityid"`
	Ip           string   `json:"ip"`
	Name         string   `json:"name"`
	Online       bool     `json:"online"`
	Position     Position `json:"position"`
	Experience   int      `json:"experience"`
	Level        int      `json:"level"`
	Health       int      `json:"health"`
	Stamina      int      `json:"stamin"`
	ZombieKills  int      `json:"zombiekills"`
	PlayerKills  int      `json:"playerkills"`
	PlayerDeaths int      `json:"playerdeaths"`
	Score        int      `json:"score"`
}

type PlayerInventory struct {
	PlayerName string                `json:"playername"`
	Bag        []PlayerInventoryItem `json:"bag"`
	Belt       []PlayerInventoryItem `json:"belt"`
	Equipment  PlayerEquipment       `json:"equipment"`
}

type PlayerInventoryItem struct {
	Count     int    `json:"count"`
	Name      string `json:"name"`
	Icon      string `json:"icon"`
	IconColor string `json:"iconcolor"`
	Quality   int    `json:"quality"`
}

type PlayerEquipment struct {
	Head     PlayerInventoryItem `json:"head"`
	Eyes     PlayerInventoryItem `json:"eyes"`
	Face     PlayerInventoryItem `json:"face"`
	Armor    PlayerInventoryItem `json:"armor"`
	Jacket   PlayerInventoryItem `json:"jacket"`
	Shirt    PlayerInventoryItem `json:"shirt"`
	LegArmor PlayerInventoryItem `json:"legarmor"`
	Pants    PlayerInventoryItem `json:"pants"`
	Boots    PlayerInventoryItem `json:"boots"`
	Gloves   PlayerInventoryItem `json:"gloves"`
}
