package types

import "errors"

var ErrNoRecord = errors.New("account: подходящей записи не найдено")

type account struct {
	id       int
	login    string
	password string
	token    string
	typeacc  uint8
}

/*func newAccount (nid int, nlog, npass string, ntype uint8){
	return &account{
		id: nid,
		login: nlog,
		password: npass,
		typeacc: ntype
	}
}*/
