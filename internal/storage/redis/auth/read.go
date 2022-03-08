package auth

import "context"

func (s *Storage) IsTokenExists(ctx context.Context, token string) (bool, error) {
	cmd := s.db.Get(ctx, s.prefix+token)

	if err := cmd.Err(); err != nil {
		return false, err
	}

	return cmd.String() != "", nil
}

func (s *Storage) IsRefreshTokenExists(ctx context.Context, token string) (bool, error) {
	cmd := s.db.Get(ctx, s.prefix+token)

	if err := cmd.Err(); err != nil {
		return false, err
	}

	return cmd.String() != "", nil
}
