package types

type groupinfo struct {
	keygroup int
	mentorid int
	flow     int
	numgr    int
	course   string
}

type Student struct {
	id    int
	name  string
	group groupinfo
}
