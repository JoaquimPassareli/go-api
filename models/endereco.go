package models

type Endereco struct {
	ID       int
	Cep      string
	Bairro   string
	Rua      string
	Numero   int
	Cidade   string
	Estado   string
	PessoaID *int
}

type EnderecoRequest struct {
	Cep       string `json:"cep"`
	Bairro    string `json:"bairro"`
	Rua       string `json:"rua"`
	Numero    int    `json:"numero"`
	Cidade    string `json:"cidade"`
	Estado    string `json:"estado"`
	PessoaDoc int    `json:"pessoaDoc"`
}

type EnderecoResponse struct {
	ID        int    `json:"id"`
	Cep       string `json:"cep"`
	Bairro    string `json:"bairro"`
	Rua       string `json:"rua"`
	Numero    int    `json:"numero"`
	Cidade    string `json:"cidade"`
	Estado    string `json:"estado"`
	PessoaDoc *int   `json:"pessoaDoc"`
}
