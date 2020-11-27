package usecase_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/vinbyte/movies/domain"
	"github.com/vinbyte/movies/domain/mocks"
	movieUsecase "github.com/vinbyte/movies/movies/usecase"
)

func TestDetailProcess(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var mockParam domain.DetailParam
		var mockDetailResponse domain.OmdbDetailResponse
		err := faker.FakeData(&mockDetailResponse)
		assert.Nil(t, err)
		mockDetailResponse.Response = "True"
		mockParam.ID = "aaaa"

		mockOmdbRepo := new(mocks.MoviesRepository)
		mockOmdbRepo.On("DetailMovie", mock.Anything, mockParam).Return("rawURL", "rawResp", mockDetailResponse, nil).Once()
		mockMysqlRepo := new(mocks.MysqlOmdbRepository)
		mockMysqlRepo.On("LogRequest", mock.Anything, "rawURL", "rawResp").Return(int64(1), nil).Once()

		u := movieUsecase.NewMovieUsecase(time.Second*3, mockOmdbRepo, mockMysqlRepo)
		resp := u.DetailProcess(context.Background(), mockParam)

		assert.Equal(t, 200, resp.Status)
		assert.NotEmpty(t, resp.Data)
	})
	t.Run("empty id", func(t *testing.T) {
		var mockParam domain.DetailParam
		var mockDetailResponse domain.OmdbDetailResponse
		err := faker.FakeData(&mockDetailResponse)
		assert.Nil(t, err)
		mockDetailResponse.Response = "True"
		mockParam.ID = ""
		mockOmdbRepo := new(mocks.MoviesRepository)
		mockMysqlRepo := new(mocks.MysqlOmdbRepository)
		u := movieUsecase.NewMovieUsecase(time.Second*3, mockOmdbRepo, mockMysqlRepo)
		resp := u.DetailProcess(context.Background(), mockParam)

		assert.Equal(t, 400, resp.Status)
		assert.Empty(t, resp.Data)
	})
	t.Run("failed", func(t *testing.T) {
		var mockParam domain.DetailParam
		var mockDetailResponse domain.OmdbDetailResponse
		err := faker.FakeData(&mockDetailResponse)
		assert.Nil(t, err)
		mockDetailResponse.Response = "False"
		mockParam.ID = "aaaa"

		mockOmdbRepo := new(mocks.MoviesRepository)
		mockOmdbRepo.On("DetailMovie", mock.Anything, mockParam).Return("rawURL", "rawResp", mockDetailResponse, nil).Once()
		mockMysqlRepo := new(mocks.MysqlOmdbRepository)
		mockMysqlRepo.On("LogRequest", mock.Anything, "rawURL", "rawResp").Return(int64(1), nil).Once()

		u := movieUsecase.NewMovieUsecase(time.Second*3, mockOmdbRepo, mockMysqlRepo)
		resp := u.DetailProcess(context.Background(), mockParam)

		assert.Equal(t, 400, resp.Status)
		assert.Empty(t, resp.Data)
	})
	t.Run("err repo", func(t *testing.T) {
		var mockParam domain.DetailParam
		var mockDetailResponse domain.OmdbDetailResponse
		err := faker.FakeData(&mockDetailResponse)
		assert.Nil(t, err)
		mockDetailResponse.Response = "True"
		mockParam.ID = "aaaa"

		mockOmdbRepo := new(mocks.MoviesRepository)
		mockOmdbRepo.On("DetailMovie", mock.Anything, mockParam).Return("rawURL", "rawResp", mockDetailResponse, errors.New("error")).Once()
		mockMysqlRepo := new(mocks.MysqlOmdbRepository)
		mockMysqlRepo.On("LogRequest", mock.Anything, "rawURL", "rawResp").Return(int64(1), errors.New("error")).Once()

		u := movieUsecase.NewMovieUsecase(time.Second*3, mockOmdbRepo, mockMysqlRepo)
		resp := u.DetailProcess(context.Background(), mockParam)

		assert.Equal(t, 200, resp.Status)
		assert.NotEmpty(t, resp.Data)
	})
}
