package sdtd_stats

type Entity struct {
	Id       string   `json:"id"`
	Name     string   `json:"name"`
	Position Position `json:"position"`
}

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
}
