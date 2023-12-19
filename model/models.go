package model

import(
	"github.com/dgrijalva/jwt-go"
)

type User struct {
	Email string
	Password string
}
type UserTmp struct{
	Email string
	Code string
}
type Claims struct{
	ID string
	jwt.StandardClaims
}

type Fundmt struct{
	Code int `json:"code"`
	Message string `json:"message"`
}
type Token struct{
	Fundmt
	Token string `json:"token"`
}