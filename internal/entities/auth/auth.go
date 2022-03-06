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
	NickName, Password string
}

type SignUpQuery struct {
	user.CreateUserQuery
}

type UserWithToken struct {
	user.User
	Tokens
}
