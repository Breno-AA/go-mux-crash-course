package repository

import (
	"database/sql"
	"go-mux-crash-course/entity"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

type sqliteRepo struct{}

func NewSQLiteRepository() PostRepository {
	os.Remove("./posts.db")

	db, err := sql.Open("sqlite3", "./posts.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `
	CREATE TABLE posts (id INTEGER NOT NULL PRIMARY KEY, title TEXT, txt TEXT);
	DELETE FROM posts
	`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
	}
	return &sqliteRepo{}
}

func (sq *sqliteRepo) Save(post *entity.Post) (*entity.Post, error) {
	db, err := sql.Open("sqlite3", "./posts.db")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	stmt, err := tx.Prepare("INSERT INTO posts(id, title, txt) VALUES(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer stmt.Close()
	_, err = stmt.Exec(post.ID, post.Title, post.Text)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return nil, nil
}

// TODO
func (sq *sqliteRepo) Delete(post *entity.Post) error {
	// NOT IMPLEMENTED
	return nil
}

func (sq *sqliteRepo) FindAll() ([]entity.Post, error) {
	// NOT IMPLEMENTED
	return nil, nil
}

func (sq *sqliteRepo) FindByID(ID string) (*entity.Post, error) {
	// NOT IMPLEMENTED
	return nil, nil
}
