package user

import "time"

type User struct {
	ID         int64     `json:"id" db:"id"`
	FirstName  string    `json:"first_name" db:"first_name"`
	SecondName string    `json:"second_name" db:"second_name"`
	NickName   string    `json:"nick_name" db:"nick_name"`
	Email      string    `json:"email" db:"email"`
	Phone      string    `json:"phone,omitempty" db:"phone" validate:"e164"`
	Password   string    `json:"-" db:"password" validate:"min=8"`
	CreatedAt  time.Time `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt  time.Time `json:"updated_at,omitempty" db:"updated_at"`
}

type CreateUserQuery struct {
	FirstName  string `json:"first_name"`
	SecondName string `json:"second_name"`
	NickName   string `json:"nick_name"`
	Email      string `json:"email" db:"email"`
	Phone      string `json:"phone" validate:"e164"`
	Password   string `json:"password" validate:"min=8"`
}

type UpdateUserQuery struct {
	ID int64 `json:"required"`
	CreateUserQuery
}
