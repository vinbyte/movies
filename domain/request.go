package domain

// SearchParam is parameter for searching
type SearchParam struct {
	Query   string `form:"query"`
	PageStr string `form:"page"`
	Page    int
}

// DetailParam is parameter for get detail
type DetailParam struct {
	ID string `form:"imdb_id"`
}
