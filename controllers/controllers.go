package controllers

import (
	"curso5/models"

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
