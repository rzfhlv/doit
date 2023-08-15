package repository

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
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
			want: nil,
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

				got := r.Ping(tt.args)
				if got != tt.want {
					t.Errorf("Ping = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
