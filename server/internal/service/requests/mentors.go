package data

import (
	"database/sql"
	"errors"
	"server/internal/types"
)

func (m *Service) InsertMentor(mentor *types.Mentor) error {
	stmt := `INSERT INTO mentors (fio, chair)
    VALUES(?, ?)`

	_, err := m.DB.Exec(stmt, mentor.FIO, mentor.CHAIR)
	if err != nil {
		return err
	}

	return nil
}

func (m *Service) GetMentor(id int) (*types.Mentor, error) {
	stmt := `SELECT id, fio, chair FROM mentors
    WHERE id = ?`

	row := m.DB.QueryRow(stmt, id)

	s := &types.Mentor{}

	err := row.Scan(&s.ID, &s.FIO, &s.CHAIR)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, types.ErrNoRecord
		} else {
			return nil, err
		}
	}

	return s, nil
}

func (m *Service) LatestMentors() ([]*types.Account, error) {
	return nil, nil
}
