package ddl

import "database/sql"

func (db *DB) CreateDB(name string) (*sql.DB,error) {
	statment := "CREATE DATABASE " + name + ";"
	_, err := db.db.Exec(statment)
	if err != nil {
		return nil,err
	}
	db.dbName=name
	newDb, err := db.Connect()
	if err != nil {
		return nil,err
	}
	db.db = newDb
	return db.db,nil
}
func (db *DB) DropDB(name string) error {
	db.dbName = "postgres"
	newDb, err := db.Connect()
	if err != nil {
		return err
	}
	db.db = newDb
	statment := "DROP DATABASE " + name + ";"
	_, err = db.db.Exec(statment)
	if err != nil {
		return err
	}
	return nil
}
