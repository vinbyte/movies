package usecase

import (
	"time"

	"github.com/vinbyte/movies/domain"
)

type movieUsecase struct {
	mysqlRepo      domain.MysqlOmdbRepository
	omdbRepo       domain.MoviesRepository
	contextTimeout time.Duration
}

// NewMovieUsecase will create an object that represent domain.MoviesUsecase
func NewMovieUsecase(timeout time.Duration, mr domain.MoviesRepository, mor domain.MysqlOmdbRepository) domain.MoviesUsecase {
	return &movieUsecase{
		mysqlRepo:      mor,
		omdbRepo:       mr,
		contextTimeout: timeout,
	}
}
