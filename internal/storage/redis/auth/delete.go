package auth

import (
	"context"
)

func (s *Storage) DeleteToken(ctx context.Context, token string) error {
	return s.db.Del(ctx, s.prefix+token).Err()
}
