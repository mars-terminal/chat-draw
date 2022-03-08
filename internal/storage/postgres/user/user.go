package user

import (
	"context"
	"fmt"
	"repositorie/internal/entities/user"
	"time"

	"github.com/jmoiron/sqlx"
)

type Store struct {
	db *sqlx.DB

	table string
}

func NewUserStore(db *sqlx.DB, table string) *Store {
	return &Store{
		db:    db,
		table: table,
	}
}

func (s *Store) GetByID(ctx context.Context, ID int64) (*user.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=?", s.table)

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
	query := fmt.Sprintf("SELECT * FROM %s WHERE id in (?)", s.table)

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

func (s *Store) GetByNickNameAndPasswordHash(ctx context.Context, nickName, passwordHash string) (*user.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE nick_name=? AND password=?", s.table)

	rows, err := s.db.QueryxContext(ctx, query, nickName, passwordHash)
	if err != nil {
		return nil, fmt.Errorf("failet to get by nick name and password: %w", err)
	}

	var u *user.User
	for rows.Next() {
		if err := rows.StructScan(&u); err != nil {
			return nil, err
		}
	}

	return u, nil
}

func (s *Store) GetByPhone(ctx context.Context, phone string) ([]*user.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE phone=?", s.table)

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

func (s *Store) Create(ctx context.Context, q *user.CreateUserQuery) (*user.User, error) {
	query := fmt.Sprintf(`
INSERT INTO %s 
	(first_name, second_name, nick_name, phone, password)
VALUES 
	(?, ?, ?, ?, ?)
RETURNING id
`, s.table)

	result, err := s.db.ExecContext(ctx, query, q.FirstName, q.SecondName, q.NickName, q.Phone, q.Password)
	if err != nil {
		return nil, fmt.Errorf("failet to create (user storage): %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failet to last insert id in create (user storage): %w", err)
	}

	return &user.User{
		ID:         id,
		FirstName:  q.FirstName,
		SecondName: q.SecondName,
		NickName:   q.NickName,
		Phone:      q.Phone,
		Password:   q.Password,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}, nil
}

func (s *Store) Update(ctx context.Context, q *user.UpdateUserQuery) (*user.User, error) {
	query := fmt.Sprintf(`
UPDATE %s 
SET 
	first_name=?,
	secont_name = ?,
	nick_name = ?,
	phone = ?,
	password = ?
WHERE 
	id = ?
`, s.table)

	result, err := s.db.ExecContext(ctx, query, q.FirstName, q.SecondName, q.NickName, q.Phone, q.Password, q.ID)
	if err != nil {
		return nil, fmt.Errorf("failet to update (user storage): %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, fmt.Errorf("failet to last insert id in update (user store): %w", err)
	}

	message, err := s.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failet to get by id in update (user store): %w", err)
	}

	return message, nil
}

func (s *Store) DeleteByID(ctx context.Context, ID int64) error {
	query := fmt.Sprintf(`
DELETE FROM %s 
WHERE id=?
`, s.table)

	result, err := s.db.ExecContext(ctx, query, ID)
	if err != nil {
		return fmt.Errorf("failet to delete (user storage): %w", err)
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if affected == 0 {
		return fmt.Errorf("not found")
	}

	return nil
}
