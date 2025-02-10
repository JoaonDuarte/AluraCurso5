package models

type Aluno struct {
	Id   int    `json: "id"`
	Nome string `json : "nome"`
	Cpf  string `json : "cpf"`
}

var Alunos []Aluno
