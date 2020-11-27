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

func TestSearchProcess(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var mockParam domain.SearchParam
		var mockRawParam domain.SearchParam
		var mockSearchResponse domain.OmdbSearchResponse
		err := faker.FakeData(&mockSearchResponse)
		assert.Nil(t, err)
		mockSearchResponse.Response = "True"
		mockParam.PageStr = "1"
		mockParam.Query = "Batman"
		mockParam.Page = 1
		mockRawParam.Query = "Batman"

		mockOmdbRepo := new(mocks.MoviesRepository)
		mockOmdbRepo.On("SearchMovies", mock.Anything, mockParam).Return("rawURL", "rawResp", mockSearchResponse, nil).Once()
		mockMysqlRepo := new(mocks.MysqlOmdbRepository)
		mockMysqlRepo.On("LogRequest", mock.Anything, "rawURL", "rawResp").Return(int64(1), nil).Once()

		u := movieUsecase.NewMovieUsecase(time.Second*3, mockOmdbRepo, mockMysqlRepo)
		resp := u.SearchProcess(context.Background(), mockRawParam)

		assert.Equal(t, 200, resp.Status)
		assert.NotEmpty(t, resp.Data)
	})
	t.Run("err convert", func(t *testing.T) {
		var mockRawParam domain.SearchParam
		mockOmdbRepo := new(mocks.MoviesRepository)
		mockMysqlRepo := new(mocks.MysqlOmdbRepository)
		mockRawParam.Query = "Batman"
		mockRawParam.PageStr = "aaa"
		u := movieUsecase.NewMovieUsecase(time.Second*3, mockOmdbRepo, mockMysqlRepo)
		resp := u.SearchProcess(context.Background(), mockRawParam)

		assert.Equal(t, 400, resp.Status)
		assert.Empty(t, resp.Data)
	})
	t.Run("err repo", func(t *testing.T) {
		var mockParam domain.SearchParam
		var mockRawParam domain.SearchParam
		var mockSearchResponse domain.OmdbSearchResponse
		err := faker.FakeData(&mockSearchResponse)
		assert.Nil(t, err)
		mockSearchResponse.Response = "True"
		mockParam.PageStr = "1"
		mockParam.Query = "Batman"
		mockParam.Page = 1
		mockRawParam.Query = "Batman"

		mockOmdbRepo := new(mocks.MoviesRepository)
		mockOmdbRepo.On("SearchMovies", mock.Anything, mockParam).Return("rawURL", "rawResp", mockSearchResponse, errors.New("something wrong")).Once()
		mockMysqlRepo := new(mocks.MysqlOmdbRepository)
		mockMysqlRepo.On("LogRequest", mock.Anything, "rawURL", "rawResp").Return(int64(1), errors.New("something error")).Once()

		u := movieUsecase.NewMovieUsecase(time.Second*3, mockOmdbRepo, mockMysqlRepo)
		resp := u.SearchProcess(context.Background(), mockRawParam)

		assert.Equal(t, 200, resp.Status)
		assert.NotEmpty(t, resp.Data)
	})
	t.Run("empty query", func(t *testing.T) {
		// var mockResponse domain.BaseResponse
		var mockParam domain.SearchParam
		var mockRawParam domain.SearchParam
		var mockSearchResponse domain.OmdbSearchResponse
		err := faker.FakeData(&mockSearchResponse)
		assert.Nil(t, err)
		mockSearchResponse.Response = "False"
		mockParam.PageStr = "1"
		mockParam.Query = ""
		mockParam.Page = 1
		mockRawParam.Query = ""

		mockOmdbRepo := new(mocks.MoviesRepository)
		mockOmdbRepo.On("SearchMovies", mock.Anything, mockParam).Return("rawURL", "rawResp", mockSearchResponse, nil).Once()
		mockMysqlRepo := new(mocks.MysqlOmdbRepository)
		mockMysqlRepo.On("LogRequest", mock.Anything, "rawURL", "rawResp").Return(int64(1), nil).Once()

		u := movieUsecase.NewMovieUsecase(time.Second*3, mockOmdbRepo, mockMysqlRepo)
		resp := u.SearchProcess(context.Background(), mockRawParam)

		assert.Equal(t, 400, resp.Status)
		assert.NotEmpty(t, resp.Data)
	})
	t.Run("page1", func(t *testing.T) {
		// var mockResponse domain.BaseResponse
		var mockParam domain.SearchParam
		var mockRawParam domain.SearchParam
		var mockSearchResponse domain.OmdbSearchResponse
		err := faker.FakeData(&mockSearchResponse)
		assert.Nil(t, err)
		mockSearchResponse.Response = "True"
		mockSearchResponse.TotalResults = "39"
		mockParam.PageStr = "1"
		mockParam.Query = "Batman"
		mockParam.Page = 1
		mockRawParam.Query = "Batman"

		mockOmdbRepo := new(mocks.MoviesRepository)
		mockOmdbRepo.On("SearchMovies", mock.Anything, mockParam).Return("rawURL", "rawResp", mockSearchResponse, nil).Once()
		mockMysqlRepo := new(mocks.MysqlOmdbRepository)
		mockMysqlRepo.On("LogRequest", mock.Anything, "rawURL", "rawResp").Return(int64(1), nil).Once()

		u := movieUsecase.NewMovieUsecase(time.Second*3, mockOmdbRepo, mockMysqlRepo)
		resp := u.SearchProcess(context.Background(), mockRawParam)

		assert.Equal(t, 200, resp.Status)
		assert.NotEmpty(t, resp.Data)
	})
	t.Run("page2", func(t *testing.T) {
		var mockParam domain.SearchParam
		var mockRawParam domain.SearchParam
		var mockSearchResponse domain.OmdbSearchResponse
		err := faker.FakeData(&mockSearchResponse)
		assert.Nil(t, err)
		mockSearchResponse.Response = "True"
		mockSearchResponse.TotalResults = "40"
		mockParam.PageStr = "1"
		mockParam.Query = "Batman"
		mockParam.Page = 1
		mockRawParam.Query = "Batman"

		mockOmdbRepo := new(mocks.MoviesRepository)
		mockOmdbRepo.On("SearchMovies", mock.Anything, mockParam).Return("rawURL", "rawResp", mockSearchResponse, nil).Once()
		mockMysqlRepo := new(mocks.MysqlOmdbRepository)
		mockMysqlRepo.On("LogRequest", mock.Anything, "rawURL", "rawResp").Return(int64(1), nil).Once()

		u := movieUsecase.NewMovieUsecase(time.Second*3, mockOmdbRepo, mockMysqlRepo)
		resp := u.SearchProcess(context.Background(), mockRawParam)

		assert.Equal(t, 200, resp.Status)
		assert.NotEmpty(t, resp.Data)
	})
}
