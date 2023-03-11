package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/amitshekhariitbhu/go-backend-clean-architecture/domain"
)

type taskRepository struct {
	database   *sql.DB
	collection string
}

func NewTaskRepository(db *sql.DB, collection string) domain.TaskRepository {
	return &taskRepository{
		database:   db,
		collection: collection,
	}
}

const (
	insertTask   = "INSERT INTO tasks (title, user_id, created_at) VALUES (?, ?, ?)"
	selectByUser = "SELECT * FROM tasks WHERE user_id=?"
)

func (tr *taskRepository) Create(c context.Context, task *domain.Task) error {
	res, err := tr.database.Exec(
		insertTask,
		task.Title,
		task.UserID,
		time.Now(),
	)

	if err != nil {
		return err
	} else {
		id, err := res.LastInsertId()
		if err != nil {
			println("Error:", err.Error())
			return err
		} else {
			println("Inserted Task ID:", id)
		}
	}
	return err
}

func (tr *taskRepository) FetchByUserID(c context.Context, userID int64) ([]domain.Task, error) {
	var tasks []domain.Task
	rows, err := tr.database.QueryContext(c, selectByUser, userID)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		item := domain.Task{}

		err = rows.Scan(&item.ID, &item.Title, &item.UserID, &item.CreatedAt, &item.UpdatedAt)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, item)
	}

	rows.Close()
	return tasks, nil
}
