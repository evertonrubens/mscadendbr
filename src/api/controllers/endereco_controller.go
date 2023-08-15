package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"msCadEndBr/src/api/models"
	"msCadEndBr/src/db"
	"net/http"

	"github.com/gorilla/mux"
)

// GetAllEnderecosHandler retorna todos os endereços cadastrados
func GetAllEnderecosHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Executando o getAllEnderecoHandler")

	enderecos, err := db.GetAllEnderecos()

	if len(enderecos) == 0 {
		mensagemResposta := fmt.Sprintf("Não existem endereços cadastrados nesta base de dados")
		resposta := models.Mensagem{Message: mensagemResposta}
		responseJSON, err := json.Marshal(resposta)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Erro ao codificar a resposta JSON"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(responseJSON)
		return
	}

	if err != nil {
		log.Println("Erro ao buscar os endereços:", err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(enderecos); err != nil {
		log.Println("Erro ao codificar endereços para JSON:", err)
		http.Error(w, "Erro ao retornar os endereços", http.StatusInternalServerError)
		return
	}
}

// GetEnderecosByCepHandler retorna os endereços com um determinado cep
func GetEnderecosByCepHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	cep := vars["cep"]

	enderecos, err := db.GetEnderecosByCep(cep)

	if len(enderecos) == 0 {
		mensagemResposta := fmt.Sprintf("Não existem endereços cadastrados para o CEP Informado: %s", cep)
		resposta := models.Mensagem{Message: mensagemResposta}
		responseJSON, err := json.Marshal(resposta)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Erro ao codificar a resposta JSON"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(responseJSON)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(enderecos)
}

// GetEnderecoByIdHandler retorna um endereço por id
func GetEnderecoByIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idPF := vars["id"]

	var emptEndereco models.Endereco

	endereco, err := db.GetEnderecoById(idPF)

	if endereco == emptEndereco {

		mensagemResposta := fmt.Sprintf("Nenhum endereço encontrado para este ID fornecido: %v", idPF)
		resposta := models.Mensagem{Message: mensagemResposta}
		responseJSON, err := json.Marshal(resposta)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Erro ao codificar a resposta JSON"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(responseJSON)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(endereco)
}

// GetEnderecosByNomePFHandler retorna os endereços com um determinado nome de pessoa física
func GetEnderecosByNomePFHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	nomePF := vars["nome"]

	enderecos, err := db.GetEnderecosByNomePF(nomePF)

	if len(enderecos) == 0 {
		mensagemResposta := fmt.Sprintf("Não existem endereços cadastrados para esta pessoa física informada: %s", nomePF)
		resposta := models.Mensagem{Message: mensagemResposta}
		responseJSON, err := json.Marshal(resposta)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Erro ao codificar a resposta JSON"))
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNotFound)
		w.Write(responseJSON)
		return
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(enderecos)
}

// CreateEnderecoHandler cria um novo endereço
func CreateEnderecoHandler(w http.ResponseWriter, r *http.Request) {
	var endereco models.Endereco
	err := json.NewDecoder(r.Body).Decode(&endereco)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//endereco.Id = util.GenerateID()
	endereco.Id = ""

	createdEndereco, err := db.CreateEndereco(endereco)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdEndereco)
}

/*
/// UpdateEnderecoHandler atualiza um endereço existente
func UpdateEnderecoHandler(w http.ResponseWriter, r *http.Request) {
	// Obter o ID a ser atualizado a partir dos parâmetros da URL
	id := mux.Vars(r)["id"]

	// Obter o endereço a ser atualizado do corpo da solicitação
	var endereco models.Endereco
	err := json.NewDecoder(r.Body).Decode(&endereco)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Atualizar o endereço no banco de dados
	err = db.UpdateEndereco(id, endereco)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Responder com status 200 OK
	w.WriteHeader(http.StatusOK)
}

// DeleteEnderecoHandler exclui um endereço existente
func DeleteEnderecoHandler(w http.ResponseWriter, r *http.Request) {
	// Obter o ID a ser excluído a partir dos parâmetros da URL
	id := mux.Vars(r)["id"]

	// Excluir o endereço do banco de dados
	err := db.DeleteEndereco(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Responder com status 200 OK
	w.WriteHeader(http.StatusOK)
}

*/
