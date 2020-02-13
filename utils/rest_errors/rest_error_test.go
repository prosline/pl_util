package rest_errors

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func TestNewInternalServerError(t *testing.T) {
	err := NewInternalServerError("This is an Error Message.",errors.New("Database Error"))
	assert.NotNil(t, err)
	assert.EqualValues(t,http.StatusInternalServerError, err.Status())
	assert.EqualValues(t,"This is an Error Message.", err.Message())
	assert.EqualValues(t,"message: This is an Error Message. - status: 500 - error: internal_server_error - causes: [Database Error]", err.Error())

	assert.NotNil(t,err.Causes)
	assert.EqualValues(t,1, len(err.Causes()))
	assert.EqualValues(t,"Database Error",err.Causes()[0])
}
func TestNewBadRequestError(t *testing.T) {
	err := NewBadRequestError("Bad Request Error")
	assert.NotNil(t, err)
	assert.EqualValues(t,http.StatusBadRequest,err.Status())
	assert.EqualValues(t,"Bad Request Error",err.Message())
	assert.EqualValues(t,"message: Bad Request Error - status: 400 - error: bad_request - causes: []",err.Error())
}
func TestNewError(t *testing.T) {
	err := NewError("New Error")
	assert.NotNil(t, err)
	assert.EqualValues(t, "New Error", err.Error())
}
func TestNotFoundError(t *testing.T) {
	err := NewNotFoundError("Not Found Error")
	assert.NotNil(t, err)
	assert.EqualValues(t,"message: Not Found Error - status: 500 - error: not_found - causes: []", err.Error() )
	assert.EqualValues(t,http.StatusInternalServerError, err.Status() )
	assert.EqualValues(t,"Not Found Error", err.Message() )
}

