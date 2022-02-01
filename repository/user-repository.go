package repository

import (
	"log"

	"github.com/evan-nurandiz/golang-gin-rest-api/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository interface {
	InsertUser(user entity.User) entity.User
}

type UserConnection struct {
	connection *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserConnection{
		connection: db,
	}
}

func (db *UserConnection) InsertUser(user entity.User) entity.User {
	user.Password = HashPassword([]byte(user.Password))
	db.connection.Save(&user)
	return user
}

func HashPassword(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to Has Password")
	}

	return string(hash)
}
