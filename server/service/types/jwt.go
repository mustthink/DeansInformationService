package types

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
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

func GenerateJWT(a Account, secret string) (string, error) {
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
