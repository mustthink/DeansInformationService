package testJWT

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"log"
	"server/internal/types"
	"strings"
	"time"
)

type header struct {
	Typ string `json:"type"`
	Alg string `json:"alg"`
}

type claims struct {
	Name    string    `json:"name"`
	Created time.Time `json:"iat"`
	Expired time.Time `json:"exp"`
	Role    int       `json:"role"`
	UserId  int       `json:"user_id"`
}

func BaseToUrl(s string) string {
	s = strings.ReplaceAll(s, "+", "-")
	s = strings.ReplaceAll(s, "/", "_")
	return strings.ReplaceAll(s, "=", "%")
}

func UrlToBase(s string) string {
	s = strings.ReplaceAll(s, "-", "+")
	s = strings.ReplaceAll(s, "_", "/")
	return strings.ReplaceAll(s, "%", "=")
}

func GenerateJWT(a types.Account, secret string) (string, error) {
	key := []byte(secret)
	hs := hmac.New(sha256.New, key)

	h, err := json.Marshal(&header{
		Typ: "JWT",
		Alg: "HS256",
	})
	if err != nil {
		return "", err
	}

	c, err := json.Marshal(&claims{
		Name:    a.Login,
		Created: time.Now(),
		Expired: time.Now().Add(time.Hour),
		Role:    a.Typeacc,
		UserId:  a.Id,
	})
	if err != nil {
		return "", err
	}

	hc := base64.StdEncoding.EncodeToString(h) + "." + base64.StdEncoding.EncodeToString(c)
	hs.Write([]byte(hc))
	t := hc + "." + base64.StdEncoding.EncodeToString(hs.Sum(nil))
	return BaseToUrl(t), nil
}

func VerificationJWT(token string, secret string) bool {
	key := []byte(secret)
	hs := hmac.New(sha256.New, key)
	hcs := strings.Split(UrlToBase(token), ".")
	if checkExpJWT(hcs[0]) {
		hc := strings.Join(hcs[0:1], ".")
		hs.Write([]byte(hc))
		if strings.Compare(hcs[2], base64.StdEncoding.EncodeToString(hs.Sum(nil))) == 0 {
			return true
		}
	}
	return false
}

func checkExpJWT(c string) bool {
	cl, err := base64.StdEncoding.DecodeString(c)
	if err != nil {
		log.Println(err)
		return false
	}
	newcl := &claims{}
	err = json.Unmarshal(cl, &newcl)
	if err != nil {
		log.Println(err)
		return false
	}
	if newcl.Expired.Sub(time.Now()) < 0 {
		return false
	}
	return true
}

func InfoFromJWT(token string) (int, int) {
	jsonclaims, err := base64.StdEncoding.DecodeString(strings.Split(UrlToBase(token), ".")[1])
	cl := &claims{}
	if err != nil {
		log.Print(err)
		return 0, 0
	}
	err = json.Unmarshal(jsonclaims, &cl)
	if err != nil {
		log.Print(err)
		return 0, 0
	}
	return cl.UserId, cl.UserId
}
