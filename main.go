package main

import (
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

var rp *repository.PessoaRepository
var rp2 *repository.CarroRepository

// ===================== PESSOAS =====================

// GET /pessoas
func getAllPessoas(c *echo.Context) error {
	pessoas, err := rp.Read()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Erro ao buscar pessoas",
		})
	}
	return c.JSON(http.StatusOK, pessoas)
}

// GET /pessoas/:doc
func getPessoa(c *echo.Context) error {
	docParam := c.Param("doc")
	doc, err := strconv.Atoi(docParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "doc inválido",
		})
	}

	p, err := rp.ReadById(doc)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Pessoa não encontrada",
		})
	}

	return c.JSON(http.StatusOK, p)
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

	if err := rp.Create(&p); err != nil {
		log.Println("Erro Create Pessoa:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Erro ao criar pessoa",
		})
	}

	return c.JSON(http.StatusCreated, p)
}

// DELETE /pessoas/:doc
func deletePessoa(c *echo.Context) error {
	docParam := c.Param("doc")
	doc, err := strconv.Atoi(docParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "doc inválido",
		})
	}

	if err := rp.Delete(doc); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Erro ao deletar",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Pessoa deletada",
	})
}

// PUT /pessoas/:doc
func updatePessoa(c *echo.Context) error {
	docParam := c.Param("doc")
	doc, err := strconv.Atoi(docParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "doc inválido",
		})
	}

	var p models.Pessoa
	if err := c.Bind(&p); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "JSON inválido",
		})
	}

	p.Doc = doc

	if err := rp.Update(&p); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Erro ao atualizar",
		})
	}

	return c.JSON(http.StatusOK, p)
}

// ===================== CARROS =====================

// GET /carros
func getAllCarros(c *echo.Context) error {
	carros, err := rp2.Read()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Erro ao buscar carros",
		})
	}
	return c.JSON(http.StatusOK, carros)
}

// GET /carros/:id
func getCarro(c *echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "id inválido",
		})
	}

	carro, err := rp2.ReadById(id)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Carro não encontrado",
		})
	}

	return c.JSON(http.StatusOK, carro)
}

// POST /carros
func createCarro(c *echo.Context) error {
	var carro models.Carro

	if err := c.Bind(&carro); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "JSON inválido",
		})
	}

	if err := rp2.Create(&carro); err != nil {
		log.Println("Erro Create Carro:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Erro ao criar carro",
		})
	}

	return c.JSON(http.StatusCreated, carro)
}

// DELETE /carros/:id
func deleteCarro(c *echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "id inválido",
		})
	}

	if err := rp2.Delete(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Erro ao deletar",
		})
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Carro deletado",
	})
}

// PUT /carros/:id
func updateCarro(c *echo.Context) error {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "id inválido",
		})
	}

	var carro models.Carro
	if err := c.Bind(&carro); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "JSON inválido",
		})
	}

	carro.ID = id

	if err := rp2.Update(&carro); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Erro ao atualizar",
		})
	}

	return c.JSON(http.StatusOK, carro)
}

// ===================== MAIN =====================

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:5173"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
	}))

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	db := database.InitDB()

	rp = repository.NewPessoaRepository(db)
	rp2 = repository.NewCarroRepository(db)

	// Rotas Pessoas
	e.GET("/pessoas", getAllPessoas)
	e.GET("/pessoas/:doc", getPessoa)
	e.POST("/pessoas", createPessoa)
	e.DELETE("/pessoas/:doc", deletePessoa)
	e.PUT("/pessoas/:doc", updatePessoa)

	// Rotas Carros
	e.GET("/carros", getAllCarros)
	e.GET("/carros/:id", getCarro)
	e.POST("/carros", createCarro)
	e.DELETE("/carros/:id", deleteCarro)
	e.PUT("/carros/:id", updateCarro)

	log.Println("🚀 Servidor rodando em http://localhost:8080")

	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
