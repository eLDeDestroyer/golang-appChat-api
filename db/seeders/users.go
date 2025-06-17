package seeders

import (
	"backend/model"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
	"math/rand"
	"time"
)

func UserSeeder(db *gorm.DB) {
	rand.Seed(time.Now().UnixNano())

	users := []model.User{
		{
			Name: "John", Username: "John Cena", Password: "12345678",
		},
		{
			Name: "Mary", Username: "Mary Cena", Password: "12345678",
		},
		{
			Name: "Jane", Username: "Jane Cena", Password: "12345678",
		},
		{
			Name: "Billy", Username: "Billy Cena", Password: "12345678",
		},
		{
			Name: "Jeremy", Username: "Jeremy Cena", Password: "12345678",
		},
		{
			Name: "Skibidi", Username: "Skibidi Cena", Password: "12345678",
		},
		{
			Name: "Minto", Username: "Minto Cena", Password: "12345678",
		},
		{
			Name: "Sulo", Username: "Sulo Cena", Password: "12345678",
		},
		{
			Name: "Rumi", Username: "Rumi Cena", Password: "12345678",
		},
		{
			Name: "Jack", Username: "Jack Cena", Password: "12345678",
		},
		{
			Name: "Grealish", Username: "Grealish Cena", Password: "12345678",
		},
		{
			Name: "Naruto", Username: "Naruto Cena", Password: "12345678",
		},
	}

	for _, user := range users {
		user.UniqueNumber = rand.Intn(90000000) + 10000000
		hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Println(err)
		}
		user.Password = string(hashPassword)
		if err := db.Create(&user).Error; err != nil {
			log.Fatalf("Gagal menambahkan user: %v", err)
		}
	}

	fmt.Println("Seeder User berhasil dijalankan!")
}
