package repository

import (
	"database/sql"

	"github.com/JoaquimPassareli/go-api/models"
)

type PessoaRepository struct {
	db *sql.DB
}

type CarroRepository struct {
	db *sql.DB
}

// Repository de Pessoas! //
func NewPessoaRepository(db *sql.DB) *PessoaRepository {
	return &PessoaRepository{db: db}
}

func (r *PessoaRepository) Create(p *models.Pessoa) error {
	stmt, err := r.db.Prepare("INSERT INTO pessoas (nome, idade, altura, doc) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	result, err := stmt.Exec(p.Nome, p.Idade, p.Altura, p.Doc)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	p.ID = int(id)
	return nil
}

func (r *PessoaRepository) Read() ([]models.Pessoa, error) {
	stmt, err := r.db.Prepare("SELECT nome, idade, altura, doc FROM pessoas")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var pessoas []models.Pessoa
	for rows.Next() {
		var p models.Pessoa
		err := rows.Scan(&p.Nome, &p.Idade, &p.Altura, &p.Doc)
		if err != nil {
			return nil, err
		}
		pessoas = append(pessoas, p)
	}
	return pessoas, nil
}

func (r *PessoaRepository) ReadById(doc int) (models.Pessoa, error) {
	stmt, err := r.db.Prepare("SELECT nome, idade, altura, doc FROM pessoas WHERE doc = ?")
	if err != nil {
		return models.Pessoa{}, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(doc)

	var pessoa models.Pessoa
	err = row.Scan(&pessoa.Nome, &pessoa.Idade, &pessoa.Altura, &pessoa.Doc)
	if err != nil {
		return models.Pessoa{}, err
	}

	return pessoa, nil
}

func (r *PessoaRepository) Update(p *models.Pessoa) error {
	stmt, err := r.db.Prepare("UPDATE pessoas SET nome = ?, idade = ?, altura = ? WHERE doc = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(p.Nome, p.Idade, p.Altura, p.Doc)
	return err
}

func (r *PessoaRepository) Delete(doc int) error {
	stmt, err := r.db.Prepare("DELETE FROM pessoas WHERE doc = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(doc)
	return err
}

//Repository de Carros !//

func NewCarroRepository(db *sql.DB) *CarroRepository {
	return &CarroRepository{db: db}
}

func (r *CarroRepository) Create(p *models.Carro) error {
	stmt, err := r.db.Prepare("INSERT INTO carros (marca, modelo, ano, cor) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(p.Marca, p.Modelo, p.Ano, p.Cor)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	p.ID = int(id)
	return nil
}

func (r *CarroRepository) Read() ([]models.Carro, error) {
	rows, err := r.db.Query("SELECT id, marca, modelo, ano, cor FROM carros")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var carros []models.Carro
	for rows.Next() {
		var c models.Carro
		if err := rows.Scan(&c.ID, &c.Marca, &c.Modelo, &c.Ano, &c.Cor); err != nil {
			return nil, err
		}
		carros = append(carros, c)
	}
	return carros, nil
}

func (r *CarroRepository) ReadById(id int) (models.Carro, error) {
	row := r.db.QueryRow("SELECT id, marca, modelo, ano, cor FROM carros WHERE id = ?", id)

	var c models.Carro
	err := row.Scan(&c.ID, &c.Marca, &c.Modelo, &c.Ano, &c.Cor)
	return c, err
}

func (r *CarroRepository) Update(p *models.Carro) error {
	stmt, err := r.db.Prepare("UPDATE carros SET marca=?, modelo=?, ano=?, cor=? WHERE id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(p.Marca, p.Modelo, p.Ano, p.Cor, p.ID)
	return err
}

func (r *CarroRepository) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM carros WHERE id=?", id)
	return err
}
