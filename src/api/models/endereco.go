package models

//import "time"

type Endereco struct {
	Id         string
	NomePF     string
	Logradouro string
	Numero     string
	Bairro     string
	Cidade     string
	Uf         string
	Cep        string
	DtCriacao  string
}
