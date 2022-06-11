package data

import (
	"server/service/types"
)

func (m *Service) InsertAccount(title, content, expires string) (int, error) {
	return 0, nil
}

func (m *Service) GetAccount(id int) (*types.Account, error) {
	return nil, nil
}

func (m *Service) LatestAccounts() ([]*types.Account, error) {
	return nil, nil
}
