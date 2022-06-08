package handlers

import (
	"database/sql"
	"server/types"
)

type AccountModel struct {
	DB *sql.DB
}

func (m *AccountModel) Insert(title, content, expires string) (int, error) {
	return 0, nil
}

// Get - Метод для возвращения данных заметки по её идентификатору ID.
func (m *AccountModel) Get(id int) (*types.Account, error) {
	return nil, nil
}

// Latest - Метод возвращает 10 наиболее часто используемые заметки.
func (m *AccountModel) Latest() ([]*types.Account, error) {
	return nil, nil
}
