package models

import (
	"errors"
	"time"

	"github.com/kartik1112/OG-Chat-Backend/db"
	"github.com/kartik1112/OG-Chat-Backend/utils"
)

type User struct {
	UserID       int
	Username     string `binding:"required"`
	Email        string `binding:"required"`
	PasswordHash string `binding:"required"`
	AvatarUrl    string
	Status       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u *User) CreateUser() error {
	query := `INSERT INTO public.users(userId, username, email, passwordHash, avatarUrl, status, createdat, updatedat) VALUES (
		nextval('users_userid_seq'::regclass),$1 ,$2  ,$3  ,$4  ,$5  , CURRENT_TIMESTAMP, CURRENT_TIMESTAMP
		);`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	hashedPassword, _ := utils.GeneratePasswordHash(u.PasswordHash)
	_, err = stmt.Exec(u.Username, u.Email, hashedPassword, u.AvatarUrl, u.Status)
	return err
}

func (u *User) ValidateUser() error {
	query := `SELECT userid,username, email, passwordhash
				FROM public.users WHERE username = $1 OR email = $2;`
	row := db.DB.QueryRow(query, u.Username, u.Email)
	var password string
	err := row.Scan(&u.UserID, &u.Username, &u.Email, &password)
	if err != nil {
		return errors.New("invalid credentials")
	}
	err = utils.ValidatePassword(password, u.PasswordHash)
	// return err
	// if username == u.Username || email == u.Email {
	// }

	return err
}
