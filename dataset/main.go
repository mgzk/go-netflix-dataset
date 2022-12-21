package main

import (
	"database/sql"
	"fmt"
	"github.com/go-the-way/exl"
	_ "github.com/lib/pq"
	"os"
	"strconv"
	"time"
)

type DatasetModel struct {
	Title               string  `excel:"Title"`
	Genre               string  `excel:"Genre"`
	Tags                string  `excel:"Tags"`
	Languages           string  `excel:"Languages"`
	Type                string  `excel:"Series or Movie"`
	Score               float32 `excel:"Hidden Gem Score"`
	CountryAvailability string  `excel:"Country Availability"`
	Runtime             string  `excel:"Runtime"`
	Director            string  `excel:"Director"`
	Writer              string  `excel:"Writer"`
	Actors              string  `excel:"Actors"`
	ViewRating          string  `excel:"View Rating"`
	ImdbScore           float32 `excel:"IMDb Score"`
	RottenTomatoesScore int     `excel:"Rotten Tomatoes Score"`
	MetacriticScore     int     `excel:"Metacritic Score"`
	AwardsReceived      int     `excel:"Awards Received"`
	AwardsNominatedFor  int     `excel:"Awards Nominated For"`
	Boxoffice           string  `excel:"Boxoffice"`
	ReleaseDate         string  `excel:"Release Date"`
	NetflixReleaseDate  string  `excel:"Netflix Release Date"`
	ProductionHouse     string  `excel:"Production House"`
	NetflixLink         string  `excel:"Netflix Link"`
	IMDbLink            string  `excel:"IMDb Link"`
	Summary             string  `excel:"Summary"`
	IMDbVotes           int     `excel:"IMDb Votes"`
	Image               string  `excel:"Image"`
	Poster              string  `excel:"Poster"`
	TMDbTrailer         string  `excel:"TMDb Trailer"`
	TrailerSite         string  `excel:"Trailer Site"`
}

const insertQuery = "insert into film (id, title, genre, tags, languages, type, score, country_availability, runtime, director, writer, actors, view_rating, " +
	"imdb_score, rotten_tomatoes_score, metacritic_score, awards_received, awards_nominated_for, boxoffice, release_date, netflix_release_date, production_house, netflix_link," +
	"imdb_link, summary, imdb_votes, image, poster, tmdb_trailer, trailer_site) values (nextval('film_id_seq'), $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, " +
	"$13, $14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24, $25, $26, $27, $28, $29)"

const selectQuery = "select id, title from film"

const datasetPath = "./dataset/Netflix-Dataset-2021.xlsx"

const dateLayout = "2006-01-02"

type DatabaseRow struct {
	ID    int
	Title string
}

func (*DatasetModel) ConfigureRM(*exl.ReadMetadata) {}

func main() {
	db, err := sql.Open("postgres", os.Args[1])

	if err != nil {
		fmt.Printf(err.Error())
	}

	models := read(datasetPath)

	for _, model := range models {
		insert(model, db)
	}
}

func insert(model *DatasetModel, db *sql.DB) {
	_, err := db.Exec(insertQuery, model.Title, model.Genre, model.Tags, model.Languages, model.Type, model.Score, model.CountryAvailability, model.Runtime,
		model.Director,
		model.Writer, model.Actors, model.ViewRating, model.ImdbScore, model.RottenTomatoesScore, model.MetacriticScore, model.AwardsReceived, model.AwardsNominatedFor,
		model.Boxoffice, convertDate(model.ReleaseDate), convertDate(model.NetflixReleaseDate), model.ProductionHouse, model.NetflixLink, model.IMDbLink,
		model.Summary, model.IMDbVotes,
		model.Image, model.Poster,
		model.TMDbTrailer,
		model.TrailerSite)

	if err != nil {
		fmt.Printf("Title: %s, Error: %s \n", model.Title, err.Error())
	}
}

func convertDate(excelDate string) *string {
	if len(excelDate) == 0 {
		return nil
	}

	dateInt, err := strconv.Atoi(excelDate)

	if err != nil {
		fmt.Printf(err.Error())
	}

	date := time.Date(1900, 1, 0, 0, 0, 0, 0, time.UTC).AddDate(0, 0, dateInt-1).Format(dateLayout)

	return &date
}

func read(path string) []*DatasetModel {
	var models []*DatasetModel

	models, err := exl.ReadFile(path, filter)

	if err != nil {
		fmt.Printf(err.Error())
	}

	return models
}

func filter(*DatasetModel) bool {
	return true
}
