// internal/data/postgres.go
package data

import (
	"database/sql"
	"fmt"
	"io/ioutil"

	// registering database driver
	_ "github.com/lib/pq"
)

func getConnection() (*sql.DB, error) {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=require",
		"ec2-3-225-41-234.compute-1.amazonaws.com",
		5432,
		"pqonavdfbwcoln",
		"c129830082ef74ba9c3cafdca381457105f64a83cfc7e8310d7290766d67533b",
		"d6kov15budmlki")

	/*psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	"127.0.0.1",
	5432,
	"postgres",
	"root",
	"no-country")*/

	//uri := os.Getenv("DATABASE_URI")

	return sql.Open("postgres", psqlInfo)
}

func MakeMigration(db *sql.DB) error {
	b, err := ioutil.ReadFile("./database/models.sql")
	if err != nil {
		return err
	}

	rows, err := db.Query(string(b))
	if err != nil {
		return err
	}

	return rows.Close()
}
