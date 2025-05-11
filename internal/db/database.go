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
	Source   string
	Selected bool
}

type Credentials struct {
	ID       uint
	Username string
	Password string
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
	queries := []string{`CREATE TABLE IF NOT EXISTS images (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			tag TEXT NOT NULL,
			source TEXT NOT NULL,
			destination INTEGER NOT NULL,
			selected BOOLEAN DEFAULT FALSE,
			UNIQUE(name, tag),
			FOREIGN KEY (destination) REFERENCES destinations(id) ON DELETE CASCADE
    	);`,
		`CREATE TABLE IF NOT EXISTS destinations (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL UNIQUE,
			url TEXT NOT NULL,
			authtype INTEGER NOT NULL,
			credentials INTEGER,
            FOREIGN KEY (authtype) REFERENCES auth_types(id) ON DELETE CASCADE,
            FOREIGN KEY (credentials) REFERENCES credentials(id) ON DELETE CASCADE
    	);`,
		`CREATE TABLE IF NOT EXISTS auth_types (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			type TEXT NOT NULL UNIQUE
    	);`,
		`CREATE TABLE IF NOT EXISTS credentials (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username TEXT NOT NULL UNIQUE,
			password TEXT
    	);`,
	}
	for _, query := range queries {
		if _, err := db.Connection.Exec(query); err != nil {
			log.Fatal("There is an error on DB initialization: ", err)
		}
	}
	// Sample data below
	sampleData := []string{
		`INSERT INTO auth_types (id, type) VALUES (1, 'anonymous'), (2, 'password');`,
		`INSERT INTO credentials (id, username, password) VALUES (1, 'support', 'testpass');`,
		`INSERT INTO credentials (id, username, password) VALUES (2, 'newuser', 'testpass1');`,
		`INSERT INTO destinations (id, name, url, authtype, credentials) VALUES (1, 'adm_gitlab', 'gitlab.adm-systems.tech:5050', 2, 1);`,
		`INSERT INTO destinations (id, name, url, authtype) VALUES (2, 'docker.io', 'docker.io', 1);`,
		`INSERT INTO images (name, tag, source, destination) VALUES ('nginx', 'latest', 'docker.local', 1);`,
		`INSERT INTO images (name, tag, source, destination) VALUES ('nginx', '1.28', 'docker.local', 2);`,
		`INSERT INTO images (name, tag, source, destination, selected) VALUES ('httpd', '2.4', 'docker.local', 2, true);`,
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

func (db *Database) ReadImagesData() ([]Image, error) {
	rows, err := db.Connection.Query("SELECT id, name, tag, selected FROM images")
	if err != nil {
		log.Fatal("There is an error in image select query: ", err)
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

func (db *Database) ReadCredentialsData() ([]Credentials, error) {
	rows, err := db.Connection.Query("SELECT id, username, password FROM credentials")
	if err != nil {
		log.Fatal("There is an error in credentials select query: ", err)
	}
	defer rows.Close()

	var credsList []Credentials
	for rows.Next() {
		var cred Credentials
		err = rows.Scan(&cred.ID, &cred.Username, &cred.Password)
		if err != nil {
			log.Fatal("Can't map credentials data from select: ", err)
			return nil, err
		}
		credsList = append(credsList, cred)
	}
	fmt.Print(credsList)
	return credsList, nil
}
