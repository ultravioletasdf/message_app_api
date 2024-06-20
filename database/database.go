package database

import (
	"database/sql"
	"fmt"
	"log"
	"message_app_api/utils"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	"github.com/tursodatabase/go-libsql"
)

func local(config *utils.Config) (*sql.DB, error) {
	return sql.Open("sqlite3", config.LocalDatabaseName)
}
func turso(config *utils.Config) (*sql.DB, string, error) {
	dir, err := os.MkdirTemp("", "libsql-*")
	if err != nil {
		return nil, "", err
	}
	fmt.Println("Created temporary directory " + dir)

	path := filepath.Join(dir, config.TursoDatabaseName)
	connector, err := libsql.NewEmbeddedReplicaConnector(path, config.TursoUrl, libsql.WithAuthToken(config.TursoToken))
	if err != nil {
		return nil, "", err
	}

	return sql.OpenDB(connector), dir, nil
}
func ConnectToDatabase(config *utils.Config) (*sql.DB, string) {
	if config.UseLocalDatabase == true {
		db, err := local(config)
		if err != nil {
			log.Fatalln("Failed to connect to local database", err.Error())
		}
		return db, ""
	}
	turso, dir, err := turso(config)
	if err != nil {
		log.Fatalln("Failed to connect to turso database", err.Error())
	}
	return turso, dir
}
