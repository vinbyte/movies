package domain

// BaseResponse will be the basic json response in every endpoint
type BaseResponse struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

// OmdbSearchResponse is search response from OMDB API
type OmdbSearchResponse struct {
	Search []struct {
		Title  string `json:"Title"`
		Year   string `json:"Year"`
		ImdbID string `json:"imdbID"`
		Type   string `json:"Type"`
		Poster string `json:"Poster"`
	} `json:"Search"`
	TotalResults string `json:"totalResults"`
	Response     string `json:"Response"`
}

// OmdbDetailResponse is search response from OMDB API
type OmdbDetailResponse struct {
	Title    string `json:"Title"`
	Year     string `json:"Year"`
	Rated    string `json:"Rated"`
	Released string `json:"Released"`
	Runtime  string `json:"Runtime"`
	Genre    string `json:"Genre"`
	Director string `json:"Director"`
	Writer   string `json:"Writer"`
	Actors   string `json:"Actors"`
	Plot     string `json:"Plot"`
	Language string `json:"Language"`
	Country  string `json:"Country"`
	Awards   string `json:"Awards"`
	Poster   string `json:"Poster"`
	Ratings  []struct {
		Source string `json:"Source"`
		Value  string `json:"Value"`
	} `json:"Ratings"`
	Metascore  string `json:"Metascore"`
	ImdbRating string `json:"imdbRating"`
	ImdbVotes  string `json:"imdbVotes"`
	ImdbID     string `json:"imdbID"`
	Type       string `json:"Type"`
	DVD        string `json:"DVD"`
	BoxOffice  string `json:"BoxOffice"`
	Production string `json:"Production"`
	Website    string `json:"Website"`
	Response   string `json:"Response"`
}

// Movie is data structure for movie
type Movie struct {
	Title  string `json:"title"`
	Year   string `json:"year"`
	ImdbID string `json:"imdb_id"`
	Type   string `json:"type"`
	Poster string `json:"poster"`
}

// SearchResponse is response for search endpoint
type SearchResponse struct {
	List        []Movie `json:"list"`
	Page        int     `json:"page"`
	TotalPage   int     `json:"total_page"`
	TotalResult int     `json:"total_result"`
	Response    bool    `json:"response"`
}

// DetailResponse is response for detail endpoint
type DetailResponse struct {
	Title    string `json:"title"`
	Year     string `json:"year"`
	Rated    string `json:"rated"`
	Released string `json:"released"`
	Runtime  string `json:"runtime"`
	Genre    string `json:"genre"`
	Director string `json:"director"`
	Writer   string `json:"writer"`
	Actors   string `json:"actors"`
	Plot     string `json:"plot"`
	Language string `json:"language"`
	Country  string `json:"country"`
	Awards   string `json:"awards"`
	Poster   string `json:"poster"`
	Ratings  []struct {
		Source string `json:"source"`
		Value  string `json:"value"`
	} `json:"ratings"`
	Metascore  string `json:"metascore"`
	ImdbRating string `json:"imdbRating"`
	ImdbVotes  string `json:"imdbVotes"`
	ImdbID     string `json:"imdbID"`
	Type       string `json:"type"`
	DVD        string `json:"dvd"`
	BoxOffice  string `json:"boxOffice"`
	Production string `json:"production"`
	Website    string `json:"website"`
	Response   string `json:"response"`
}
