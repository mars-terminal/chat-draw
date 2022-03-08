package auth

import (
	"repositorie/internal/entities/user"
)

type Tokens struct {
	AccessToken  string
	RefreshToken string
	ExpireAt     int64
}

type SignInQuery struct {
	NickName string `json:"nick_name" validate:"min=1"`
	Password string `json:"password" validate:"min=8"`
}

type SignUpQuery struct {
	user.CreateUserQuery
}

type UserWithToken struct {
	user.User
	Tokens
}
