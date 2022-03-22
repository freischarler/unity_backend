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
		"ec2-52-201-195-11.compute-1.amazonaws.com",
		5432,
		"xlzgnrpqsxzsoy",
		"35db71b3fbd0ff3be8561318ecc7c0f83b7fd934108b67c759db89246bf572b0",
		"d5nd16i4s66ahm")

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
