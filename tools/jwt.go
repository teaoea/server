package tools

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"server/config/vars"
	"time"
)

type token struct {
	Id        int64 `json:"id,omitempty"`
	CreatedAt int64 `json:"created_at,omitempty"`
	jwt.Claims
}

// Create
/// id 签发标识,userId
/// name 签发人,userName
func Create(id int64) string {
	claims := &token{
		Id:        id,
		CreatedAt: time.Now().UnixNano(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	str, err := token.SignedString(vars.KeyToken)
	if err != nil {
		Err("Create", fmt.Sprintf("%s", err))
	}
	return str
}

// Parse
/// 解析token
func Parse(str string) (claims jwt.Claims) {
	var token *jwt.Token
	token, _ = jwt.Parse(str, func(token *jwt.Token) (interface{}, error) {
		return vars.KeyToken, nil
	})
	claims = token.Claims
	return
}
