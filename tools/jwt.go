package tools

import (
	"time"

	"server/config/vars"

	"github.com/dgrijalva/jwt-go"
)

type token struct {
	Id      int64  `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Expired int64  `json:"expired,omitempty"`
	jwt.Claims
}

// Create
/// generate token
func Create(id int64, name string) string {
	claims := &token{
		Id:      id,
		Name:    name,
		Expired: time.Now().AddDate(0, 0, 7).UnixNano(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	str, err := token.SignedString(vars.KeyToken)
	if err != nil {
		Err("Create", err)
	}
	return str
}

// parse
/// method of parsing token
func parse(str string) (claims jwt.Claims) {
	var token *jwt.Token
	token, _ = jwt.Parse(str, func(token *jwt.Token) (interface{}, error) {
		return vars.KeyToken, nil
	})
	claims = token.Claims
	return
}

// Parse
/// parsing token and return user's id
func Parse(value string) interface{} {
	parse := parse(value)
	id := parse.(jwt.MapClaims)["id"]
	return id
}
