package repository

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/rzfhlv/doit/modules/user/model"
	"github.com/stretchr/testify/assert"
)

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
