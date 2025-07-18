package config

import (
	"crud-oportunides/schemas"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Inicializa e retorna a conexão com SQLite
func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")

	// Caminho do banco de dados (relativo ao diretório raiz do projeto)
	dbPath := "./db/main.db"

	// Verifica se o arquivo do banco de dados existe
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("Arquivo do banco não existe. Criando diretório e arquivo.")

		// Cria o diretório ./db com permissões padrão
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			logger.Error("Erro ao criar diretório:", err)
			return nil, err
		}

		// Cria o arquivo do banco de dados (isso é opcional porque o GORM/SQLite cria automaticamente)
		file, err := os.Create(dbPath)
		if err != nil {
			logger.Error("Erro ao criar arquivo do banco de dados:", err)
			return nil, err
		}
		file.Close()
	}

	// Conecta ao banco de dados SQLite usando GORM
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Error("Erro ao conectar ao banco de dados SQLite:", err)
		return nil, err
	}

	// Realiza o auto-mapeamento do schema (criação de tabelas)
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Error("Erro ao migrar schemas:", err)
		return nil, err
	}

	logger.Info("Banco de dados SQLite inicializado com sucesso.")

	return db, nil
}
