package auth

import "context"

func (s *Storage) GetUserIDByToken(ctx context.Context, token string) (int64, error) {
	cmd := s.db.Get(ctx, s.prefix+token)

	if err := cmd.Err(); err != nil {
		return 0, err
	}

	return cmd.Int64()
}

func (s *Storage) IsRefreshTokenExists(ctx context.Context, token string) (bool, error) {
	cmd := s.db.Get(ctx, s.prefix+token)

	if err := cmd.Err(); err != nil {
		return false, err
	}

	return cmd.String() != "", nil
}
