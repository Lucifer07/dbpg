package ddl

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (db *DB) Connect() (*sql.DB, error) {
	err:=db.insertData()
	if err != nil {
		return nil, err
	}
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", db.host, db.port, db.username, db.password, db.dbName,
	)
	database, err := sql.Open("pgx", psqlInfo)
	if err != nil {

		return nil, err
	}

	err = database.Ping()
	if err != nil {
		return nil, err
	}
	db.db = database
	return database, nil
}
func (db *DB) insertData() error {
	dataenv, err := readEnvFile(".env")
	if err != nil {
		return err
	}
	db.host = dataenv["DB_HOST"]
	port, err := strconv.Atoi(dataenv["DB_PORT"])
	if err != nil {
		return err
	}
	db.port = port
	db.dbName = dataenv["DB_DATABASE"]
	db.username = dataenv["DB_USERNAME"]
	db.password = dataenv["DB_PASSWORD"]
	return nil
}
func readEnvFile(filename string) (map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	envVars := make(map[string]string)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "=", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			envVars[key] = value
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return envVars, nil
}
