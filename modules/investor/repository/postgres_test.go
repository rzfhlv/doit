package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/rzfhlv/doit/modules/investor/model"
	"github.com/rzfhlv/doit/utilities/param"
	"github.com/stretchr/testify/assert"
)

type testCase struct {
	name       string
	args       context.Context
	beforeTest func(s sqlmock.Sqlmock)
	want       error
	wantError  bool
}

var (
	ctx       = context.Background()
	investors = []model.Investor{
		{
			ID: 1, Name: "Test 1",
		},
		{
			ID: 2, Name: "Test 2",
		},
	}
	errFoo    = errors.New("foo")
	paramTest = param.Param{
		Limit: 10,
		Page:  1,
	}
)

func TestGetPsqlRepository(t *testing.T) {
	testCase := []testCase{
		{
			name: "Testcase #1: Positive",
			args: ctx,
			beforeTest: func(s sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name"}).
					AddRow(investors[0].ID, investors[0].Name).
					AddRow(investors[1].ID, investors[1].Name)

				s.ExpectQuery("SELECT * FROM investors").WillReturnRows(rows)
			},
			want:      nil,
			wantError: false,
		},
		{
			name: "Testcase #2: Negative",
			args: ctx,
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectQuery("SELECT * FROM investors").WillReturnError(errFoo)
			},
			want:      errFoo,
			wantError: true,
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

			result, err := r.GetPsql(ctx)
			if tt.wantError {
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

func TestGetAllRepository(t *testing.T) {
	testCase := []testCase{
		{
			name: "Testcase #1: Positive",
			args: ctx,
			beforeTest: func(s sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name"}).
					AddRow(investors[0].ID, investors[0].Name).
					AddRow(investors[1].ID, investors[1].Name)

				s.ExpectQuery("SELECT * FROM investors ORDER BY investors.id DESC LIMIT $1 OFFSET $2;").
					WithArgs(paramTest.Limit, paramTest.CalculateOffset()).
					WillReturnRows(rows)
			},
			want:      nil,
			wantError: false,
		},
		{
			name: "Testcase #2: Negative",
			args: ctx,
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectQuery("SELECT * FROM investors ORDER BY investors.id DESC LIMIT $1 OFFSET $2;").
					WillReturnError(errFoo)
			},
			want:      errFoo,
			wantError: true,
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

			result, err := r.GetAll(ctx, paramTest)
			if tt.wantError {
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

func TestGetByIDRepository(t *testing.T) {
	testCase := []testCase{
		{
			name: "Testcase #1: Positive",
			args: ctx,
			beforeTest: func(s sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"id", "name"}).
					AddRow(investors[0].ID, investors[0].Name)

				s.ExpectQuery("SELECT * FROM investors WHERE id = $1;").
					WithArgs(investors[0].ID).
					WillReturnRows(rows)
			},
			want:      nil,
			wantError: false,
		},
		{
			name: "Testcase #2: Negative",
			args: ctx,
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectQuery("SELECT * FROM investors WHERE id = $1;").
					WithArgs(investors[0].ID).
					WillReturnError(errFoo)
			},
			want:      errFoo,
			wantError: true,
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

			result, err := r.GetByID(ctx, investors[0].ID)
			if tt.wantError {
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

func TestCountRepository(t *testing.T) {
	expectedCount := int64(10)
	testCase := []testCase{
		{
			name: "Testcase #1: Positive",
			args: ctx,
			beforeTest: func(s sqlmock.Sqlmock) {
				rows := sqlmock.NewRows([]string{"count"}).
					AddRow(expectedCount)

				s.ExpectQuery("SELECT count(*) FROM investors;").
					WillReturnRows(rows)
			},
			want:      nil,
			wantError: false,
		},
		{
			name: "Testcase #2: Negative",
			args: ctx,
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectQuery("SELECT count(*) FROM investors;").
					WillReturnError(errFoo)
			},
			want:      errFoo,
			wantError: true,
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

			_, err := r.Count(ctx)
			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			if err := mockSQL.ExpectationsWereMet(); err != nil {
				assert.Error(t, err)
			}
		})
	}
}

func TestGenerateRepository(t *testing.T) {
	testCase := []testCase{
		{
			name: "Testcase #1: Positive",
			args: ctx,
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectExec("INSERT INTO investors (name) VALUES ($1) RETURNING id;").
					WithArgs(investors[0].Name).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
			want:      nil,
			wantError: false,
		},
		{
			name: "Testcase #2: Negative",
			args: ctx,
			beforeTest: func(s sqlmock.Sqlmock) {
				s.ExpectExec("INSERT INTO investors (name) VALUES ($1) RETURNING id;").
					WithArgs(investors[0].Name).
					WillReturnError(errFoo)
			},
			want:      errFoo,
			wantError: true,
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

			err := r.Generate(ctx, investors[0].Name)
			if tt.wantError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}

			if err := mockSQL.ExpectationsWereMet(); err != nil {
				assert.Error(t, err)
			}
		})
	}
}
