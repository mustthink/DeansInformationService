package data

import (
	"database/sql"
	"errors"
	"server/internal/types"
	"server/pkg/testJWT"
)

type NewLog struct {
	Jwt      string `json:"jwt"`
	Userrole int    `json:"userRole"`
}

func (m *Service) NewAuth(acc *types.Account) (NewLog, error) {
	stmt := `SELECT id, login, pass, type FROM accounts
    WHERE login = ?`

	l := &NewLog{}
	row := m.DB.QueryRow(stmt, acc.Login)
	exac := &types.Account{}
	err := row.Scan(&exac.Id, &exac.Login, &exac.Password, &exac.Typeacc)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return *l, types.ErrNoRecord
		} else {
			return *l, err
		}
	}
	if acc.Password == exac.Password {
		token, err := testJWT.GenerateJWT(*exac, m.Secret)
		if err != nil {
			return *l, err
		} else {
			l.Jwt = token
			l.Userrole = exac.Typeacc
			return *l, nil
		}
	} else {
		return *l, nil
	}

}

func (m *Service) GetMyAcc(id int) (*types.Account, error) {

	return nil, nil
}

func (m *Service) LatestAccounts() ([]*types.Account, error) {
	return nil, nil
}
