package importing

import "moviesDB/modele"

// const bddMovies = "Movies.gcs"

var BddGCStarDir = "/media/veracrypt60/GCStar/"
var collec string

var movies moviesStruct

func LoadMovies(gcstarName string, moviesDBName string) {
	modele.CreateMoviesTable()
	movies.chargeBase(gcstarName)
	modele.OpenDB(moviesDBName)
	modele.CreateMoviesTable()
	movies.saveBase()
}
