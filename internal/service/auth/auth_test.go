package auth

import (
	"repositorie/internal/service"
	"repositorie/internal/storage"
	"testing"
	"time"
)

func TestService_generateRefreshToken(t *testing.T) {
	type fields struct {
		salt           string
		signInKey      string
		tokenTTL       time.Duration
		authStorage    storage.AuthStorage
		userStorage    storage.UserStorage
		userService    service.UserService
		messageService service.MessageService
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "ok",
			fields: fields{},
			want:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Service{
				salt:           tt.fields.salt,
				signInKey:      tt.fields.signInKey,
				tokenTTL:       tt.fields.tokenTTL,
				authStorage:    tt.fields.authStorage,
				userStorage:    tt.fields.userStorage,
				userService:    tt.fields.userService,
				messageService: tt.fields.messageService,
			}
			if got, _ := s.generateRefreshToken(); got != tt.want {
				t.Errorf("generateRefreshToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
