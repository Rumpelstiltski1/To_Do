package models

type TaskEvent struct {
	TaskID    int
	Action    string
	Title     *string //Различение «отсутствующего» значения и «нулевого»
	Content   *string
	Status    *bool
	CreatedAt string
}
