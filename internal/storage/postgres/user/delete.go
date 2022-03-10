package user

import (
	"context"
	"fmt"
)

func (s *Store) DeleteByID(ctx context.Context, ID int64) error {
	query := fmt.Sprintf(`
DELETE FROM %s 
WHERE id=$1
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
