package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/conqdat/books-api/internal/config"
	_ "github.com/lib/pq"
)

func Connect(cfg config.DatabaseConfig) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.ConnectionString())
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	// Set connection pool settings
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(25)

	log.Println("Database connected successfully")
	return db, nil
}

func CreateTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS books (
		id SERIAL PRIMARY KEY,
		title VARCHAR(255) NOT NULL,
		author VARCHAR(255) NOT NULL,
		isbn VARCHAR(20) UNIQUE NOT NULL,
		published_at DATE NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);

	-- Create indexes for better performance
	CREATE INDEX IF NOT EXISTS idx_books_title ON books(title);
	CREATE INDEX IF NOT EXISTS idx_books_author ON books(author);
	CREATE INDEX IF NOT EXISTS idx_books_isbn ON books(isbn);

	-- Insert sample data if table is empty
	INSERT INTO books (title, author, isbn, published_at) 
	SELECT 'The Go Programming Language', 'Alan Donovan', '978-0134190440', '2015-11-16'
	WHERE NOT EXISTS (SELECT 1 FROM books);

	INSERT INTO books (title, author, isbn, published_at) 
	SELECT 'Clean Code', 'Robert C. Martin', '978-0132350884', '2008-08-01'
	WHERE NOT EXISTS (SELECT 1 FROM books WHERE isbn = '978-0132350884');

	INSERT INTO books (title, author, isbn, published_at) 
	SELECT 'Design Patterns', 'Gang of Four', '978-0201633612', '1994-10-31'
	WHERE NOT EXISTS (SELECT 1 FROM books WHERE isbn = '978-0201633612');
	`

	if _, err := db.Exec(query); err != nil {
		return fmt.Errorf("failed to create table: %w", err)
	}

	log.Println("Table created and sample data inserted successfully")
	return nil
}