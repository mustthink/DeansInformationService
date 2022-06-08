package data

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

func (m *AccountModel) Get(id int) (*types.Account, error) {
	return nil, nil
}

func (m *AccountModel) Latest() ([]*types.Account, error) {
	return nil, nil
}
