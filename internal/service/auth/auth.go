package auth

import (
	"context"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"

	"github.com/mars-terminal/chat-draw/internal/entities/auth"
	"github.com/mars-terminal/chat-draw/internal/entities/user"
	"github.com/mars-terminal/chat-draw/internal/service"
	"github.com/mars-terminal/chat-draw/internal/storage"
)

type Service struct {
	salt     string
	signKey  string
	tokenTTL time.Duration

	authStorage storage.AuthStorage

	userService    service.UserService
	messageService service.MessageService
}

func NewService(
	salt string,
	signInKey string,
	tokenTTL time.Duration,

	authStorage storage.AuthStorage,
	userService service.UserService,
	messageService service.MessageService,
) *Service {

	return &Service{
		salt:           salt,
		signKey:        signInKey,
		tokenTTL:       tokenTTL,
		authStorage:    authStorage,
		userService:    userService,
		messageService: messageService,
	}
}

func (s *Service) generateAccessToken(user *user.User, expireAt int64) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: expireAt,
		Subject:   strconv.FormatInt(user.ID, 10),
	})

	return token.SignedString([]byte(s.signKey))
}

func (s *Service) generateRefreshToken() (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(uuid.New().String()), 1)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (s *Service) generateTokens(ctx context.Context, u *user.User) (*auth.Tokens, error) {
	expireAt := time.Now().Add(s.tokenTTL).Unix()

	accessToken, err := s.generateAccessToken(u, expireAt)
	if err != nil {
		return nil, err
	}
	refreshToken, err := s.generateRefreshToken()
	if err != nil {
		return nil, err
	}

	if err = s.authStorage.SetToken(ctx, accessToken, u.ID, s.tokenTTL); err != nil {
		return nil, err
	}

	if err = s.authStorage.SetToken(ctx, refreshToken, u.ID, s.tokenTTL*2); err != nil {
		return nil, err
	}

	return &auth.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		ExpireAt:     expireAt,
	}, nil
}
