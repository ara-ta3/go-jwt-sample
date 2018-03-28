package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

type User struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

func claimAndParse() error {
	t := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		&User {
			ID: "hoge",
		},
	)
	sig := []byte("hogehoge")
	ts, err := t.SignedString(sig)
	if err != nil {
		return err
	}

	fmt.Println(ts)
	tt, err := jwt.Parse(ts, func(t *jwt.Token) (interface{}, error) {
		return sig, nil
	})

	if err != nil {
		return err
	}

	fmt.Println(tt)

	_, err = jwt.Parse(ts, func(t *jwt.Token) (interface{}, error) {
		return []byte("aaa"), nil
	})
	return err
}

func main() {
	e := claimAndParse()
	fmt.Printf("%+v\n", e)
}
