package repository

import (
	"database/sql"
	"fmt"

	"github.com/JoaquimPassareli/go-api/models"
)

type PessoaRepository struct {
	db *sql.DB
}

type CarroRepository struct {
	db *sql.DB
}

func NewPessoaRepository(db *sql.DB) *PessoaRepository {
	return &PessoaRepository{db: db}
}

func NewCarroRepository(db *sql.DB) *CarroRepository {
	return &CarroRepository{db: db}
}

//
// ========================
// PESSOA REPOSITORY
// ========================
//

func (r *PessoaRepository) Create(p *models.Pessoa) error {
	res, err := r.db.Exec(
		"INSERT INTO pessoas (nome, idade, altura, doc) VALUES (?, ?, ?, ?)",
		p.Nome, p.Idade, p.Altura, p.Doc,
	)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	p.ID = int(id)
	return nil
}

func (r *PessoaRepository) Read() ([]models.Pessoa, error) {
	rows, err := r.db.Query(
		"SELECT id, nome, idade, altura, doc FROM pessoas",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lista []models.Pessoa

	for rows.Next() {
		var p models.Pessoa

		if err := rows.Scan(
			&p.ID,
			&p.Nome,
			&p.Idade,
			&p.Altura,
			&p.Doc,
		); err != nil {
			return nil, err
		}

		lista = append(lista, p)
	}

	return lista, nil
}

func (r *PessoaRepository) ReadByID(id int) (models.Pessoa, error) {
	var p models.Pessoa

	err := r.db.QueryRow(
		"SELECT id, nome, idade, altura, doc FROM pessoas WHERE id = ?",
		id,
	).Scan(
		&p.ID,
		&p.Nome,
		&p.Idade,
		&p.Altura,
		&p.Doc,
	)

	return p, err
}

func (r *PessoaRepository) ReadByDoc(doc int) (models.Pessoa, error) {
	var p models.Pessoa

	err := r.db.QueryRow(
		"SELECT id, nome, idade, altura, doc FROM pessoas WHERE doc = ?",
		doc,
	).Scan(
		&p.ID,
		&p.Nome,
		&p.Idade,
		&p.Altura,
		&p.Doc,
	)

	return p, err
}

func (r *PessoaRepository) Update(p *models.Pessoa) error {
	_, err := r.db.Exec(
		"UPDATE pessoas SET nome = ?, idade = ?, altura = ?, doc = ? WHERE id = ?",
		p.Nome,
		p.Idade,
		p.Altura,
		p.Doc,
		p.ID,
	)

	return err
}

func (r *PessoaRepository) Delete(id int) error {
	res, err := r.db.Exec(
		"DELETE FROM pessoas WHERE id = ?",
		id,
	)
	if err != nil {
		return err
	}

	count, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if count == 0 {
		return fmt.Errorf("not found")
	}

	return nil
}

//
// ========================
// CARRO REPOSITORY
// ========================
//

func (r *CarroRepository) Create(c *models.Carro) error {
	res, err := r.db.Exec(
		"INSERT INTO carros (marca, modelo, ano, cor, pessoa_id) VALUES (?, ?, ?, ?, ?)",
		c.Marca,
		c.Modelo,
		c.Ano,
		c.Cor,
		c.PessoaID,
	)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	c.ID = int(id)
	return nil
}

func (r *CarroRepository) Read() ([]models.Carro, error) {
	rows, err := r.db.Query(
		"SELECT id, marca, modelo, ano, cor, pessoa_id FROM carros",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lista []models.Carro

	for rows.Next() {
		var c models.Carro

		if err := rows.Scan(
			&c.ID,
			&c.Marca,
			&c.Modelo,
			&c.Ano,
			&c.Cor,
			&c.PessoaID,
		); err != nil {
			return nil, err
		}

		lista = append(lista, c)
	}

	return lista, nil
}

func (r *CarroRepository) ReadByPessoaId(id int) ([]models.Carro, error) {
	rows, err := r.db.Query(
		"SELECT id, marca, modelo, ano, cor, pessoa_id FROM carros WHERE pessoa_id = ?",
		id,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var lista []models.Carro

	for rows.Next() {
		var c models.Carro

		if err := rows.Scan(&c.ID, &c.Marca, &c.Modelo, &c.Ano, &c.Cor, &c.PessoaID); err != nil {
			return nil, err
		}

		lista = append(lista, c)
	}

	return lista, nil
}

func (r *CarroRepository) Update(id int, c *models.Carro) error {
	_, err := r.db.Exec(
		"UPDATE carros SET marca = ?, modelo = ?, ano = ?, cor = ?, pessoa_id = ? WHERE id = ?",
		c.Marca,
		c.Modelo,
		c.Ano,
		c.Cor,
		c.PessoaID,
		id,
	)

	return err
}

func (r *CarroRepository) Delete(id int) error {
	_, err := r.db.Exec(
		"DELETE FROM carros WHERE id = ?",
		id,
	)

	return err
}

func (r *CarroRepository) DeleteByPessoaId(id int) error {
	_, err := r.db.Exec(
		"DELETE FROM carros WHERE pessoa_id = ?",
		id,
	)

	return err
}
