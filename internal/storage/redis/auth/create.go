package auth

import (
	"context"
	"time"
)

func (s *Storage) SetToken(ctx context.Context, token string, userID int64, ttl time.Duration) error {
	if cmd := s.db.Set(ctx, s.prefix+token, userID, ttl); cmd.Err() != nil {
		return cmd.Err()
	}

	return nil
}
