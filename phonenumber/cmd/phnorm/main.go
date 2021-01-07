package main

import (
	"database/sql"
	"fmt"
	norm "github.com/gophercises/phonenumber/normalize"
	_ "github.com/lib/pq"
	"log"
)

//todo move db work stuff to seperate package
const tableName = "phones"

func setupDB() (*sql.DB, error) {
	const (
		host     = "localhost"
		port     = 5432
		user     = "denis"
		password = "a" //any non empty string
		dbname   = "gophercises"
	)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected!")

	return db, nil
}

func resetTable(db *sql.DB, name string) error {
	fmt.Println("Table reset attempt")
	sqlStatement := `DROP TABLE IF EXISTS ` + name
	_, err := db.Exec(sqlStatement)
	if err != nil {
		return err
	}

	sqlStatement = `CREATE TABLE IF NOT EXISTS ` + name + ` (
		id SERIAL primary key,
		phone VARCHAR(32) 
    )`
	_, err = db.Exec(sqlStatement)
	return err
}

type phoneRecord struct {
	id int
	phone string
}

func getPhones(db *sql.DB) ([]phoneRecord, error) {
	var entries []phoneRecord
	var sqlStatement = "SELECT id, phone FROM " + tableName
	rows, err := db.Query(sqlStatement)//, "phone")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var phone phoneRecord
		if err = rows.Scan(&phone.id, &phone.phone); err != nil {
			return nil, err
		}

		entries = append(entries, phone)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return entries, nil
}

func findPhone(db *sql.DB, phone string) (*phoneRecord, error) {
	var sqlStatement = "SELECT id, phone FROM " + tableName + " WHERE phone = $1"

	row := db.QueryRow(sqlStatement, phone)
	if row != nil {
		var p phoneRecord
		err := row.Scan(&p.id, &p.phone)
		if err != nil {
			return nil, err
		}

		return &p, nil
	}

	return nil, nil
}

func updPhone(db *sql.DB, phone phoneRecord) error {
	sqlStatement := "UPDATE " + tableName + " SET phone = $2 WHERE id = $1"

	_, err := db.Exec(sqlStatement, phone.id, phone.phone)
	if err != nil {
		return err
	}

	return nil
}

func delPhone(db *sql.DB, id int) error {
	sqlStatement := "DELETE FROM " + tableName + " WHERE id = $1"

	_, err := db.Exec(sqlStatement, id)
	if err != nil {
		return err
	}

	return nil
}

func addPhone(db *sql.DB, phone string) error {
	sqlStatement := `INSERT INTO ` + tableName + ` (phone)` +
		` VALUES ($1)`

	_, err := db.Exec(sqlStatement, phone)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	db, err := setupDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err = resetTable(db, tableName); err != nil {
		log.Fatal(err)
	}

	var rawPhones = []string {
		"1234567890",
		"123 456 7891",
		"(123) 456 7892",
		"(123) 456-7893",
		"123-456-7894",
		"123-456-7890",
		"1234567892",
		"(123)456-7892",
	}

	//fill with initial data
	for _, ph := range rawPhones {
		err = addPhone(db, ph)
		if err != nil {
			log.Fatal(err)
		}
	}

	//read
	dbPhones, err := getPhones(db)
	if err != nil {
		log.Fatal(err)
	}

	//normalize and write
	for _, ph := range dbPhones {
		phNorm, err := norm.Normalize(ph.phone)
		if err != nil {
			log.Println(err)
			continue
		}
		if phNorm != ph.phone {
			//find existing
			existingPhone, err := findPhone(db, phNorm)
			if err != nil && err != sql.ErrNoRows {
				log.Println(err)
				continue
			}
			if existingPhone != nil {
				//if found, delete current
				err = delPhone(db, ph.id)
				if err != nil {
					log.Println(err)
					continue
				}
			} else {
				//else update with formatted
				err = updPhone(db, phoneRecord{ph.id, phNorm})
				if err != nil {
					log.Println(err)
					continue
				}
			}
		} else {
			fmt.Println("no change required")
		}
	}

}
