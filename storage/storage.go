package storage

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

//Instance of storage

type Storage struct {
	// конфиг - как подключиться
	config *Config
	// db fileDescriptor
	db *sql.DB
	// Subfield's for repo interfacing
	UserRepository    *UserRepository
	ArticleRepository *ArticleRepository
}

func New(config *Config) *Storage {
	return &Storage{
		config: config,
	}
}

func (storage *Storage) Open() error {
	db, err := sql.Open("postgres", storage.config.DatabaseURI)

	if err != nil {
		log.Printf("Error: %v", err)
		return err
	}
	if err := db.Ping(); err != nil {
		log.Printf("Error: %v", err)
		return err
	}
	storage.db = db
	log.Println("DB connection created successfully!")
	return nil
}

func (storage *Storage) Close() {
	storage.db.Close()
}

// Public Repo for User

func (s *Storage) User() *UserRepository {
	if s.UserRepository != nil {
		return s.UserRepository
	}
	s.UserRepository = &UserRepository{storage: s}
	return s.UserRepository
}

// Public Repo for Article

func (s *Storage) Article() *ArticleRepository {
	if s.ArticleRepository != nil {
		return s.ArticleRepository
	}
	s.ArticleRepository = &ArticleRepository{storage: s}
	return s.ArticleRepository
}
