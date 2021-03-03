package repositories

import (
	"log"

	"github.com/ragil000/go-restful.git/entities"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// UserRepository is contract what userRepository can do to database
type UserRepository interface {
	InsertUser(user entities.User) entities.User
	UpdateUser(user entities.User) entities.User
	VerifyCredential(email string, password string) interface{}
	IsDuplicateEmail(email string) (tx *gorm.DB)
	FindByEmail(email string) entities.User
	ProfileUser(userID string) entities.User
}

type userConnection struct {
	connection *gorm.DB
}

// NewUserRespository is creates a new instance of UserRepository
func NewUserRespository(db *gorm.DB) UserRepository {
	return &userConnection{
		connection: db,
	}
}

func (db *userConnection) InsertUser(user entities.User) entities.User {
	user.Password = hashAndSalt(([]byte(user.Password)))
	db.connection.Save(&user)
	return user
}

func (db *userConnection) UpdateUser(user entities.User) entities.User {
	if user.Password != "" {
		user.Password = hashAndSalt(([]byte(user.Password)))
	} else {
		var tempUser entities.User
		db.connection.Find(&tempUser, user.ID)
		user.Password = tempUser.Password
	}
	db.connection.Save(&user)
	return user
}

func (db *userConnection) VerifyCredential(email string, password string) interface{} {
	var user entities.User
	res := db.connection.Where("email = ?", email).Take(&user)
	if res.Error == nil {
		return user
	}
	return nil
}

func (db *userConnection) IsDuplicateEmail(email string) (tx *gorm.DB) {
	var user entities.User
	return db.connection.Where("email = ?", email).Take(&user)
}

func (db *userConnection) FindByEmail(email string) entities.User {
	var user entities.User
	db.connection.Where("email = ?", email).Take(&user)
	return user
}

func (db *userConnection) ProfileUser(userID string) entities.User {
	var user entities.User
	db.connection.Preload("Books.User").Find(&user, userID)
	return user
}

func hashAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}
