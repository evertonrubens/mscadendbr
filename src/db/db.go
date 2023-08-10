package db

import (
	"database/sql"
	"fmt"
	"log"
	"msCadEndBr/src/api/models"
	tools "msCadEndBr/src/utils"
	"os"
	"regexp"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	// carrega variáveis de ambiente
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	err = godotenv.Load(dir + "/dev.env")
	if err != nil {
		log.Print("Arquivo .env não encontrado.")
	}

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_DATABASE"))
	//log.Println("Conexão string: ", connStr)
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Println("Erro ao tentar a conectividade com o banco de dados: ", connStr)
		log.Fatal(err)
	}

	DB = db

	err = DB.Ping()
	if err != nil {
		log.Println("Não foi possível realizar um ping na base de dados que está sendo apontado no arquivo .env: ", connStr)
		log.Fatal(err)
	}

	log.Println("Conectado no banco de dados com sucesso!")

}

func GetAllEnderecos() ([]models.Endereco, error) {
	InitDB()

	rows, err := DB.Query("SELECT * FROM endereco")
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar os endereços: %v", err)
	}
	defer rows.Close()

	var enderecos []models.Endereco

	for rows.Next() {
		var endereco models.Endereco
		err := rows.Scan(&endereco.Id, &endereco.NomePF, &endereco.Logradouro, &endereco.Numero, &endereco.Bairro, &endereco.Cidade, &endereco.Uf, &endereco.Cep, &endereco.DtCriacao)
		if err != nil {
			return nil, fmt.Errorf("erro ao obter o endereço: %v", err)
		}

		enderecos = append(enderecos, endereco)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("erro ao percorrer os registros do endereço: %v", err)
	}

	log.Println("Consulta de todos os enderecos realizada com sucesso!")

	defer DB.Close()
	log.Println("Desconectado no banco de dados com sucesso!")

	return enderecos, nil
}

// CreateEndereco cria um novo endereço no banco de dados com um ID único gerado a partir de um hash SHA-256
func CreateEndereco(endereco models.Endereco) (models.Endereco, error) {
	InitDB()

	// Gerar hash SHA-256 a partir dos campos do endereço
	id := tools.GenerateID()
	idStr := fmt.Sprintf("%v", id)                  // Converte o id para uma string
	idStr = strings.Trim(idStr, "%!(EXTRA string=") // Remove o prefixo %!(EXTRA string=
	idStr = strings.Trim(idStr, ")")                // Remove o sufixo )
	idStr = strings.TrimSpace(idStr)
	// Usa expressão regular para extrair o hash de 64 caracteres
	r, _ := regexp.Compile("[a-f0-9]{64}")
	idStr = r.FindString(idStr)

	fmt.Println("conseguimos fazer a extração apenas do hash com 64bytes: ", idStr)

	// Inserir endereço no banco de dados
	sqlStatement := `
	  INSERT INTO endereco (id, nomePF, logradouro, numero, bairro, cidade, uf, cep)
	  VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	result, err := DB.Exec(sqlStatement, idStr, endereco.NomePF, endereco.Logradouro, endereco.Numero, endereco.Bairro, endereco.Cidade, endereco.Uf, endereco.Cep)
	if err != nil {
		log.Printf("Erro ao criar endereço no banco de dados: %v", err)
		return models.Endereco{}, err
	}

	idResult, _ := result.LastInsertId()
	endereco.Id = fmt.Sprintf("%d", idResult)

	log.Println("Insert realizado com sucesso!")

	defer DB.Close()
	log.Println("Desconectado no banco de dados com sucesso!")

	return endereco, nil
}

func GetEnderecoById(id string) (models.Endereco, error) {
	InitDB()

	row := DB.QueryRow("SELECT * FROM endereco WHERE id = ?", id)

	var endereco models.Endereco
	err := row.Scan(&endereco.Id, &endereco.NomePF, &endereco.Logradouro, &endereco.Numero, &endereco.Bairro, &endereco.Cidade, &endereco.Uf, &endereco.Cep, &endereco.DtCriacao)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Endereco{}, fmt.Errorf("não foi encontrado nenhum endereço com o ID fornecido: %v", err)
		}
		return models.Endereco{}, fmt.Errorf("erro ao obter o endereço com o ID fornecido: %v", err)
	}

	log.Println("Consulta por ID realizada com sucesso!")

	defer DB.Close()
	log.Println("Desconectado no banco de dados com sucesso!")

	return endereco, nil
}

// GetEnderecosByCep retorna todos os endereços que possuem o CEP informado
func GetEnderecosByCep(cep string) ([]models.Endereco, error) {
	InitDB()

	rows, err := DB.Query("SELECT * FROM endereco WHERE cep = ?", cep)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar endereços por CEP: %v", err)
	}
	defer rows.Close()

	var enderecos []models.Endereco
	for rows.Next() {
		var endereco models.Endereco
		err := rows.Scan(&endereco.Id, &endereco.NomePF, &endereco.Logradouro, &endereco.Numero, &endereco.Bairro, &endereco.Cidade, &endereco.Uf, &endereco.Cep, &endereco.DtCriacao)
		if err != nil {
			return nil, fmt.Errorf("erro ao ler dados de endereço: %v", err)
		}
		enderecos = append(enderecos, endereco)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("erro ao iterar endereços: %v", err)
	}

	log.Println("Consulta por CEP realizada com sucesso!")

	defer DB.Close()
	log.Println("Desconectado no banco de dados com sucesso!")

	return enderecos, nil
}

// GetEnderecosByNomePF retorna todos os endereços que possuem o nome da pessoa física informado
func GetEnderecosByNomePF(nomePF string) ([]models.Endereco, error) {

	InitDB()
	//rows, err := DB.Query("SELECT * FROM endereco WHERE nomePF = ?", nomePF)

	// Prepara a consulta SQL
	sqlStatement := `SELECT id, nomePF, logradouro, numero, bairro, cidade, uf, cep, dtCriacao
		FROM endereco WHERE nomePF = ?`

	// Executa a consulta e verifica se houve erro
	rows, err := DB.Query(sqlStatement, nomePF)
	if err != nil {
		return nil, fmt.Errorf("erro ao buscar endereços por nome da pessoa física: %v", err)
	}
	defer rows.Close()

	//var enderecos []models.Endereco
	enderecos := make([]models.Endereco, 0)
	for rows.Next() {
		var endereco models.Endereco
		err := rows.Scan(&endereco.Id, &endereco.NomePF, &endereco.Logradouro, &endereco.Numero, &endereco.Bairro, &endereco.Cidade, &endereco.Uf, &endereco.Cep, &endereco.DtCriacao)
		if err != nil {
			return nil, fmt.Errorf("erro ao ler dados de endereço: %v", err)
		}
		enderecos = append(enderecos, endereco)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("erro ao iterar endereços: %v", err)
	}

	log.Println("Consulta por nomePF realizada com sucesso!")

	defer DB.Close()
	log.Println("Desconectado no banco de dados com sucesso!")

	return enderecos, nil
}
