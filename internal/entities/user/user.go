package user

import "time"

type User struct {
	ID         int64     `db:"id"`
	FirstName  string    `db:"first_name"`
	SecondName string    `db:"second_name"`
	NickName   string    `db:"nick_name"`
	Phone      string    `db:"phone"`
	Password   string    `db:"password"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

type CreateUserQuery struct {
	FirstName  string
	SecondName string
	NickName   string
	Phone      string
	Password   string
}

type UpdateUserQuery struct {
	ID int64
	CreateUserQuery
}
