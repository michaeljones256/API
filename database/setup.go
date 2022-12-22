package database

import (
	"database/sql"
	"fmt"
)

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "easy123"
	DB_NAME     = "GolangApp"
)

// easy123 db master
// same for invidual servers
// using pgadmin
func Connect() (*sql.DB, error) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	if err != nil {
		return &sql.DB{}, err
	}

	return db, nil
}

func SelectAllTable(db *sql.DB) {

	rows, err := db.Query("SELECT * FROM testTable")
	if err != nil {
		for rows.Next() {
			var name string
			var age int
			err = rows.Scan(&name, &age)
			fmt.Println(name)
			fmt.Println(age)
		}
	}

}
