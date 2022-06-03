package types

import (
	"io/ioutil"
	"os"
	"strings"
)

/*
DAY MONTH YEAR NUMBER SURNAME FIRSTNAME MIDDLENAME GROUP, YEARSOFSTUDY FACULTY INFORMATION LDAY LMONTH LYEAR DEAN MANAGER
*/

type doc struct {
	DAY          uint8
	MONTH        uint8
	YEAR         uint8
	NUMBER       int
	SURNAME      string
	FIRSTNAME    string
	MIDDLENAME   string
	GROUP        string
	YEARSOFSTUDY uint8
	FACULTY      string
	INFORMATION  string
	LDAY         uint8
	LMONTH       uint8
	DEAN         string
	MANAGER      string
}

func newDoc(DAY uint8, MONTH uint8, YEAR uint8, NUMBER int, SURNAME string, FIRSTNAME string, MIDDLENAME string, GROUP string, YEARSOFSTUDY uint8, FACULTY string, INFORMATION string, LDAY uint8, LMONTH uint8, DEAN string, MANAGER string) *doc {
	return &doc{
		DAY:          DAY,
		MONTH:        MONTH,
		YEAR:         YEAR,
		NUMBER:       NUMBER,
		SURNAME:      SURNAME,
		FIRSTNAME:    FIRSTNAME,
		MIDDLENAME:   MIDDLENAME,
		GROUP:        GROUP,
		YEARSOFSTUDY: YEARSOFSTUDY,
		FACULTY:      FACULTY,
		INFORMATION:  INFORMATION,
		LDAY:         LDAY,
		LMONTH:       LMONTH,
		DEAN:         DEAN,
		MANAGER:      MANAGER}
}

func (d doc) generateDoc(namefile string) {
	f, err := ioutil.ReadFile("html/" + namefile + ".htm")
	if err != nil {
		panic(err)
	}
	file := strings.Replace(string(f), "DAY", string(d.DAY), 1)
	nf, err := os.Open("html/temperary.htm")
	if err != nil {
		panic(err)
	}
	nf.WriteString(file)
	nf.Close()
}

func (d doc) deleteDoc() {
	os.Remove("html/temperary.htm")
}
