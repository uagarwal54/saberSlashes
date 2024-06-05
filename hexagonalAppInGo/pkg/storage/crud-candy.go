package storage

import (
	"hexagonalAppInGo/pkg/adding"
)

func (s *Storage) GetAllCandies() (candies []string, err error) {
	results, err := s.db.Query("SELECT * from candies")
	if err != nil {
		return
	}
	defer results.Close()
	var candy adding.Candy

	for results.Next() {
		results.Scan(&candy.Id, &candy.Name, &candy.Company, &candy.Ppu, &candy.Stock)
		candies = append(candies, candy.Name)
	}
	return
}

func (s *Storage) AddCandy(c adding.Candy) (id string, err error) {
	res, err := s.db.Exec("insert into candies (name,comapnany,ppu,stock) values (?,?,?,?)", c.Name, c.Company, c.Ppu, c.Stock)
	if err != nil {
		return
	}
	var idInt int64
	idInt, err = res.LastInsertId()
	id = string(idInt)
	return
}
