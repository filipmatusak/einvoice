package auth

import (
	"errors"
	"github.com/google/uuid"
	"github.com/slovak-egov/einvoice/authproxy/db"
	"github.com/slovak-egov/einvoice/authproxy/user"
	"math/rand"
	"strconv"
)

type UserManager interface {
	Create() *user.User
	Remove(token string) error
	Exists(token string) bool
	GetUser(token string) *user.User
}

type userManager struct {
	db db.AuthDB
}

func NewUserManager(db db.AuthDB) UserManager {
	return userManager{db}
}

func (userManager userManager) Create() *user.User {
	random, _ := uuid.NewRandom()
	token := random.String()

	random, _ = uuid.NewRandom()
	id := strconv.Itoa(rand.Intn(10000000))

	usr := &user.User{Token: token, Id: id}

	userManager.db.Add(usr)

	return usr
}

func (userManager userManager) Remove(token string) error {
	err := userManager.db.Remove(token)
	if err != nil {
		return errors.New("Invalid token")
	}
	return nil
}

func (userManager userManager) Exists(token string) bool {
	return userManager.db.Exists(token)
}

func (userManager userManager) GetUser(token string) *user.User {
	return userManager.db.GetUser(token)
}
