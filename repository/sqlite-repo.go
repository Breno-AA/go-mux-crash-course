package repository

import (
	"database/sql"
	"go-mux-crash-course/entity"
	"log"

	_ "modernc.org/sqlite"
)

type sqliteRepo struct{}

func NewSQLiteRepository() PostRepository {
	db, err := sql.Open("sqlite", "./posts.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := `CREATE TABLE IF NOT EXISTS posts (id INTEGER NOT NULL PRIMARY KEY, title TEXT, txt TEXT)`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
	}
	return &sqliteRepo{}
}

func (sq *sqliteRepo) Save(post *entity.Post) (id int64, err error) {
	db, err := sql.Open("sqlite", "./posts.db")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO posts(id, title, txt) VALUES(?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer stmt.Close()

	result, err := stmt.Exec(post.ID, post.Title, post.Text)
	if err != nil {
		log.Fatal(err)
		return
	}
	return result.LastInsertId()
}

// TODO
func (sq *sqliteRepo) Delete(ID string) (err error) {
	db, err := sql.Open("sqlite", "./posts.db")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer db.Close()

	db.QueryRow("DELETE FROM posts WHERE id = ?", ID)
	return
}

func (sq *sqliteRepo) FindAll() (posts []entity.Post, err error) {
	db, err := sql.Open("sqlite", "./posts.db")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT * FROM posts")
	if err != nil {
		log.Fatal(err)
		return
	}

	for rows.Next() {
		var post entity.Post
		rows.Scan(&post.ID, &post.Title, &post.Text)
		posts = append(posts, post)
	}
	return
}

func (sq *sqliteRepo) FindByID(ID string) (post entity.Post, err error) {
	db, err := sql.Open("sqlite", "./posts.db")
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	err = db.QueryRow("SELECT * FROM posts WHERE id = ?", ID).Scan(&post.ID, &post.Title, &post.Text)
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}
