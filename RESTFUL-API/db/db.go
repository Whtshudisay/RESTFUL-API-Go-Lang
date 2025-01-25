package db

import (
	"database/sql"
	"log"
	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite", "api.db")
	if err != nil {
		log.Fatal("Could not connect to the database: ", err)
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsersTable := `CREATE TABLE IF NOT EXISTS 
    users( 
	id INTEGER PRIMARY KEY AUTOINCREMENT, 
	email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL)`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		log.Fatal("Could not User's table: ", err)
	}

	createEventsTable := `CREATE TABLE IF NOT EXISTS 
    events(
    id INTEGER PRIMARY KEY AUTOINCREMENT, 
    name TEXT NOT NULL,
    description TEXT NOT NULL, 
    location TEXT NOT NULL, 
    dateTime DATETIME NOT NULL, 
    user_id INTEGER,
	FOREIGN KEY(user_id) REFERENCES users(id))`

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		log.Fatal("Could not create table: ", err)
	}

	createRegistrationsTable := `CREATE TABLE IF NOT EXISTS registrations (id INTEGER PRIMARY KEY AUTOINCREMENT, 
	Event_id INTEGER, user_id INTEGER, 
	FOREIGN KEY(event_id) REFERENCES events(id), 
	FOREIGN KEY(user_id) REFERENCES users(id))`

	_, err = DB.Exec(createRegistrationsTable)
	if err != nil {
		log.Fatal("Could not Registration's table: ", err)
	}

}
