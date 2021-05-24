package tools

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"server/config/vars"
	"time"
)

type token struct {
	Id      int64  `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Expired int64  `json:"expired,omitempty"`
	jwt.Claims
}

// Create
/// id 签发标识,userId
/// name 签发人,userName
func Create(id int64, name string) string {
	claims := &token{
		Id:      id,
		Name:    name,
		Expired: time.Now().AddDate(0, 0, 7).UnixNano(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	str, err := token.SignedString(vars.KeyToken)
	if err != nil {
		Err("Create", fmt.Sprintf("%s", err))
	}
	return str
}

// parse
/// 解析token的方法
func parse(str string) (claims jwt.Claims) {
	var token *jwt.Token
	token, _ = jwt.Parse(str, func(token *jwt.Token) (interface{}, error) {
		return vars.KeyToken, nil
	})
	claims = token.Claims
	return
}

// Parse
/// 解析token,并返回id
func Parse(value string) interface{} {
	parse := parse(value)
	id := parse.(jwt.MapClaims)["id"]
	return id
}
