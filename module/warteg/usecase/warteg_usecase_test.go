package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/cpartogi/warteg/module/warteg/mocks"
	"github.com/cpartogi/warteg/schema/request"
	"github.com/cpartogi/warteg/schema/response"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewWartegUsecase(t *testing.T) {
	mockWarteg := new(mocks.Usecase)
	NewWartegUsecase(mockWarteg, 5)
}

func TestWartegAdd(t *testing.T) {
	resWartegAdd := response.WartegAdd{}
	mockRepo := new(mocks.Repository)

	mockWartegAdd := request.Warteg{
		WartegName:        "nama warteg",
		WartegDesc:        "deskripsi warteg",
		WartegAddr:        "alamat warteg",
		WartegContactName: "nama kontak warteg",
		WartegPhone:       "09223422",
	}

	t.Run("Success", func(t *testing.T) {
		tempWartegAdd := mockWartegAdd
		u := NewWartegUsecase(mockRepo, 2)

		mockRepo.On("WartegAdd", mock.Anything, mock.AnythingOfType("request.Warteg")).
			Return(resWartegAdd, nil).Once()

		_, err := u.WartegAdd(context.TODO(), tempWartegAdd)
		assert.NoError(t, err)
	})

	t.Run("failed", func(t *testing.T) {
		tempWartegAdd := mockWartegAdd
		u := NewWartegUsecase(mockRepo, 2)

		mockRepo.On("WartegAdd", mock.Anything, mock.AnythingOfType("request.Warteg")).
			Return(resWartegAdd, errors.New("apa deh")).Once()

		_, err := u.WartegAdd(context.TODO(), tempWartegAdd)
		assert.Error(t, err)
	})

}

func TestWartegDelete(t *testing.T) {
	resWartegDelete := response.WartegDelete{}
	mockRepo := new(mocks.Repository)

	mockWartegDelete := "234234"

	t.Run("success", func(t *testing.T) {
		tempWartegDelete := mockWartegDelete
		u := NewWartegUsecase(mockRepo, 2)

		mockRepo.On("WartegDelete", mock.Anything, "234234").
			Return(resWartegDelete, nil).Once()

		_, err := u.WartegDelete(context.TODO(), tempWartegDelete)
		assert.NoError(t, err)
	})

	t.Run("failed", func(t *testing.T) {
		tempWartegDelete := mockWartegDelete
		u := NewWartegUsecase(mockRepo, 2)

		mockRepo.On("WartegDelete", mock.Anything, "234234").
			Return(resWartegDelete, errors.New("apa deh")).Once()

		_, err := u.WartegDelete(context.TODO(), tempWartegDelete)
		assert.Error(t, err)
	})
}

func TestWartegUpdate(t *testing.T) {
	resWartegUpdate := response.WartegUpdate{}
	mockRepo := new(mocks.Repository)

	mockWartegId := "abcdef"
	mockWartegUpdate := request.WartegUpdate{
		WartegName:        "abc",
		WartegDesc:        "def",
		WartegAddr:        "ghi",
		WartegContactName: "jkl",
		WartegPhone:       "23423",
	}

	t.Run("success", func(t *testing.T) {
		tempWartegUpdate := mockWartegUpdate
		u := NewWartegUsecase(mockRepo, 2)

		mockRepo.On("WartegUpdate", mock.Anything, "abcdef", mock.AnythingOfType("request.WartegUpdate")).
			Return(resWartegUpdate, nil).Once()

		_, err := u.WartegUpdate(context.TODO(), mockWartegId, tempWartegUpdate)
		assert.NoError(t, err)
	})

	t.Run("failed", func(t *testing.T) {
		tempWartegUpdate := mockWartegUpdate
		u := NewWartegUsecase(mockRepo, 2)

		mockRepo.On("WartegUpdate", mock.Anything, "abcdef", mock.AnythingOfType("request.WartegUpdate")).
			Return(resWartegUpdate, errors.New("apa deh")).Once()

		_, err := u.WartegUpdate(context.TODO(), mockWartegId, tempWartegUpdate)
		assert.Error(t, err)
	})

}

func TestWartegList(t *testing.T) {
	resWartegList := []response.WartegList{}
	mockRepo := new(mocks.Repository)

	mockWartegName := "abcd"

	t.Run("success", func(t *testing.T) {
		tempWartegList := mockWartegName
		u := NewWartegUsecase(mockRepo, 2)

		mockRepo.On("WartegList", mock.Anything, "abcd").
			Return(resWartegList, nil).Once()

		_, err := u.WartegList(context.TODO(), mockWartegName)
		assert.NoError(t, err)
		assert.Equal(t, mockWartegName, tempWartegList)
	})

	t.Run("failed", func(t *testing.T) {
		tempWartegList := mockWartegName
		u := NewWartegUsecase(mockRepo, 2)

		mockRepo.On("WartegList", mock.Anything, "abcd").
			Return(resWartegList, errors.New("apa deh")).Once()

		_, err := u.WartegList(context.TODO(), mockWartegName)
		assert.Error(t, err)
		assert.Equal(t, mockWartegName, tempWartegList)
	})
}

func TestWartegDetail(t *testing.T) {
	resWartegDetail := response.WartegDetail{}
	mockRepo := new(mocks.Repository)

	mockWartegId := "abcd-efgh"

	t.Run("success", func(t *testing.T) {
		tempWartegDetail := mockWartegId
		u := NewWartegUsecase(mockRepo, 2)

		mockRepo.On("WartegDetail", mock.Anything, "abcd-efgh").
			Return(resWartegDetail, nil).Once()

		_, err := u.WartegDetail(context.TODO(), mockWartegId)
		assert.NoError(t, err)
		assert.Equal(t, mockWartegId, tempWartegDetail)
	})

	t.Run("failed", func(t *testing.T) {
		tempWartegDetail := mockWartegId
		u := NewWartegUsecase(mockRepo, 2)

		mockRepo.On("WartegDetail", mock.Anything, "abcd-efgh").
			Return(resWartegDetail, errors.New("apa deh")).Once()

		_, err := u.WartegDetail(context.TODO(), mockWartegId)
		assert.Error(t, err)
		assert.Equal(t, mockWartegId, tempWartegDetail)
	})
}
