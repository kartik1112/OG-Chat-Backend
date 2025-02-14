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
	return err
}

func (u *User) GetUserByEmail() {
	query := `SELECT userId, username, avatarurl, status, createdat,updatedat FROM public.users WHERE email = $1;`
	row := db.DB.QueryRow(query, u.Email)
	row.Scan(&u.UserID, &u.Username, &u.AvatarUrl, &u.Status, &u.CreatedAt, &u.UpdatedAt)
}

func (u *User) UdpateUserByEmail() error {
	query := `UPDATE public.users
			  SET username = $1, avatarurl=$2, status=$3, updatedat=CURRENT_TIMESTAMP
			  WHERE email=$4;
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(u.Username, u.AvatarUrl, u.Status, u.Email)
	if err != nil {
		return err
	}
	return err
}
