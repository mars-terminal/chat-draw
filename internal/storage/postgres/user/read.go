package user

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/mars-terminal/chat-draw/internal/entities"
	"github.com/mars-terminal/chat-draw/internal/entities/user"
)

func (s *Store) GetByID(ctx context.Context, ID int64) (*user.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", s.table)

	rows, err := s.db.QueryxContext(ctx, query, ID)
	if err != nil {
		return nil, fmt.Errorf("failet to get by id (user storage): %w", err)
	}

	u := user.User{}
	for rows.Next() {
		err := rows.StructScan(&u)
		if err != nil {
			return nil, err
		}
	}

	return &u, nil
}

func (s *Store) GetByIDs(ctx context.Context, IDs []int64) ([]*user.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id in ($1)", s.table)

	rows, err := s.db.QueryxContext(ctx, query, IDs)
	if err != nil {
		return nil, fmt.Errorf("failet to get by id's (user storage): %w", err)
	}

	var users []*user.User
	for rows.Next() {
		var u user.User
		err := rows.StructScan(&u)
		if err != nil {
			return nil, err
		}
		users = append(users, &u)
	}

	return users, nil
}

func (s *Store) GetByNickName(ctx context.Context, nickName string) ([]*user.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE nick_name ILIKE '%%?%%'", s.table)

	rows, err := s.db.QueryxContext(ctx, query, nickName)
	if err != nil {
		return nil, fmt.Errorf("failet to get by nick name (user storage): %w", err)
	}

	var users []*user.User
	for rows.Next() {
		var u user.User
		err := rows.StructScan(&u)

		if err != nil {
			return nil, err
		}

		users = append(users, &u)
	}

	return users, nil
}

func (s *Store) GetByNickNameStrict(ctx context.Context, nickName string) (*user.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE nick_name = $1", s.table)

	row := s.db.QueryRowxContext(ctx, query, nickName)

	var u user.User
	err := row.StructScan(&u)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user: %w", entities.ErrNotFound)
		}
		return nil, fmt.Errorf("failed to get by strict nickname: %w", err)
	}

	return &u, nil
}

func (s *Store) GetByEmailAndPasswordHash(ctx context.Context, email, passwordHash string) (*user.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE email=$1 AND password = $2", s.table)

	row := s.db.QueryRowxContext(ctx, query, email, passwordHash)
	if err := row.Err(); err != nil {
		return nil, err
	}

	var u user.User
	err := row.StructScan(&u)
	if err != nil {
		return nil, fmt.Errorf("failed to get by email and password: %w", err)
	}

	return &u, nil
}

func (s *Store) GetByPhone(ctx context.Context, phone string) ([]*user.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE phone=$1", s.table)

	rows, err := s.db.QueryxContext(ctx, query, phone)
	if err != nil {
		return nil, fmt.Errorf("failet to get by phone (user storage): %w", err)
	}
	var users []*user.User
	for rows.Next() {
		var u user.User
		err := rows.StructScan(&u)

		if err != nil {
			return nil, err
		}
		users = append(users, &u)
	}
	return users, nil
}
