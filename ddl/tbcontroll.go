package ddl
func (db *DB) DropTable(name string) error {
	statment := "DROP TABLE " + name + ";"
	_, err := db.db.Exec(statment)
	if err != nil {
		return err
	}
	return nil
}
