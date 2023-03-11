package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/amitshekhariitbhu/go-backend-clean-architecture/domain"
)

type userRepository struct {
	database   *sql.DB
	collection string
}

func NewUserRepository(db *sql.DB, collection string) domain.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}

const (
	insertUser            = "INSERT INTO users (name, email, password, created_at) VALUES (?, ?, ?, ?)"
	selectUserWithoutPass = "SELECT id, name, email, created_at FROM users WHERE id = ?"
	selectUserWithPass    = "SELECT id, name, email, password, created_at FROM users WHERE email=?"
)

func (ur *userRepository) Create(c context.Context, user *domain.User) error {
	res, err := ur.database.Exec(
		insertUser,
		user.Name,
		user.Email,
		user.Password,
		time.Now(),
	)
	if err != nil {
		return err
	} else {
		id, err := res.LastInsertId()
		if err != nil {
			return err
		} else {
			println("Inserted User ID:", id)
		}
	}
	return err
}

func (ur *userRepository) fetch(c context.Context, query string, includePass bool, args ...interface{}) ([]*domain.User, error) {
	rows, err := ur.database.QueryContext(c, query, args...)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	payload := make([]*domain.User, 0)

	for rows.Next() {
		data := new(domain.User)
		var err error
		if includePass {
			err = rows.Scan(&data.ID, &data.Name, &data.Email, &data.Password, &data.CreatedAt)
		} else {
			err = rows.Scan(&data.ID, &data.Name, &data.Email, &data.CreatedAt)
		}
		if err != nil {
			return nil, err
		}
		payload = append(payload, data)
	}
	return payload, nil
}

func (ur *userRepository) FetchAll(c context.Context) ([]*domain.User, error) {
	return ur.fetch(c, selectUserWithoutPass, false, 1)
}

func (ur *userRepository) GetByEmail(c context.Context, email string) (*domain.User, error) {
	var user *domain.User

	rows, err := ur.fetch(c, selectUserWithPass, true, email)
	if err != nil {
		return user, err
	}

	if len(rows) > 0 {
		user = rows[0]
	} else {
		return user, domain.ErrNotFound
	}

	return user, nil
}

func (ur *userRepository) GetByID(c context.Context, id int64) (*domain.User, error) {
	var user *domain.User

	rows, err := ur.fetch(c, selectUserWithoutPass, false, id)

	if err != nil {
		return user, err
	}

	if len(rows) > 0 {
		user = rows[0]
	}

	return user, nil
}
