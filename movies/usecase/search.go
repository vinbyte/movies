package usecase

import (
	"context"
	"strconv"

	"github.com/vinbyte/movies/domain"

	log "github.com/sirupsen/logrus"
)

func (u *movieUsecase) SearchProcess(ctx context.Context, param domain.SearchParam) domain.BaseResponse {
	var response domain.BaseResponse
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var emptyData struct{}
	response.Status = 200
	response.Data = emptyData
	response.Message = "success"

	var err error
	if param.PageStr == "" || param.PageStr == "0" {
		param.PageStr = "1"
	}
	param.Page, err = strconv.Atoi(param.PageStr)
	if err != nil {
		log.Error("param.PageStr : ", err)
		response.Status = 400
		response.Message = err.Error()
		return response
	}

	rawURL, rawResp, res, err := u.omdbRepo.SearchMovies(ctx, param)
	if err != nil {
		log.Error(res)
	}
	_, err = u.mysqlRepo.LogRequest(ctx, rawURL, rawResp)
	if err != nil {
		log.Error("LogRequest ", err)
	}
	var data domain.SearchResponse
	data.List = make([]domain.Movie, 0)
	if res.Response == "True" {
		for _, m := range res.Search {
			var movieData domain.Movie
			movieData.ImdbID = m.ImdbID
			movieData.Poster = m.Poster
			movieData.Title = m.Title
			movieData.Type = m.Type
			movieData.Year = m.Year
			data.List = append(data.List, movieData)
		}
		data.Response = true
		data.TotalResult, _ = strconv.Atoi(res.TotalResults)
		mod := data.TotalResult % 10
		bagi := data.TotalResult / 10
		if bagi != 0 && mod == 0 {
			data.TotalPage = bagi
		} else if bagi != 0 && mod != 0 {
			data.TotalPage = bagi + 1
		} else {
			data.TotalPage = 1
		}
		data.Page = param.Page
	} else {
		response.Status = 400
		response.Message = "please fill the query parameter"
	}
	response.Data = data

	return response
}
