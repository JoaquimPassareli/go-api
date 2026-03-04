package models

type Pessoa struct {
	ID     int
	Nome   string
	Idade  int
	Altura float64
	Doc    int
}

type PessoaRequest struct {
	Nome   string  `json:"nome"`
	Idade  int     `json:"idade"`
	Altura float64 `json:"altura"`
	Doc    int     `json:"doc"`
}

type PessoaResponse struct {
	ID        int                `json:"id"`
	Nome      string             `json:"nome"`
	Idade     int                `json:"idade"`
	Altura    float64            `json:"altura"`
	Doc       int                `json:"doc"`
	Carros    []CarroResponse    `json:"carros"`
	Enderecos []EnderecoResponse `json:"enderecos"`
}

type Carro struct {
	ID       int
	Marca    string
	Modelo   string
	Ano      int
	Cor      string
	PessoaID *int
}

type CarroRequest struct {
	Marca     string `json:"marca"`
	Modelo    string `json:"modelo"`
	Ano       int    `json:"ano"`
	Cor       string `json:"cor"`
	PessoaDoc *int   `json:"pessoaDoc"`
	PessoaID  *int   `json:"pessoaId"`
}

type CarroResponse struct {
	ID        int    `json:"id"`
	Marca     string `json:"marca"`
	Modelo    string `json:"modelo"`
	Ano       int    `json:"ano"`
	Cor       string `json:"cor"`
	Pessoadoc *int   `json:"pessoaDoc"`
	PessoaID  *int   `json:"pessoaId"`
}
