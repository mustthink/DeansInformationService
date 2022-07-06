package data

import (
	"database/sql"
	"errors"
	"server/service/types"
	"time"
)

func (m *Service) InsertNews(news *types.News) error {

	stmt := `INSERT INTO news (title, content, created)
    VALUES(?, ?, NOW())`

	_, err := m.DB.Exec(stmt, news.Title, news.Content)
	if err != nil {
		return err
	}

	return nil
}

func (m *Service) GetNews(id int) (*types.News, error) {

	stmt := `SELECT id, title, content, created FROM news
    WHERE id = ?`

	row := m.DB.QueryRow(stmt, id)

	n := &types.News{}
	var t time.Time
	err := row.Scan(&n.ID, &n.Title, &n.Content, &t)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, types.ErrNoRecord
		} else {
			return nil, err
		}
	}
	n.Created = t.Format("2 Jan 2006 15:04")
	return n, nil
}

func (m *Service) LatestNews() ([]*types.News, error) {

	stmt := `SELECT id, title, content, created FROM news
    ORDER BY id DESC LIMIT 10`

	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var news []*types.News

	for rows.Next() {

		n := &types.News{}
		var t time.Time
		err := rows.Scan(&n.ID, &n.Title, &n.Content, &t)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, types.ErrNoRecord
			} else {
				return nil, err
			}
		}
		n.Created = t.Format("2 Jan 2006 15:04")

		news = append(news, n)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return news, nil
}

func (m *Service) DeleteNews(id int) error {

	stmt := `DELETE FROM news
    WHERE id = ?`

	_, err := m.DB.Exec(stmt, id)
	if err != nil {
		return err
	}
	return nil
}
