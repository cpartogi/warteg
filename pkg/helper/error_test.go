package helper

import (
	"errors"
	"testing"

	"github.com/cpartogi/warteg/constant"
	"gotest.tools/assert"
)

func TestCommonErrorInternalServerError(t *testing.T) {
	err := errors.New("")
	result, err := CommonError(err)

	assert.Equal(t, 500, result)
}

func TestCommonErrorConflict(t *testing.T) {
	err := constant.ErrNotFound
	result, err := CommonError(err)

	assert.Equal(t, 404, result)
}

func TestCommonError(t *testing.T) {
	err := constant.ErrConflict
	result, err := CommonError(err)
	assert.Equal(t, 409, result)
}
