package postgres

import (
	"context"
	"fmt"
	"repositorie/internal/entities"
)

type UserStore struct {
	*Store

	table string
}

func NewUserStore(store *Store, table string) *UserStore {
	return &UserStore{
		Store: store,
		table: table,
	}
}

func (u *UserStore) GetByID(ctx context.Context, ID int64) (*entities.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=? LIMIT 1", u.table)

	rows, err := u.db.QueryxContext(ctx, query, ID)
	if err != nil {
		return nil, err
	}

	user := entities.User{}
	for rows.Next() {
		err := rows.StructScan(&user)
		if err != nil {
			return nil, err
		}
	}

	return &user, nil
}

func (u *UserStore) GetByIDs(ctx context.Context, IDs []int64) ([]*entities.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE id in (?)", u.table)

	rows, err := u.db.QueryxContext(ctx, query, IDs)
	if err != nil {
		return nil, err
	}

	var users []*entities.User
	for rows.Next() {
		var user entities.User
		err := rows.StructScan(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}

func (u *UserStore) GetByNickName(ctx context.Context, nickName string) ([]*entities.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE nick_name=? ", u.table)

	rows, err := u.db.QueryxContext(ctx, query, nickName)
	if err != nil {
		return nil, err
	}

	var users []*entities.User
	for rows.Next() {
		var user entities.User
		err := rows.StructScan(&user)

		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}

	return users, nil
}

func (u *UserStore) GetByPhone(ctx context.Context, phone string) ([]*entities.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE phone=?", u.table)

	rows, err := u.db.QueryxContext(ctx, query, phone)
	if err != nil {
		return nil, err
	}
	var users []*entities.User
	for rows.Next() {
		var user entities.User
		err := rows.StructScan(&user)

		if err != nil {
			return nil, err
		}
		users = append(users, &user)
	}
	return users, nil
}

func (u *UserStore) Create(ctx context.Context, q *entities.CreateUserQuery) (*entities.User, error) {
	query := fmt.Sprintf(`
INSERT INTO %s 
	(first_name, second_name, nick_name, phone, password)
VALUES 
	(?, ?, ?, ?, ?)
RETURNING id
`, u.table)

	result, err := u.db.ExecContext(ctx, query, q.FirstName, q.SecondName, q.NickName, q.Phone, q.Password)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user, err := u.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserStore) Update(ctx context.Context, q *entities.UpdateUserQuery) (*entities.User, error) {
	query := fmt.Sprintf(`
INSERT INTO %s
	(first_name, second_name, nick_name, phone, password)
VALUES 
	(?, ?, ?, ?, ?)
RETURNING id
`, u.table)

	result, err := u.db.Exec(query, q.FirstName, q.SecondName, q.NickName, q.Phone, q.Password)
	if err != nil {
		return nil, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	user, err := u.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *UserStore) DeleteByID(ctx context.Context, ID int64) error {
	query := fmt.Sprintf(`
DELETE FROM %s 
WHERE id=?
`, u.table)

	result, err := u.db.ExecContext(ctx, query, ID)
	if err != nil {
		return err
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
