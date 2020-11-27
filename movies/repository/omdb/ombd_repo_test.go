package omdb_test

import (
	"context"
	"testing"
	"time"

	"github.com/bxcodec/faker"
	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
	"github.com/vinbyte/movies/domain"
	omdbRepo "github.com/vinbyte/movies/movies/repository/omdb"
)

func TestSearchMovies(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockOmdbResponse := domain.OmdbSearchResponse{}
		_ = faker.FakeData(&mockOmdbResponse)
		mockOmdbResponse.Response = "True"
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder("GET", "http://www.omdbapi.com",
			httpmock.NewJsonResponderOrPanic(200, mockOmdbResponse))
		mockParam := domain.SearchParam{}
		mockParam.Query = "Batman"
		mockParam.PageStr = "1"
		or := omdbRepo.NewOmdbRepository(time.Second * 3)
		rawURL, rawResp, data, err := or.SearchMovies(context.Background(), mockParam)
		assert.Nil(t, err)
		assert.Equal(t, "True", data.Response)
		assert.NotEqual(t, "", rawURL)
		assert.NotEqual(t, "", rawResp)
	})
	t.Run("failed", func(t *testing.T) {
		mockOmdbResponse := domain.OmdbSearchResponse{}
		_ = faker.FakeData(&mockOmdbResponse)
		mockOmdbResponse.Response = "True"
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder("GET", "ws://www.omdbapi.com",
			httpmock.NewJsonResponderOrPanic(200, mockOmdbResponse))
		mockParam := domain.SearchParam{}
		mockParam.Query = "Batman"
		mockParam.PageStr = "1"
		or := omdbRepo.NewOmdbRepository(time.Second * 3)
		rawURL, rawResp, data, err := or.SearchMovies(context.Background(), mockParam)
		assert.NotNil(t, err)
		assert.Equal(t, "", data.Response)
		assert.NotEqual(t, "", rawURL)
		assert.Equal(t, "", rawResp)
	})
}

func TestDetailMovie(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockOmdbResponse := domain.OmdbDetailResponse{}
		_ = faker.FakeData(&mockOmdbResponse)
		mockOmdbResponse.Response = "True"
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder("GET", "http://www.omdbapi.com",
			httpmock.NewJsonResponderOrPanic(200, mockOmdbResponse))
		mockParam := domain.DetailParam{}
		mockParam.ID = "aaaa"
		or := omdbRepo.NewOmdbRepository(time.Second * 3)
		rawURL, rawResp, data, err := or.DetailMovie(context.Background(), mockParam)
		assert.Nil(t, err)
		assert.Equal(t, "True", data.Response)
		assert.NotEqual(t, "", rawURL)
		assert.NotEqual(t, "", rawResp)
	})
	t.Run("failed", func(t *testing.T) {
		mockOmdbResponse := domain.OmdbDetailResponse{}
		_ = faker.FakeData(&mockOmdbResponse)
		mockOmdbResponse.Response = "True"
		httpmock.Activate()
		defer httpmock.DeactivateAndReset()
		httpmock.RegisterResponder("GET", "ws://www.omdbapi.com",
			httpmock.NewJsonResponderOrPanic(200, mockOmdbResponse))
		mockParam := domain.DetailParam{}
		mockParam.ID = "aaaa"
		or := omdbRepo.NewOmdbRepository(time.Second * 3)
		rawURL, rawResp, data, err := or.DetailMovie(context.Background(), mockParam)
		assert.NotNil(t, err)
		assert.Equal(t, "", data.Response)
		assert.NotEqual(t, "", rawURL)
		assert.Equal(t, "", rawResp)
	})
}
