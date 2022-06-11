package data

import (
	"database/sql"
	"errors"
	"server/types"
)

func (m *Service) InsertStudent(fio string, keygroup, expires int) (int, error) {

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

func (m *Service) GetStudent(id int) (*types.Student, error) {

	stmt := `SELECT id, fio, keygroup, expires FROM students
    WHERE expires > CURRENT_DATE() AND id = ?`

	row := m.DB.QueryRow(stmt, id)

	s := &types.Student{}

	err := row.Scan(&s.ID, &s.FIO, &s.GROUP.KEYgroup, &s.EXPIRES)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, types.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (m *Service) LatestStudents() ([]*types.Student, error) {

	stmt := `SELECT id, fio, keygroup, expires FROM students
    WHERE expires > CURDATE() ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var students []*types.Student

	for rows.Next() {

		s := &types.Student{}

		err = rows.Scan(&s.ID, &s.FIO, &s.GROUP.KEYgroup, &s.EXPIRES)
		if err != nil {
			return nil, err
		}

		students = append(students, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return students, nil
}
