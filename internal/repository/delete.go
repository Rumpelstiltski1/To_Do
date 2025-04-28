package repository

import "fmt"

func (s *Storage) DeleteTask(id int) error {
	result, err := s.db.Exec("delete from tasks where id=$1", id)
	if err != nil {
		return err
	}
	rowsAffectd, _ := result.RowsAffected()
	if rowsAffectd == 0 {
		return fmt.Errorf("задача с ID %d не найдена", id)
	}
	return nil
}
