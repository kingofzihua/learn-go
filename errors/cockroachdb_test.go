package errors

import (
	"database/sql"
	"fmt"
	cerrors "github.com/cockroachdb/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func getUser() error {
	return cerrors.New("fakedb: no database name")
}

func getUserNoRow() error {
	return sql.ErrNoRows
}

func stack1() error {
	return cerrors.Wrap(getUser(), "user is nil1")
}

func stack2() error {
	return cerrors.Wrap(stack1(), "user is nil2")
}

func TestStack(t *testing.T) {
	fmt.Printf("error %+v\n", stack2())
}

func TestSafeDetail(t *testing.T) {
	var err = getUserNoRow()

	err = cerrors.WithSafeDetails(err, "user not found")

	err = cerrors.WithSafeDetails(err, "user not found %s", "2")

	assert.True(t, cerrors.Is(err, sql.ErrNoRows))

	err = cerrors.WithMessage(err, "user is nil")

	assert.True(t, cerrors.Is(err, sql.ErrNoRows))

	fmt.Printf("error : [%s]\n", err)
	fmt.Printf("\n\nerror : %+v\n", err)
	fmt.Printf("\n\nGetSafeDetails : %+v\n", cerrors.GetSafeDetails(err).SafeDetails)
	fmt.Printf("\n\nGetAllSafeDetails : %+v\n", cerrors.GetAllSafeDetails(err))
}
