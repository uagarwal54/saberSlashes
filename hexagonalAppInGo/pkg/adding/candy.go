package adding

type Candy struct {
	Id      int
	Name    string `json:"name"`
	Company string `json:"company"`
	Ppu     int    `json:"ppu"`
	Stock   int    `json:"stock"`
}
