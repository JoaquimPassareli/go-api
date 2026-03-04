package repository

import (
	"database/sql"

	"github.com/JoaquimPassareli/go-api/models"
)

type EnderecoRepository struct {
	db *sql.DB
}

func NewEnderecoRepository(db *sql.DB) *EnderecoRepository {
	return &EnderecoRepository{db: db}
}

func (r *EnderecoRepository) Create(e *models.Endereco) error {
	res, err := r.db.Exec(
		"INSERT INTO enderecos (cep, bairro, rua, numero, cidade, estado, pessoa_id) VALUES (?, ?, ?, ?, ?, ?, ?)",
		e.Cep, e.Bairro, e.Rua, e.Numero, e.Cidade, e.Estado, e.PessoaID)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}
	e.ID = int(id)
	return nil
}

func (r *EnderecoRepository) Read() ([]models.Endereco, error) {
	rows, err := r.db.Query(
		"SELECT id, cep, bairro, rua, numero, cidade, estado, pessoa_id FROM enderecos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lista []models.Endereco
	for rows.Next() {
		var e models.Endereco
		err := rows.Scan(&e.ID, &e.Cep, &e.Bairro, &e.Rua, &e.Numero, &e.Cidade, &e.Estado, &e.PessoaID)
		if err != nil {
			return nil, err
		}
		lista = append(lista, e)
	}
	return lista, nil
}

func (r *EnderecoRepository) ReadByPessoaId(pessoaId int) ([]models.Endereco, error) {
	rows, err := r.db.Query(
		"SELECT id, cep, bairro, rua, numero, cidade, estado, pessoa_id FROM enderecos WHERE pessoa_id = ?", pessoaId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lista []models.Endereco
	for rows.Next() {
		var e models.Endereco
		err := rows.Scan(&e.ID, &e.Cep, &e.Bairro, &e.Rua, &e.Numero, &e.Cidade, &e.Estado, &e.PessoaID)
		if err != nil {
			return nil, err
		}
		lista = append(lista, e)
	}
	return lista, nil
}

func (r *EnderecoRepository) Update(id int, e *models.Endereco) error {
	_, err := r.db.Exec(
		"UPDATE enderecos SET cep = ?, bairro = ?, rua = ?, numero = ?, cidade = ?, estado = ?, pessoa_id = ? WHERE id = ?",
		e.Cep, e.Bairro, e.Rua, e.Numero, e.Cidade, e.Estado, e.PessoaID, id)
	return err
}

func (r *EnderecoRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM enderecos WHERE id = ?", id)
	return err
}
