package models

type Genre struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (g *Genre) TableName() string {
	return "genres"
}
