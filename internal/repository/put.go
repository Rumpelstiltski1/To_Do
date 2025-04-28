package repository

import "fmt"

func (s *Storage) PutTask(status bool, id int) error {
	result, err := s.db.Exec("update tasks set status= $1 WHERE id=$2", status, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("не удалось получить количество изменённых строк: %w", err)
	}
	if rowsAffected == 0 {
		return fmt.Errorf("задача с ID %d не найдена", id)
	}
	return nil
}
