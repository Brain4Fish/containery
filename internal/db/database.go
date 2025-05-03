package db

import (
	"database/sql"
	"fmt"
	"log"
	_ "modernc.org/sqlite"
	"os"
)

type Image struct {
	ID       string
	Name     string
	Tag      string
	Selected bool
}

type Database struct {
	Connection   *sql.DB
	DatabasePath string `default:"settings.db"`
}

func (db *Database) dbFileExists() bool {
	fmt.Println("Checking path: ", db.DatabasePath)
	_, err := os.Stat(db.DatabasePath)
	return !os.IsNotExist(err)
}

func (db *Database) initDatabase() {
	query := `CREATE TABLE IF NOT EXISTS images (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			tag TEXT NOT NULL,
			selected BOOLEAN DEFAULT FALSE,
			UNIQUE(name, tag)
    	);`
	_, err := db.Connection.Exec(query)
	if err != nil {
		log.Fatal("There is an error on DB initialization: ", err)
	}
	// Sample data below
	sampleData := []string{
		`INSERT INTO images (name, tag) VALUES ('nginx', 'latest');`,
		`INSERT INTO images (name, tag) VALUES ('nginx', '1.28');`,
		`INSERT INTO images (name, tag, selected) VALUES ('httpd', '2.4', true);`,
	}
	for _, query := range sampleData {
		if _, err := db.Connection.Exec(query); err != nil {
			log.Fatal("Unable to insert sample data: ", err)
		}
	}
}

func (db *Database) CreateConnection() {
	var err error
	db.Connection, err = sql.Open("sqlite", "settings.db")
	if err != nil {
		log.Fatal("Failed to open DB: ", err)
	}
	if !db.dbFileExists() {
		fmt.Println("Must init database")
		db.initDatabase()
	}
}

func (db *Database) ReadData() ([]Image, error) {
	rows, err := db.Connection.Query("SELECT id, name, tag, selected FROM images")
	if err != nil {
		log.Fatal("There is an error in select query: ", err)
	}
	defer rows.Close()

	var images []Image
	for rows.Next() {
		var img Image
		err = rows.Scan(&img.ID, &img.Name, &img.Tag, &img.Selected)
		if err != nil {
			log.Fatal("Can't map data from select: ", err)
			return nil, err
		}
		images = append(images, img)
	}
	fmt.Print(images)
	return images, nil
}
