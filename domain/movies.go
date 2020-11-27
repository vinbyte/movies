package domain

import "context"

// MoviesUsecase represent movies usecase
type MoviesUsecase interface {
	SearchProcess(ctx context.Context, param SearchParam) BaseResponse
	DetailProcess(ctx context.Context, param DetailParam) BaseResponse
}

// MoviesRepository represent movies repository
type MoviesRepository interface {
	SearchMovies(ctx context.Context, param SearchParam) (string, string, OmdbSearchResponse, error)
	DetailMovie(ctx context.Context, param DetailParam) (string, string, OmdbDetailResponse, error)
}

// MysqlOmdbRepository represent mysql omdb repository
type MysqlOmdbRepository interface {
	LogRequest(ctx context.Context, request string, response string) (int64, error)
}
