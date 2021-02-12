package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/cpartogi/warteg/module/warteg/mocks"
	"github.com/cpartogi/warteg/schema/response"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var errorWarteg = errors.New("error warteg")

func TestWartegHandlerNewWartegHandler(t *testing.T) {
	e := echo.New()
	mockWarteg := new(mocks.Usecase)
	NewWartegHandler(e, mockWarteg)
}

func TestWartegAdd(t *testing.T) {
	type input struct {
		req map[string]interface{}
	}

	type output struct {
		err        error
		statusCode int
	}

	cases := []struct {
		name           string
		expectedInput  input
		expectedOutput output
		configureMock  func(
			payload input,
			mockWarteg *mocks.Usecase,
		)
	}{
		{
			name: "#1 success insert data",
			expectedInput: input{
				req: map[string]interface{}{
					"warteg_addr":         "a",
					"warteg_contact_name": "b",
					"warteg_desc":         "c",
					"warteg_name":         "d",
					"warteg_phone":        "e",
				},
			},
			expectedOutput: output{nil, http.StatusCreated},
			configureMock: func(
				payload input,
				mockWarteg *mocks.Usecase,
			) {
				wResponse := response.WartegAdd{}

				mockWarteg.
					On("WartegAdd", mock.Anything, mock.Anything, mock.Anything).
					Return(wResponse, nil)
			},
		},
		{
			name: "#2 unprocessable add warteg",
			expectedInput: input{
				req: map[string]interface{}{
					"warteg_addr":         "a",
					"warteg_contact_name": "b",
					"warteg_desc":         "c",
					"warteg_name":         1,
					"warteg_phone":        "e",
				},
			},
			expectedOutput: output{nil, http.StatusUnprocessableEntity},
			configureMock: func(
				payload input,
				mockWarteg *mocks.Usecase,
			) {
				wResponse := response.WartegAdd{}

				mockWarteg.
					On("WartegAdd", mock.Anything, mock.Anything).
					Return(wResponse, nil)
			},
		},
		{
			name: "#3 bad request add warteg",
			expectedInput: input{
				req: map[string]interface{}{
					"warteg_addr":         "a",
					"warteg_contact_name": "b",
				},
			},
			expectedOutput: output{nil, http.StatusBadRequest},
			configureMock: func(
				payload input,
				mockWarteg *mocks.Usecase,
			) {
				wResponse := response.WartegAdd{}

				mockWarteg.
					On("WartegAdd", mock.Anything, mock.Anything).
					Return(wResponse, nil)
			},
		},
		{
			name: "#4 internal server error add warteg",
			expectedInput: input{
				req: map[string]interface{}{
					"warteg_addr":         "a",
					"warteg_contact_name": "b",
					"warteg_desc":         "c",
					"warteg_name":         "d",
					"warteg_phone":        "e",
				},
			},
			expectedOutput: output{nil, http.StatusInternalServerError},
			configureMock: func(
				payload input,
				mockWarteg *mocks.Usecase,
			) {
				wResponse := response.WartegAdd{}

				mockWarteg.
					On("WartegAdd", mock.Anything, mock.Anything).
					Return(wResponse, errorWarteg)
			},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			mockWarteg := new(mocks.Usecase)

			payload, err := json.Marshal(testCase.expectedInput.req)

			assert.NoError(t, err)

			e := echo.New()

			req, err := http.NewRequest(echo.POST, "/v1/warteg",
				strings.NewReader(string(payload)))

			assert.NoError(t, err)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/v1/warteg")

			testCase.configureMock(
				testCase.expectedInput,
				mockWarteg,
			)

			handler := WartegHandler{
				wartegUsecase: mockWarteg,
			}

			err = handler.WartegAdd(c)
			assert.Equal(t, testCase.expectedOutput.err, err)
			assert.Equal(t, testCase.expectedOutput.statusCode, rec.Code)

		})
	}
}

func TestWartegDelete(t *testing.T) {
	type input struct {
		warteg_id string
	}

	type output struct {
		err        error
		statusCode int
	}

	cases := []struct {
		name           string
		expectedInput  input
		expectedOutput output
		configureMock  func(
			payload input,
			mockWarteg *mocks.Usecase,
		)
	}{
		{
			name: "#1 success delete warteg",
			expectedInput: input{
				warteg_id: "abc",
			},
			expectedOutput: output{nil, http.StatusOK},
			configureMock: func(
				payload input,
				mockWarteg *mocks.Usecase,
			) {
				wResponse := response.WartegDelete{}

				mockWarteg.
					On("WartegDelete", mock.Anything, mock.Anything).
					Return(wResponse, nil)
			},
		},
		{
			name: "#2 internal server error delete warteg",
			expectedInput: input{
				warteg_id: "abc",
			},
			expectedOutput: output{nil, http.StatusInternalServerError},
			configureMock: func(
				payload input,
				mockWarteg *mocks.Usecase,
			) {
				wResponse := response.WartegDelete{}

				mockWarteg.
					On("WartegDelete", mock.Anything, mock.Anything).
					Return(wResponse, errorWarteg)
			},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			mockWarteg := new(mocks.Usecase)

			warteg_id := testCase.expectedInput.warteg_id

			e := echo.New()

			req, err := http.NewRequest(echo.DELETE, "/v1/warteg/:warteg_id",
				strings.NewReader(string(warteg_id)))

			assert.NoError(t, err)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/v1/warteg/")

			testCase.configureMock(
				testCase.expectedInput,
				mockWarteg,
			)

			handler := WartegHandler{
				wartegUsecase: mockWarteg,
			}

			err = handler.WartegDelete(c)
			assert.Equal(t, testCase.expectedOutput.err, err)
			assert.Equal(t, testCase.expectedOutput.statusCode, rec.Code)

		})
	}
}

func TestWartegUpdate(t *testing.T) {
	type input struct {
		warteg_id string
		req       map[string]interface{}
	}

	type output struct {
		err        error
		statusCode int
	}

	cases := []struct {
		name           string
		expectedInput  input
		expectedOutput output
		configureMock  func(
			payload input,
			mockWarteg *mocks.Usecase,
		)
	}{
		{
			name: "#1 success update warteg",
			expectedInput: input{
				warteg_id: "abc",
				req: map[string]interface{}{
					"warteg_addr":         "a",
					"warteg_contact_name": "b",
					"warteg_desc":         "c",
					"warteg_name":         "d",
					"warteg_phone":        "e",
				},
			},
			expectedOutput: output{nil, http.StatusOK},
			configureMock: func(
				payload input,
				mockWarteg *mocks.Usecase,
			) {
				wResponse := response.WartegUpdate{}

				mockWarteg.
					On("WartegUpdate", mock.Anything, mock.Anything).
					Return(wResponse, nil)
			},
		},
		{
			name: "#2 bad request update",
			expectedInput: input{
				warteg_id: "abc",
				req: map[string]interface{}{
					"warteg_addr":         "a",
					"warteg_contact_name": "b",
					"warteg_desc":         "c",
				},
			},
			expectedOutput: output{nil, http.StatusBadRequest},
			configureMock: func(
				payload input,
				mockWarteg *mocks.Usecase,
			) {
				wResponse := response.WartegUpdate{}

				mockWarteg.
					On("WartegUpdate", mock.Anything, mock.Anything).
					Return(wResponse, nil)
			},
		},
		{
			name: "#3 unprocessable update",
			expectedInput: input{
				warteg_id: "abc",
				req: map[string]interface{}{
					"warteg_addr":         "a",
					"warteg_contact_name": "b",
					"warteg_desc":         "c",
					"warteg_name":         "d",
					"warteg_phone":        1,
				},
			},
			expectedOutput: output{nil, http.StatusUnprocessableEntity},
			configureMock: func(
				payload input,
				mockWarteg *mocks.Usecase,
			) {
				wResponse := response.WartegUpdate{}

				mockWarteg.
					On("WartegUpdate", mock.Anything, mock.Anything).
					Return(wResponse, nil)
			},
		},
		{
			name: "#4 internal server error update",
			expectedInput: input{
				warteg_id: "abc",
				req: map[string]interface{}{
					"warteg_addr":         "a",
					"warteg_contact_name": "b",
					"warteg_desc":         "c",
					"warteg_name":         "d",
					"warteg_phone":        "e",
				},
			},
			expectedOutput: output{nil, http.StatusInternalServerError},
			configureMock: func(
				payload input,
				mockWarteg *mocks.Usecase,
			) {
				wResponse := response.WartegUpdate{}

				mockWarteg.
					On("WartegUpdate", mock.Anything, mock.Anything).
					Return(wResponse, errorWarteg)
			},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			mockWarteg := new(mocks.Usecase)

			j, err := json.Marshal(testCase.expectedInput.req)

			e := echo.New()

			req, err := http.NewRequest(echo.PUT, "/v1/warteg/:warteg_id",
				strings.NewReader(string(j)))

			assert.NoError(t, err)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/v1/warteg/")

			testCase.configureMock(
				testCase.expectedInput,
				mockWarteg,
			)

			handler := WartegHandler{
				wartegUsecase: mockWarteg,
			}

			err = handler.WartegUpdate(c)
			assert.Equal(t, testCase.expectedOutput.err, err)
			assert.Equal(t, testCase.expectedOutput.statusCode, rec.Code)

		})
	}
}

func TestWartegList(t *testing.T) {
	type input struct {
		warteg_name string
	}

	type output struct {
		err        error
		statusCode int
	}

	cases := []struct {
		name           string
		expectedInput  input
		expectedOutput output
		configureMock  func(
			payload input,
			mockWarteg *mocks.Usecase,
		)
	}{
		{
			name: "#1 success get warteg list",
			expectedInput: input{
				warteg_name: "asdfsdfsd",
			},
			expectedOutput: output{nil, http.StatusOK},
			configureMock: func(
				payload input,
				mockWarteg *mocks.Usecase,
			) {
				wResponse := []response.WartegList{}

				mockWarteg.
					On("WartegList", mock.Anything, mock.Anything).
					Return(wResponse, nil)
			},
		},
		{
			name: "#2 internal server error ",
			expectedInput: input{
				warteg_name: "asdfsdfsd",
			},
			expectedOutput: output{nil, http.StatusInternalServerError},
			configureMock: func(
				payload input,
				mockWarteg *mocks.Usecase,
			) {
				wResponse := []response.WartegList{}

				mockWarteg.
					On("WartegList", mock.Anything, mock.Anything).
					Return(wResponse, errorWarteg)
			},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			mockWarteg := new(mocks.Usecase)

			warteg_name := testCase.expectedInput.warteg_name

			e := echo.New()

			req, err := http.NewRequest(echo.GET, "/v1/wartegs/list",
				strings.NewReader(string(warteg_name)))

			assert.NoError(t, err)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/v1/wartegs/list")

			testCase.configureMock(
				testCase.expectedInput,
				mockWarteg,
			)

			handler := WartegHandler{
				wartegUsecase: mockWarteg,
			}

			err = handler.WartegList(c)
			assert.Equal(t, testCase.expectedOutput.err, err)
			assert.Equal(t, testCase.expectedOutput.statusCode, rec.Code)

		})
	}
}

func TestWartegDetail(t *testing.T) {
	type input struct {
		warteg_id string
	}

	type output struct {
		err        error
		statusCode int
	}

	cases := []struct {
		name           string
		expectedInput  input
		expectedOutput output
		configureMock  func(
			payload input,
			mockWarteg *mocks.Usecase,
		)
	}{
		{
			name: "#1 success get data",
			expectedInput: input{
				warteg_id: "abc",
			},
			expectedOutput: output{nil, http.StatusOK},
			configureMock: func(
				payload input,
				mockWarteg *mocks.Usecase,
			) {
				wResponse := response.WartegDetail{}

				mockWarteg.
					On("WartegDetail", mock.Anything, mock.Anything).
					Return(wResponse, nil)
			},
		},
		{
			name: "#2 internal server error update",
			expectedInput: input{
				warteg_id: "abc",
			},
			expectedOutput: output{nil, http.StatusInternalServerError},
			configureMock: func(
				payload input,
				mockWarteg *mocks.Usecase,
			) {
				wResponse := response.WartegDetail{}

				mockWarteg.
					On("WartegDetail", mock.Anything, mock.Anything).
					Return(wResponse, errorWarteg)
			},
		},
	}

	for _, testCase := range cases {
		t.Run(testCase.name, func(t *testing.T) {
			mockWarteg := new(mocks.Usecase)

			warteg_id := testCase.expectedInput.warteg_id

			e := echo.New()

			req, err := http.NewRequest(echo.GET, "/v1/warteg/:warteg_id",
				strings.NewReader(string(warteg_id)))

			assert.NoError(t, err)
			req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			c.SetPath("/v1/warteg/")

			testCase.configureMock(
				testCase.expectedInput,
				mockWarteg,
			)

			handler := WartegHandler{
				wartegUsecase: mockWarteg,
			}

			err = handler.WartegDetail(c)
			assert.Equal(t, testCase.expectedOutput.err, err)
			assert.Equal(t, testCase.expectedOutput.statusCode, rec.Code)

		})
	}
}
