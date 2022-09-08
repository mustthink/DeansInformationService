package data

import (
	"server/internal/types"
)

func (m *Service) InsertGroup(title, content, expires string) (int, error) {
	return 0, nil
}

func (m *Service) GetGroup(id int) (*types.Account, error) {
	return nil, nil
}

func (m *Service) LatestGroups() ([]*types.Account, error) {
	return nil, nil
}
