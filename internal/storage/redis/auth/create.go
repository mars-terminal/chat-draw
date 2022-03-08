package auth

import "context"

func (s *Storage) SetToken(ctx context.Context, token string) error {
	if cmd := s.db.Set(ctx, s.prefix+token, "", 0); cmd.Err() != nil {
		return cmd.Err()
	}

	return nil
}
