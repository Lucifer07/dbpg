package migration

import (
	"database/sql"
	"strings"
)

type Table struct {
	query strings.Builder
}

func (t *Table) Create(name string, cols []*Col) *Table {
	t.query.WriteString("CREATE TABLE ")
	t.query.WriteString(name)
	t.query.WriteString("(")
	for i := 0; i < len(cols); i++ {
		q := cols[i].query.String()
		t.query.WriteString(q)
		if string(q[len(q)-1]) != ";" {
			t.query.WriteString(",")
		}
	}
	return t
}
func (t *Table) Execute(db *sql.DB) error {
	dbStart, err := db.Begin()
	if err != nil {
		return err
	}
	_, err = db.Exec(t.query.String())
	if err != nil {
		// dbStart.Rollback()
		return err
	}
	err = dbStart.Commit()
	if err != nil {
		// dbStart.Rollback()
		return err
	}
	return nil
}
