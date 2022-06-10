package types

import "errors"

var ErrNoRecord = errors.New("account: подходящего аккаунта не найдено")

type Account struct {
	id       int
	login    string
	password string
	token    string
	typeacc  uint8
}
