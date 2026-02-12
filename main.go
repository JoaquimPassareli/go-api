package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type Pessoa struct {
	Nome   string  `json:"nome"`
	Idade  int     `json:"idade"`
	Altura float64 `json:"altura"`
	Doc    int     `json:"doc"`
}

var pessoas = make(map[int]Pessoa)

// GET /pessoas/:doc - busca uma pessoa pelo documento.
func getPessoa(c *echo.Context) error {
	docParam := c.Param("doc")
	doc, err := strconv.Atoi(docParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "doc inválido, use apenas números",
		})
	}

	p, ok := pessoas[doc]
	if !ok {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "pessoa não encontrada",
		})
	}

	return c.JSON(http.StatusOK, p)
}

// POST /pessoas - cria uma nova pessoa.
func createPessoa(c *echo.Context) error {
	var p Pessoa
	if err := c.Bind(&p); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "JSON inválido",
		})
	}

	if _, exists := pessoas[p.Doc]; exists {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "usuário já cadastrado",
		})
	}

	pessoas[p.Doc] = p
	return c.JSON(http.StatusCreated, p)
}

func deletePessoa(c *echo.Context) error {
	docParam := c.Param("doc")
	doc, err := strconv.Atoi(docParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "doc inválido, tente outro número",
		})
	}

	if _, ok := pessoas[doc]; !ok {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "pessoa não encontrada",
		})
	}

	delete(pessoas, doc)

	return c.JSON(http.StatusOK, map[string]string{
		"message": "Usuário deletado com sucesso",
		"doc":     docParam,
	})
}

// PUT /pessoas/:doc - atualiza uma pessoa existente.
func updatePessoa(c *echo.Context) error {
	docParam := c.Param("doc")
	doc, err := strconv.Atoi(docParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "doc inválido, tente outro número",
		})
	}

	var p Pessoa
	if err := c.Bind(&p); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "JSON inválido",
		})
	}

	if _, ok := pessoas[doc]; !ok {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "pessoa não encontrada",
		})
	}

	// Garante que o Doc venha da URL, não do corpo.
	p.Doc = doc
	pessoas[doc] = p

	return c.JSON(http.StatusOK, p)
}

func main() {
	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())

	e.GET("/pessoas/:doc", getPessoa)
	e.POST("/pessoas", createPessoa)
	e.DELETE("/pessoas/:doc", deletePessoa)
	e.PUT("/pessoas/:doc", updatePessoa)

	if err := e.Start(":8080"); err != nil {
		log.Fatal(err)
	}
}
