package data

import (
	"database/sql"
	"fmt"

	// PostgreSQL driver
	"github.com/lib/pq"
	_ "github.com/lib/pq"

	"github.com/davidzaragoza/tangelo/pkg/domain"
)

const (
	createTableImages = `CREATE TABLE IF NOT EXISTS images(
		url text PRIMARY KEY,
		content bytea
	)`
)

type Repository struct {
	config *domain.Configuration
	db     *sql.DB
}

func NewRepository(config *domain.Configuration) (domain.Repository, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s dbname=%s search_path=%s user=%s password=%s sslmode=%s",
		config.Database.Host, config.Database.Port, config.Database.DBName, config.Database.Schema, config.Database.User, config.Database.Password, config.Database.SSLMode))
	if err != nil {
		return nil, err
	}
	if _, err := db.Exec(fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s", config.Database.Schema)); err != nil {
		return nil, err
	}
	if _, err := db.Exec(createTableImages); err != nil {
		return nil, err
	}
	return &Repository{config: config, db: db}, nil
}

func (rep *Repository) SaveImage(url string, content []byte) error {
	stmt, err := rep.db.Prepare("INSERT INTO images (url, content) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	if _, err := stmt.Exec(url, content); err != nil {
		if postgreErr, ok := err.(*pq.Error); ok {
			// This error happens when primary key exist
			if postgreErr.Code == "23505" {
				stmtUpdate, err := rep.db.Prepare("UPDATE images SET content = $1")
				if err != nil {
					return err
				}
				if _, err := stmtUpdate.Exec(content); err != nil {
					return err
				}
				return nil
			}
		}
		return err
	}
	return nil
}

func (rep *Repository) GetImage(url string) ([]byte, error) {
	return nil, nil
}
