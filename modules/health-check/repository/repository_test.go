package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/go-redis/redismock/v9"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestRepository(t *testing.T) {
	tests := []struct {
		name       string
		args       context.Context
		beforeTest func(sqlmock.Sqlmock)
		want       error
		wantErr    bool
	}{
		{
			name: "Ping Postgres",
			args: context.Background(),
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectExec("SELECT 1").WillReturnError(nil)
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "Ping Postgres Fail",
			args: context.Background(),
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectExec("SELECT 1").WillReturnError(errors.New("error ping postgres"))
			},
			want:    errors.New("error ping postgres"),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.name == "Ping Postgres" {
				mockDB, mockSQL, _ := sqlmock.New()
				defer mockDB.Close()

				db := sqlx.NewDb(mockDB, "sqlmock")

				r := Repository{
					db: db,
				}

				if tt.beforeTest != nil {
					tt.beforeTest(mockSQL)
				}

				err := r.Ping(tt.args)
				if tt.wantErr {
					assert.Error(t, err)
					assert.EqualError(t, err, tt.want.Error())
				} else {
					assert.NoError(t, err)
				}
			}
		})
	}
}

func TestRepository_Redis(t *testing.T) {
	client, mock := redismock.NewClientMock()

	r := &Repository{
		redis: client,
	}

	mock.ExpectPing().SetVal("PONG")

	err := r.RedisPing(context.Background())

	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestRepository_RedisFail(t *testing.T) {
	client, mock := redismock.NewClientMock()

	r := &Repository{
		redis: client,
	}

	mock.ExpectPing().SetErr(errors.New("redis error"))

	err := r.RedisPing(context.Background())

	assert.NoError(t, mock.ExpectationsWereMet())

	assert.Error(t, err)
}
