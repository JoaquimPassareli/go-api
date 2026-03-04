package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/JoaquimPassareli/go-api/database"
	"github.com/JoaquimPassareli/go-api/models"
	"github.com/JoaquimPassareli/go-api/repository"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"

	_ "github.com/mattn/go-sqlite3"
)

var pessoaRepo *repository.PessoaRepository
var carroRepo *repository.CarroRepository
var enderecoRepo *repository.EnderecoRepository

// ===================== PESSOAS =====================

// GET /pessoas
func getAllPessoas(c *echo.Context) error {
	pessoas, err := pessoaRepo.Read()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Erro ao buscar pessoas",
		})
	}

	response := []models.PessoaResponse{}

	for _, p := range pessoas {
		carros, _ := carroRepo.ReadByPessoaId(p.ID)
		enderecos, _ := enderecoRepo.ReadByPessoaId(p.ID) // ← adiciona isso

		carrosResponse := []models.CarroResponse{}
		for _, car := range carros {
			carrosResponse = append(carrosResponse, models.CarroResponse{
				ID:       car.ID,
				Marca:    car.Marca,
				Modelo:   car.Modelo,
				Ano:      car.Ano,
				Cor:      car.Cor,
				PessoaID: car.PessoaID,
			})
		}

		enderecosResponse := []models.EnderecoResponse{}
		for _, e := range enderecos {
			enderecosResponse = append(enderecosResponse, models.EnderecoResponse{
				ID:     e.ID,
				Cep:    e.Cep,
				Bairro: e.Bairro,
				Rua:    e.Rua,
				Numero: e.Numero,
				Cidade: e.Cidade,
				Estado: e.Estado,
			})
		}

		response = append(response, models.PessoaResponse{
			ID:        p.ID,
			Nome:      p.Nome,
			Idade:     p.Idade,
			Altura:    p.Altura,
			Doc:       p.Doc,
			Carros:    carrosResponse,
			Enderecos: enderecosResponse, // ← adiciona isso
		})
	}

	return c.JSON(http.StatusOK, response)
}

// GET /pessoas/:id
func getPessoaByID(c *echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "ID inválido",
		})
	}

	pessoa, err := pessoaRepo.ReadByID(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Pessoa não encontrada",
		})
	}

	return c.JSON(http.StatusOK, pessoa)
}

// GET /pessoas/doc/:doc
func getPessoaByDoc(c *echo.Context) error {
	doc, err := strconv.Atoi(c.Param("doc"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Doc inválido",
		})
	}

	pessoa, err := pessoaRepo.ReadByDoc(doc)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Pessoa não encontrada",
		})
	}

	return c.JSON(http.StatusOK, pessoa)
}

// POST /pessoas
func createPessoa(c *echo.Context) error {
	var p models.Pessoa

	if err := c.Bind(&p); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "JSON inválido",
		})
	}

	if p.Idade < 18 {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Menor de idade",
		})
	}

	if err := pessoaRepo.Create(&p); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Erro ao criar pessoa",
		})
	}

	return c.JSON(http.StatusCreated, p)
}

// PUT /pessoas/:id
func updatePessoa(c *echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "ID inválido",
		})
	}

	var req models.PessoaRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "JSON inválido",
		})
	}

	pessoa := models.Pessoa{
		ID:     id,
		Nome:   req.Nome,
		Idade:  req.Idade,
		Altura: req.Altura,
	}

	if err := pessoaRepo.Update(&pessoa); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Erro ao atualizar",
		})
	}

	return c.JSON(http.StatusOK, pessoa)
}

// DELETE /pessoas/:id
func deletePessoa(c *echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "ID inválido",
		})
	}

	carros, err := carroRepo.ReadByPessoaId(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Erro ao verificar carros",
		})
	}

	if len(carros) > 0 {
		if err := carroRepo.DeleteByPessoaId(id); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": "Erro ao deletar carros",
			})
		}
	}

	if err := pessoaRepo.Delete(id); err != nil {
		fmt.Println("ERRO AO DELETAR PESSOA:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Pessoa deletada com sucesso (e carros, se existiam)",
	})
}

// ===================== CARROS =====================

// GET /carros
func getAllCarros(c *echo.Context) error {
	carros, err := carroRepo.Read()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Erro ao buscar carros",
		})
	}
	var carrosResponse []models.CarroResponse
	for _, c := range carros {
		carrosResponse = append(carrosResponse, models.CarroResponse{
			ID:       c.ID,
			Marca:    c.Marca,
			Modelo:   c.Modelo,
			Ano:      c.Ano,
			Cor:      c.Cor,
			PessoaID: c.PessoaID,
		})
	}
	return c.JSON(http.StatusOK, carrosResponse)
}

// POST /carros
func createCarro(c *echo.Context) error {
	var req models.CarroRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "JSON inválido",
		})
	}

	// converter request → model
	carro := models.Carro{
		Marca:    req.Marca,
		Modelo:   req.Modelo,
		Ano:      req.Ano,
		Cor:      req.Cor,
		PessoaID: req.PessoaID, // se ainda for ponteiro
	}

	if err := carroRepo.Create(&carro); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Erro ao criar carro",
		})
	}

	return c.JSON(http.StatusCreated, carro)
}

func updateCarro(c *echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "ID inválido",
		})
	}
	var req models.CarroRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "JSON inválido",
		})
	}

	pessoa, err := pessoaRepo.ReadByDoc(*req.PessoaDoc)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Pessoa não encontrada com esse DOC",
		})
	}

	carro := models.Carro{
		ID:       id,
		Marca:    req.Marca,
		Modelo:   req.Modelo,
		Ano:      req.Ano,
		Cor:      req.Cor,
		PessoaID: &pessoa.ID,
	}

	if err := carroRepo.Update(id, &carro); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Erro ao atualizar carro",
		})
	}
	return c.JSON(http.StatusOK, carro)
}

// DELETE /carros/:id
func deleteCarro(c *echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "ID inválido",
		})
	}

	if err := carroRepo.Delete(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Erro ao deletar carro",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Carro deletado",
	})
}

// ===================== ENDEREÇOS =====================

func getAllEnderecos(c *echo.Context) error {
	enderecos, err := enderecoRepo.Read()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Erro ao buscar endereços",
		})
	}

	response := []models.EnderecoResponse{}
	for _, e := range enderecos {
		var pessoaDoc int
		if e.PessoaID != nil {
			pessoa, err := pessoaRepo.ReadByID(*e.PessoaID)
			if err == nil {
				pessoaDoc = pessoa.Doc
			}
		}

		response = append(response, models.EnderecoResponse{
			ID:        e.ID,
			Cep:       e.Cep,
			Bairro:    e.Bairro,
			Rua:       e.Rua,
			Numero:    e.Numero,
			Cidade:    e.Cidade,
			Estado:    e.Estado,
			PessoaDoc: &pessoaDoc,
		})
	}
	return c.JSON(http.StatusOK, response)
}

func getEnderecosByPessoaId(c *echo.Context) error {
	pessoaId, err := strconv.Atoi(c.Param("pessoaId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "ID de pessoa inválido",
		})
	}
	enderecos, err := enderecoRepo.ReadByPessoaId(pessoaId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Erro ao buscar endereços",
		})
	}

	return c.JSON(http.StatusOK, enderecos)
}

func createEndereco(c *echo.Context) error {
	var req models.EnderecoRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "JSON inválido",
		})
	}

	fmt.Println("PessoaDoc recebido:", req.PessoaDoc)

	pessoa, err := pessoaRepo.ReadByDoc(req.PessoaDoc)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Pessoa não encontrada com esse documento",
		})
	}

	endereco := models.Endereco{
		Cep:      req.Cep,
		Bairro:   req.Bairro,
		Rua:      req.Rua,
		Numero:   req.Numero,
		Cidade:   req.Cidade,
		Estado:   req.Estado,
		PessoaID: &pessoa.ID,
	}

	if err := enderecoRepo.Create(&endereco); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, endereco)
}

func updateEndereco(c *echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "ID inválido",
		})
	}
	var req models.EnderecoRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "JSON inválido",
		})
	}

	pessoa, err := pessoaRepo.ReadByDoc(req.PessoaDoc)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Pessoa não encontrada com esse documento",
		})
	}

	endereco := models.Endereco{
		ID:       id,
		Cep:      req.Cep,
		Bairro:   req.Bairro,
		Rua:      req.Rua,
		Numero:   req.Numero,
		Cidade:   req.Cidade,
		Estado:   req.Estado,
		PessoaID: &pessoa.ID,
	}

	if err := enderecoRepo.Update(id, &endereco); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Erro ao atualizar endereço",
		})
	}
	return c.JSON(http.StatusOK, endereco)
}

func deleteEndereco(c *echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "ID inválido",
		})
	}
	if err := enderecoRepo.Delete(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Erro ao deletar endereço",
		})
	}
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Endereço deletado",
	})
}

// ===================== MAIN =====================

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // Ou coloque seu http://localhost:5173
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	e.Use(middleware.Recover())

	db := database.InitDB()

	pessoaRepo = repository.NewPessoaRepository(db)
	carroRepo = repository.NewCarroRepository(db)
	enderecoRepo = repository.NewEnderecoRepository(db)

	// Rotas Pessoas
	e.GET("/pessoas", getAllPessoas)
	e.GET("/pessoas/:id", getPessoaByID)
	e.GET("/pessoas/doc/:doc", getPessoaByDoc)
	e.POST("/pessoas", createPessoa)
	e.PUT("/pessoas/:id", updatePessoa)
	e.DELETE("/pessoas/:id", deletePessoa)

	// Rotas Endereços
	e.GET("/enderecos", getAllEnderecos)
	e.GET("/enderecos/pessoa/:pessoaId", getEnderecosByPessoaId)
	e.POST("/enderecos", createEndereco)
	e.PUT("/enderecos/:id", updateEndereco)
	e.DELETE("/enderecos/:id", deleteEndereco)

	// Rotas Carros
	e.GET("/carros", getAllCarros)
	e.PUT("/carros/:id", updateCarro)
	e.POST("/carros", createCarro)
	e.DELETE("/carros/:id", deleteCarro)

	log.Println("🚀 Servidor rodando em http://localhost:8080")
	e.Start(":8080")
}
