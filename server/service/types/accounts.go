package types

import "errors"

var ErrNoRecord = errors.New("account: подходящего аккаунта не найдено")

type Account struct {
	Id       int
	Login    string
	Password string
	Typeacc  int
}
