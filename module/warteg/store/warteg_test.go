package store

import (
	"context"
	"testing"

	"errors"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/cpartogi/warteg/schema/request"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestWartegAddQuery(t *testing.T) {
	arg := request.Warteg{
		WartegName:        "a",
		WartegDesc:        "b",
		WartegAddr:        "c",
		WartegContactName: "d",
		WartegPhone:       "e",
	}

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectExec(addWarteg).WithArgs(arg.WartegName, arg.WartegDesc, arg.WartegAddr, arg.WartegContactName, arg.WartegPhone).WillReturnResult(sqlmock.NewResult(1, 1))

	a := NewStore(db)
	warteg, err := a.WartegAdd(context.TODO(), arg)

	assert.NoError(t, err)
	require.NotEmpty(t, warteg)
}

func TestWategAddQueryError(t *testing.T) {
	arg := request.Warteg{
		WartegName:        "a",
		WartegDesc:        "b",
		WartegAddr:        "c",
		WartegContactName: "d",
		WartegPhone:       "e",
	}

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectExec(addWarteg).WithArgs(arg.WartegName, arg.WartegDesc, arg.WartegAddr, arg.WartegContactName, arg.WartegPhone).WillReturnError(errors.New("error connect to db"))

	a := NewStore(db)
	_, err = a.WartegAdd(context.TODO(), arg)

	assert.Error(t, err)
}

func TestWartegAddQueryNoRows(t *testing.T) {
	arg := request.Warteg{
		WartegName:        "a",
		WartegDesc:        "b",
		WartegAddr:        "c",
		WartegContactName: "d",
		WartegPhone:       "e",
	}

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectExec(addWarteg).WithArgs(arg.WartegName, arg.WartegDesc, arg.WartegAddr, arg.WartegContactName, arg.WartegPhone).WillReturnResult(sqlmock.NewResult(0, 0))

	a := NewStore(db)
	_, err = a.WartegAdd(context.TODO(), arg)

	assert.NoError(t, err)
	assert.Equal(t, "", "")
}

func TestWartegDelete(t *testing.T) {
	warteg_id := "abcdef"

	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	mock.ExpectExec(deleteWarteg).WithArgs(warteg_id).WillReturnResult((sqlmock.NewResult(1, 1)))

	a := NewStore(db)

	_, err = a.WartegDelete(context.TODO(), warteg_id)

	assert.NoError(t, err)
}
