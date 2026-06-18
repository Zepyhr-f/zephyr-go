package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	hash := "$2a$10$7JB720yubVSZvUIV7EqK/.VqGOZTH.ulu33dHOiBE8ByOhJIrdAu2"
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte("123456"))
	fmt.Println(err)
}
