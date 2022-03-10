package auth

import (
	"repositorie/internal/entities/user"
)

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpireAt     int64  `json:"expire_at"`
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
