package data

import (
	"server/types"
)

func (m *Service) InsertMentor(title, content, expires string) (int, error) {
	return 0, nil
}

func (m *Service) GetMentor(id int) (*types.Account, error) {
	return nil, nil
}

func (m *Service) LatestMentors() ([]*types.Account, error) {
	return nil, nil
}
