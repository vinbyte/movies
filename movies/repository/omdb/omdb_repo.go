package omdb

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/vinbyte/movies/domain"
)

type omdbRepository struct {
	Timeout time.Duration
}

// NewOmdbRepository will create an object that represent the general.Repository interface
func NewOmdbRepository(timeout time.Duration) domain.MoviesRepository {
	return &omdbRepository{timeout}
}

func (o *omdbRepository) SearchMovies(ctx context.Context, param domain.SearchParam) (string, string, domain.OmdbSearchResponse, error) {
	var data domain.OmdbSearchResponse
	var client = &http.Client{}
	req, err := http.NewRequest("GET", "http://www.omdbapi.com", nil)
	if err != nil {
		log.Error("http.NewRequest : ", err)
		return "", "", data, err
	}
	apiKey := os.Getenv("OMDB_KEY")
	q := req.URL.Query()
	q.Add("apikey", apiKey)
	q.Add("s", param.Query)
	q.Add("page", param.PageStr)
	req.URL.RawQuery = q.Encode()
	rawURL := req.URL.String()
	res, err := client.Do(req)
	if err != nil {
		log.Error("client.Do : ", err)
		return rawURL, "", data, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	rawResponse := string(body)
	res.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		log.Error("Decode : ", err)
		return rawURL, "", data, err
	}
	return rawURL, rawResponse, data, nil
}

func (o *omdbRepository) DetailMovie(ctx context.Context, param domain.DetailParam) (string, string, domain.OmdbDetailResponse, error) {
	var data domain.OmdbDetailResponse
	var client = &http.Client{}
	req, err := http.NewRequest("GET", "http://www.omdbapi.com", nil)
	if err != nil {
		log.Error("http.NewRequest : ", err)
		return "", "", data, err
	}
	apiKey := os.Getenv("OMDB_KEY")
	q := req.URL.Query()
	q.Add("apikey", apiKey)
	q.Add("i", param.ID)
	req.URL.RawQuery = q.Encode()
	rawURL := req.URL.String()
	res, err := client.Do(req)
	if err != nil {
		log.Error("client.Do : ", err)
		return rawURL, "", data, err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	rawResponse := string(body)
	res.Body = ioutil.NopCloser(bytes.NewBuffer(body))
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		log.Error("Decode : ", err)
		return rawURL, "", data, err
	}
	return rawURL, rawResponse, data, nil
}
