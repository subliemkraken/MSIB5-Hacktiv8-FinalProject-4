package seed

import (
	"FinalProject4/model/entity"

	"golang.org/x/crypto/bcrypt"
)

var passwordHash, _ = bcrypt.GenerateFromPassword([]byte("iniadmin"), bcrypt.MinCost)

var User = entity.User{
	FullName: "iniadmin",
	Email:    "iniadmin@mail.com",
	Password: string(passwordHash),
	Role:     "admin",
	Balance:  0,
}
