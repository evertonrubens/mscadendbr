package models

//import "time"

type Endereco struct {
	Id         string `json: id`
	NomePF     string `json: nomePF`
	Logradouro string `json: logradouro`
	Numero     string `json: numero`
	Bairro     string `json: bairro`
	Cidade     string `json: cidade`
	Uf         string `json: uf`
	Cep        string `json: cep`
	DtCriacao  string `json: dtCriacao`
}

type Mensagem struct {
	Message string `json:"message"`
}
