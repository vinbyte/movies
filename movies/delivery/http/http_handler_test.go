package http_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bxcodec/faker"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/vinbyte/movies/domain"
	"github.com/vinbyte/movies/domain/mocks"

	httpHandler "github.com/vinbyte/movies/movies/delivery/http"
)

func TestNewMovieHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	mockUsecase := new(mocks.MoviesUsecase)
	httpHandler.NewMovieHandler(gin.New(), mockUsecase)
}

func TestSearch(t *testing.T) {
	mockUsecase := new(mocks.MoviesUsecase)
	mockResponse := domain.BaseResponse{}
	res := domain.SearchResponse{}
	_ = faker.FakeData(&res)
	mockResponse.Data = res
	mockResponse.Status = 200
	mockResponse.Message = "success"
	mockParam := domain.SearchParam{}
	gin.SetMode(gin.TestMode)
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	var err error
	c.Request, err = http.NewRequest("GET", "/v1/search", nil)
	assert.Nil(t, err)
	c.Bind(&mockParam)
	ctx := c.Request.Context()
	mockUsecase.On("SearchProcess", ctx, mockParam).Return(mockResponse).Once()
	handler := httpHandler.MovieHandler{
		MovieUsecase: mockUsecase,
	}
	handler.Search(c)
	assert.Equal(t, 200, rec.Code)
}

func TestDetail(t *testing.T) {
	mockUsecase := new(mocks.MoviesUsecase)
	mockResponse := domain.BaseResponse{}
	res := domain.DetailResponse{}
	_ = faker.FakeData(&res)
	mockResponse.Data = res
	mockResponse.Status = 200
	mockResponse.Message = "success"
	mockParam := domain.DetailParam{}
	gin.SetMode(gin.TestMode)
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	var err error
	c.Request, err = http.NewRequest("GET", "/v1/detail", nil)
	assert.Nil(t, err)
	c.Bind(&mockParam)
	ctx := c.Request.Context()
	mockUsecase.On("DetailProcess", ctx, mockParam).Return(mockResponse).Once()
	handler := httpHandler.MovieHandler{
		MovieUsecase: mockUsecase,
	}
	handler.Detail(c)
	assert.Equal(t, 200, rec.Code)
}
