package http

import (
	"github.com/gin-gonic/gin"
	"github.com/vinbyte/movies/domain"
)

// MovieHandler  represent the httphandler for movies
type MovieHandler struct {
	MovieUsecase domain.MoviesUsecase
}

// NewMovieHandler will initialize the course resources endpoint
func NewMovieHandler(r *gin.Engine, mu domain.MoviesUsecase) {
	handler := &MovieHandler{
		MovieUsecase: mu,
	}
	v1 := r.Group("v1")
	{
		v1.GET("/search", handler.Search)
		v1.GET("/detail", handler.Detail)
	}
}

// Search is perform searching movie
func (h *MovieHandler) Search(c *gin.Context) {
	ctx := c.Request.Context()
	var param domain.SearchParam
	c.Bind(&param)
	res := h.MovieUsecase.SearchProcess(ctx, param)
	c.JSON(res.Status, res)
}

// Detail is perform fetch detail movie
func (h *MovieHandler) Detail(c *gin.Context) {
	ctx := c.Request.Context()
	var param domain.DetailParam
	c.Bind(&param)
	res := h.MovieUsecase.DetailProcess(ctx, param)
	c.JSON(res.Status, res)
}
