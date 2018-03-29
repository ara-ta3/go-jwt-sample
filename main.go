package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"io/ioutil"
)

type User struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

var (
	SigKey          = []byte("hogehoge")
	PublicKeyPath 	= "./app.rsa.pub"
	PrivateKeyPath 	= "./app.rsa"
)

func claimAndParseWithSameKey() error {
	t := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		&User {
			ID: "hoge",
		},
	)
	ts, err := t.SignedString(SigKey)
	if err != nil {
		return err
	}

	fmt.Println(ts)
	tt, err := jwt.Parse(ts, func(t *jwt.Token) (interface{}, error) {
		return SigKey, nil
	})

	if err != nil {
		return err
	}

	fmt.Println("parsed")
	fmt.Printf("%+v\n", tt.Claims)

	_, err = jwt.Parse(ts, func(t *jwt.Token) (interface{}, error) {
		return []byte("aaa"), nil
	})
	fmt.Printf("%+v\n", err)
	return nil
}

func claimAndParseWithKeyPair() error {
	privateKey, err := ioutil.ReadFile(PrivateKeyPath)
	if err != nil {
		return err
	}
	signKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateKey)
	if err != nil {
		return err
	}

	t := jwt.NewWithClaims(
		jwt.SigningMethodRS256,
		&User {
			ID: "hoge",
		},
	)

	ts, err := t.SignedString(signKey)
	fmt.Println(ts)
	if err != nil {
		return err
	}
	tt, err := jwt.Parse(ts, func(t *jwt.Token) (interface{}, error) {
		publicKey, err := ioutil.ReadFile(PublicKeyPath)
		if err != nil {
			return nil, err
		}
		return jwt.ParseRSAPublicKeyFromPEM(publicKey)
	})
	if err != nil {
		return err
	}
	fmt.Println("parsed")
	fmt.Printf("%+v\n", tt.Claims)

	return nil
}

func main() {
	fmt.Println("same key")
	e := claimAndParseWithSameKey()
	if e != nil {
		fmt.Printf("%+v\n", e)
	}
	fmt.Println("key pair")
	e = claimAndParseWithKeyPair()
	if e != nil {
		fmt.Printf("%+v\n", e)
	}
}
