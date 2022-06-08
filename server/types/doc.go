package types

import (
	"io/ioutil"
	"os"
	"strings"
)

var keywords = []string{"DAY", "MONTH", "YEAR", "NUMBER", "SURNAME", "FIRSTNAME", "MIDDLENAME", "GROUP", "YEARSOFSTUDY", "FACULTY", "INFORMATION", "LDAY", "LMONTH", "LYEAR", "DEAN", "MANAGER"}

func СreateDoc(info []string) map[string]string {
	doc := map[string]string{}
	for i, word := range keywords {
		doc[word] = info[i]
	}
	return doc
}

func GenerateDoc(namefile string, doc map[string]string) {
	f, err := ioutil.ReadFile("html/" + namefile + ".htm")
	if err != nil {
		panic(err)
	}

	file := string(f)
	for _, word := range keywords {
		file = strings.Replace(file, word, doc[word], 1)
	}

	nf, err := os.Create("html/temp.htm")
	if err != nil {
		panic(err)
	}
	nf.WriteString(file)
	nf.Close()
}

func DeleteDoc() {
	os.Remove("html/temp.htm")
}
