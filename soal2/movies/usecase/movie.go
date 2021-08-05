package usecase

import (
	"encoding/json"
	"log"
	movies "movies/proto"
	"net/http"
	"strconv"

	"github.com/micro/micro/v3/service/config"
)

type movieDataWrapper struct {
	Response     string      `json:"Response"`
	TotalResults string      `json:"totalResults"`
	Data         []movieData `json:"Search"`
}

type movieData struct {
	Title      string        `json:"Title"`
	Year       string        `json:"Year"`
	Rated      string        `json:"Rated"`
	Released   string        `json:"Released"`
	Runtime    string        `json:"Runtime"`
	Genre      string        `json:"Genre"`
	Director   string        `json:"Director"`
	Writer     string        `json:"Writer"`
	Actors     string        `json:"Actors"`
	Plot       string        `json:"Plot"`
	Language   string        `json:"Language"`
	Country    string        `json:"Country"`
	Awards     string        `json:"Awards"`
	Poster     string        `json:"Poster"`
	Metascore  string        `json:"Metascore"`
	ImdbRating string        `json:"imdbRating"`
	ImdbVotes  string        `json:"imdbVotes"`
	ImdbID     string        `json:"imdbID"`
	Type       string        `json:"Type"`
	DVD        string        `json:"DVD"`
	BoxOffice  string        `json:"BoxOffice"`
	Production string        `json:"Production"`
	Website    string        `json:"Website"`
	Ratings    []movieRating `json:"Ratings"`
}

type movieRating struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}

func MovieSerializer(mv movieData, resp *movies.MovieData) {
	resp.Title = mv.Title
	resp.Year = mv.Year
	resp.Rated = mv.Rated
	resp.Released = mv.Released
	resp.Runtime = mv.Runtime
	resp.Genre = mv.Genre
	resp.Director = mv.Director
	resp.Writer = mv.Writer
	resp.Actors = mv.Actors
	resp.Plot = mv.Plot
	resp.Language = mv.Language
	resp.Country = mv.Country
	resp.Awards = mv.Awards
	resp.PosterUrl = mv.Poster
	resp.Metascore = mv.Metascore
	resp.ImdbRating = mv.ImdbRating
	resp.ImdbVotes = mv.ImdbVotes
	resp.ImdbID = mv.ImdbID
	resp.Type = mv.Type
	resp.Dvd = mv.DVD
	resp.BoxOffice = mv.BoxOffice
	resp.Production = mv.Production
	resp.Website = mv.Website
	resp.Ratings = []*movies.MovieRating{}
	for _, elem := range mv.Ratings {
		temp := &movies.MovieRating{
			Source: elem.Source,
			Value:  elem.Value,
		}
		resp.Ratings = append(resp.Ratings, temp)
	}
}

func CreateImdbDetailRequest(id string) (result movieData, err error) {
	val, err := config.Get("movies.imdb_api_key")
	apiKey := val.String("")
	val, err = config.Get("movies.imdb_base_url")
	imdb_base_url := val.String("")

	client := &http.Client{}
	httpReq, _ := http.NewRequest("GET", imdb_base_url, nil)
	q := httpReq.URL.Query()
	q.Add("apikey", apiKey)
	q.Add("i", id)
	httpReq.URL.RawQuery = q.Encode()

	httpRsp, err := client.Do(httpReq)
	if err != nil {
		log.Fatal(err)
	}
	defer httpRsp.Body.Close()

	jsonRsp := movieData{}
	json.NewDecoder(httpRsp.Body).Decode(&jsonRsp)

	return jsonRsp, err
}

func CreateImdbSearchRequest(query string, currentPage int) (result movieDataWrapper, err error) {
	val, err := config.Get("movies.imdb_api_key")
	apiKey := val.String("")
	val, err = config.Get("movies.imdb_base_url")
	imdb_base_url := val.String("")

	client := &http.Client{}
	httpReq, _ := http.NewRequest("GET", imdb_base_url, nil)
	q := httpReq.URL.Query()
	q.Add("apikey", apiKey)
	q.Add("s", query)
	q.Add("page", strconv.Itoa(currentPage))
	httpReq.URL.RawQuery = q.Encode()

	httpRsp, err := client.Do(httpReq)
	if err != nil {
		log.Fatal(err)
	}
	defer httpRsp.Body.Close()

	jsonRsp := movieDataWrapper{}
	json.NewDecoder(httpRsp.Body).Decode(&jsonRsp)

	return jsonRsp, err
}
