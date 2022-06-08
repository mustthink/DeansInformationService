package data

import (
	"database/sql"
	"server/types"
)

type StudentModel struct {
	DB *sql.DB
}

func (m *StudentModel) Insert(fio string, keygroup, expires int) (int, error) {

	stmt := `INSERT INTO students (fio, keygroup, expires)
    VALUES(?, ?, DATE_ADD(CURDATE(), INTERVAL ? DAY))`

	result, err := m.DB.Exec(stmt, fio, keygroup, expires)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (m *StudentModel) Get(id int) (*types.Student, error) {
	return nil, nil
}

func (m *StudentModel) Latest() ([]*types.Student, error) {
	return nil, nil
}
