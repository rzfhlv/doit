package repository

import (
	"context"
	"database/sql"
	"errors"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-redis/redismock/v9"
	"github.com/jmoiron/sqlx"
	"github.com/rzfhlv/doit/modules/user/model"
	"github.com/stretchr/testify/assert"
)

var (
	now   = time.Date(2023, time.August, 15, 12, 0, 0, 0, time.UTC)
	ctx   = context.Background()
	ttl   = time.Duration(1 * time.Hour)
	key   = "testKey"
	value = "testValue"
)

func TestNewRepository(t *testing.T) {
	mockDB, _, _ := sqlmock.New()
	defer mockDB.Close()

	db := sqlx.NewDb(mockDB, "sqlmock")

	client, _ := redismock.NewClientMock()

	r := NewRepository(db, client)
	assert.NotNil(t, r)
}

func TestRegisterRepository(t *testing.T) {
	user := model.User{
		ID:        1,
		Name:      "test",
		Email:     "test@example.com",
		Username:  "test",
		Password:  "password",
		CreatedAt: now,
	}
	testCase := []struct {
		name       string
		args       context.Context
		beforeTest func(s sqlmock.Sqlmock)
		want       error
		wantErr    bool
	}{
		{
			name: "Testcase #1: Positive",
			args: context.Background(),
			beforeTest: func(s sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name", "email", "username", "password", "created_at"}).
					AddRow(user.ID, user.Name, user.Email, user.Username, user.Password, user.CreatedAt)

				s.ExpectQuery(`INSERT INTO users 
				(name, email, username, password, created_at) 
				VALUES ($1, $2, $3, $4, $5) RETURNING *`).
					WithArgs(user.Name, user.Email, user.Username, user.Password, user.CreatedAt).
					WillReturnRows(rows)
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "Testcase #2: Negative",
			args: context.Background(),
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectQuery("INSERT INTO users").
					WithArgs(user.Name, user.Email, user.Username, user.Password, user.CreatedAt).
					WillReturnError(sql.ErrNoRows)
			},
			want:    sql.ErrNoRows,
			wantErr: true,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			defer mockDB.Close()

			db := sqlx.NewDb(mockDB, "sqlmock")

			r := &Repository{
				db: db,
			}

			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}

			result, err := r.Register(context.Background(), user)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Empty(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
			}

			if err := mockSQL.ExpectationsWereMet(); err != nil {
				assert.Error(t, err)
			}
		})
	}
}

func TestLoginRepository(t *testing.T) {
	login := model.Login{
		Username: "test",
		Password: "password",
	}
	user := model.User{
		ID:        1,
		Name:      "test",
		Email:     "test@example.com",
		Username:  "test",
		Password:  "password",
		CreatedAt: now,
	}
	testCase := []struct {
		name       string
		args       context.Context
		beforeTest func(s sqlmock.Sqlmock)
		want       error
		wantErr    bool
	}{
		{
			name: "Testcase #1: Positive",
			args: context.Background(),
			beforeTest: func(s sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name", "email", "username", "password", "created_at"}).
					AddRow(user.ID, user.Name, user.Email, user.Username, user.Password, user.CreatedAt)

				s.ExpectQuery("SELECT * FROM users WHERE username = $1").
					WithArgs(login.Username).WillReturnRows(rows)
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "Testcase #2: Negative",
			args: context.Background(),
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectQuery("SELECT * FROM users WHERE username = $1").
					WithArgs(login.Username).WillReturnError(sql.ErrNoRows)
			},
			want:    sql.ErrNoRows,
			wantErr: true,
		},
	}
	for _, tt := range testCase {
		t.Run(tt.name, func(t *testing.T) {
			mockDB, mockSQL, _ := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			defer mockDB.Close()

			db := sqlx.NewDb(mockDB, "sqlmock")

			r := &Repository{
				db: db,
			}

			if tt.beforeTest != nil {
				tt.beforeTest(mockSQL)
			}

			result, err := r.Login(context.Background(), login)
			if tt.wantErr {
				assert.Error(t, err)
				assert.Empty(t, result)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, result)
			}

			if err := mockSQL.ExpectationsWereMet(); err != nil {
				assert.Error(t, err)
			}
		})
	}
}

func TestSetRepository(t *testing.T) {
	client, mock := redismock.NewClientMock()
	r := &Repository{
		redis: client,
	}

	t.Run("Testcase #1: Positive", func(t *testing.T) {
		mock.ExpectSet(key, value, ttl).SetVal("OK")
		err := r.Set(ctx, key, value, ttl)
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {
		mock.ExpectSet(key, value, ttl).SetErr(errors.New("error"))
		err := r.Set(ctx, key, value, ttl)
		assert.Error(t, err)
	})
}

func TestGetRepository(t *testing.T) {
	client, mock := redismock.NewClientMock()
	r := &Repository{
		redis: client,
	}

	t.Run("Testcase #1: Positive", func(t *testing.T) {
		mock.ExpectGet(key).SetVal(value)
		result, err := r.Get(ctx, key)
		assert.NoError(t, err)
		assert.Equal(t, value, result)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {
		mock.ExpectGet(key).SetErr(errors.New("error"))
		_, err := r.Get(ctx, key)
		assert.Error(t, err)
	})
}

func TestDelRepository(t *testing.T) {
	client, mock := redismock.NewClientMock()
	r := &Repository{
		redis: client,
	}

	t.Run("Testcase #1: Positive", func(t *testing.T) {
		mock.ExpectDel(key).SetVal(int64(1))
		err := r.Del(ctx, key)
		assert.NoError(t, err)
	})

	t.Run("Testcase #2: Negative", func(t *testing.T) {
		mock.ExpectDel(key).SetErr(errors.New("error"))
		err := r.Del(ctx, key)
		assert.Error(t, err)
	})
}
