package usecase

import (
	"context"

	"github.com/vinbyte/movies/domain"

	log "github.com/sirupsen/logrus"
)

func (u *movieUsecase) DetailProcess(ctx context.Context, param domain.DetailParam) domain.BaseResponse {
	var response domain.BaseResponse
	ctx, cancel := context.WithTimeout(ctx, u.contextTimeout)
	defer cancel()

	var emptyData struct{}
	response.Status = 200
	response.Data = emptyData
	response.Message = "success"

	if param.ID == "" {
		response.Status = 400
		response.Message = "imdb_id is required"
		return response
	}
	rawURL, rawResp, data, err := u.omdbRepo.DetailMovie(ctx, param)
	if err != nil {
		log.Error(err)
	}
	_, err = u.mysqlRepo.LogRequest(ctx, rawURL, rawResp)
	if err != nil {
		log.Error("LogRequest ", err)
	}
	if data.Response == "True" {
		response.Data = data
	} else {
		response.Status = 400
		response.Message = "failed to get detail"
	}

	return response
}
