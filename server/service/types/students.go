package types

import "time"

type groupinfo struct {
	KEYgroup int
	mentorid int
	flow     int
	numgr    int
	course   string
}

type Student struct {
	ID      int
	FIO     string
	GROUP   groupinfo
	EXPIRES time.Time
}
