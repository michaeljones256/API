package database

import (
	"database/sql"
	"fmt"
)

type Database struct {
	db *sql.DB
}

const (
	DB_USER     = "postgres"
	DB_PASSWORD = "easy123"
	DB_NAME     = "GolangApp"
)

// easy123 db master
// same for invidual servers
// using pgadmin
func Connect() (*Database, error) {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	return &Database{db: db}, err
}

func (database *Database) SelectAllTable() string {

	rows, err := database.db.Query("SELECT * FROM testTable")
	var finalName string
	var finalAge int
	if err != nil {
		for rows.Next() {
			var name string
			var age int
			err = rows.Scan(&name, &age)
			fmt.Println(name)
			fmt.Println(age)
			finalName = name
			finalAge = age
		}
	}
	return fmt.Sprintf("%s and %d", finalName, finalAge)

}

func (database *Database) Insert(item interface{}) {
	fmt.Println("this is an insert")
}
func (database *Database) Get(id int) {
	fmt.Println("this is a get")
}
