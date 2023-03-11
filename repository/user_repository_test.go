package repository_test

import (
	"context"
	"database/sql/driver"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/domain"
	"github.com/amitshekhariitbhu/go-backend-clean-architecture/repository"
	"github.com/stretchr/testify/assert"
)

type AnyTime struct{}

// Match satisfies sqlmock.Argument interface
func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}

func TestCreate(t *testing.T) {

	collectionName := domain.CollectionUser
	insertUser := "INSERT INTO users"

	mockUser := &domain.User{
		Name:     "Test",
		Email:    "test@gmail.com",
		Password: "password",
	}

	mockEmptyUser := &domain.User{}
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	t.Run("success", func(t *testing.T) {

		ur := repository.NewUserRepository(db, collectionName)
		// set up the expectation for the Create query
		mock.ExpectExec(insertUser).
			WithArgs(mockUser.Name, mockUser.Email, mockUser.Password, AnyTime{}).
			WillReturnResult(sqlmock.NewResult(1, 1))

		// call the Create method
		err = ur.Create(context.Background(), mockUser)
		assert.NoError(t, err)

		// ensure that all expectations were met
		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		ur := repository.NewUserRepository(db, collectionName)

		// set up the expectation for the Create query
		mock.ExpectExec(insertUser).
			WithArgs(mockEmptyUser.Name, mockEmptyUser.Email, mockEmptyUser.Password, AnyTime{}).
			WillReturnResult(sqlmock.NewResult(1, 1))

		// call the Create method
		err = ur.Create(context.Background(), mockEmptyUser)
		assert.NoError(t, err)

		// ensure that all expectations were met
		err = mock.ExpectationsWereMet()
		assert.NoError(t, err)
	})

}
