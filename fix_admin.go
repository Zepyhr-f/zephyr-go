package main

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type SysUser struct {
	Id       int64  `gorm:"column:id;primary_key"`
	Code     string `gorm:"column:code"`
	NickName string `gorm:"column:nick_name"`
	RealName string `gorm:"column:real_name"`
	Password string `gorm:"column:password"`
}

func (SysUser) TableName() string {
	return "zephyr_sys_user"
}

func main() {
	dsn := "postgres://postgres:postgres@127.0.0.1:5432/zephyr_db?sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	var users []SysUser
	db.Find(&users)

	fmt.Println("Existing Users:")
	for _, u := range users {
		fmt.Printf("ID: %d, Code: %s, NickName: %s, RealName: %s, Password: %s\n", u.Id, u.Code, u.NickName, u.RealName, u.Password)
	}

	// Update admin password
	hash, err := bcrypt.GenerateFromPassword([]byte("123456"), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("failed to generate bcrypt hash: %v", err)
	}

	err = db.Model(&SysUser{}).Where("code = ?", "admin").Or("nick_name = ?", "admin").Update("password", string(hash)).Error
	if err != nil {
		log.Fatalf("failed to update admin password: %v", err)
	}
	fmt.Printf("Updated admin password to bcrypt hash of '123456': %s\n", string(hash))
}
