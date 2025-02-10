package controllers

import (
	"curso5/database"
	"curso5/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.JSON(200, models.Alunos)
}

func Saudacao(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz": "E ai " + nome + ", tudo certo?",
	})
}

func CriarAluno(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	_, err := database.DB.Exec("insert into alunos_teste (nome, cpf) values ($1, $2)", aluno.Nome, aluno.Cpf)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "erro ao inserir aluno" + err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Aluno criado com sucesso!",
		"aluno":   aluno.Nome,
	})
}

func BuscaAluno(c *gin.Context) {
	pesquisa := c.Params.ByName("nome")

	for _, v := range models.Alunos {
		if v.Nome == pesquisa {
			c.JSON(200, v)
			break
		} else {
			continue
		}
	}

}
