package repository

import (
	"database/sql"
	"go-mux-crash-course/entity"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type mysqlRepo struct{}

func NewMySQLRepository() PostRepository {
	dsn := "root:@tcp(localhost:3306)/posts"

	// Open a connection to MySQL
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the "posts" table
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS posts (
		id INT NOT NULL AUTO_INCREMENT,
		title VARCHAR(255),
		txt TEXT,
		PRIMARY KEY (id)
	);
	`

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s\n", err, sqlStmt)
	}

	return &mysqlRepo{}
}

func (mr *mysqlRepo) Save(post *entity.Post) (*entity.Post, error) {
	dsn := "root:@tcp(localhost:3306)/posts"

	// Open a connection to MySQL
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Start a transaction
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Prepare the INSERT statement
	stmt, err := tx.Prepare("INSERT INTO posts(title, txt) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer stmt.Close()

	// Execute the INSERT statement
	result, err := stmt.Exec(post.Title, post.Text)
	if err != nil {
		log.Fatal(err)
		tx.Rollback()
		return nil, err
	}

	// Get the last inserted ID
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
		tx.Rollback()
		return nil, err
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// Update the post ID
	post.ID = int(lastInsertID)

	return post, nil
}

func (mr *mysqlRepo) FindAll() ([]entity.Post, error) {
	dsn := "root:@tcp(localhost:3306)/posts"

	// Open a connection to MySQL
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer db.Close()

	// Execute the SELECT statement
	rows, err := db.Query("SELECT id, title, txt FROM posts")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer rows.Close()

	var posts []entity.Post

	// Iterate through the result set
	for rows.Next() {
		var post entity.Post
		err := rows.Scan(&post.ID, &post.Title, &post.Text)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		posts = append(posts, post)
	}

	// Check for errors during iteration
	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return posts, nil
}
