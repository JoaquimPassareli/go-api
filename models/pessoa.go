package models

type Pessoa struct {
	ID     int     `json:"id"`
	Nome   string  `json:"nome"`
	Idade  int     `json:"idade"`
	Altura float64 `json:"altura"`
	Doc    int     `json:"doc"`
}

type Carro struct {
	ID     int    `json:"id"`
	Marca  string `json:"marca"`
	Modelo string `json:"modelo"`
	Ano    int    `json:"ano"`
	Cor    string `json:"cor"`
}
