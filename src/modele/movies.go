package modele

import (
	"database/sql"
	"fmt"
	"sort"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

var Db *sql.DB
var err error

type Movie struct {
	ID                int64
	Title             string
	OriginalTitle     string
	Picture           string
	Trailer           string //Bande annonce
	DateAjout         string
	DateModif         string
	DateSortie        string
	Directors         string
	Duration          int
	AgeMini           string
	Countries         string
	Genres            string
	Actors            string
	Synopsis          string
	Seen              bool
	BadMovie          bool
	Rating            int
	RatingPress       int
	Comment           string
	Control           bool
	Replace           bool
	ReplaceInProgress bool
	Deleted           bool
	Missing           bool
	ToReEncode        bool
	TimeLag           bool
	BADQuality        bool
	TS                bool
	MD                bool
	Sound             bool
	VFQ               bool
	VOSTFR            bool
	OtherPb           string
	RIPQuality        string
	EncQuality        string
	Source            string
	FileSize          float32
	Container         string
	BitRateT          int
	CodecV            string
	Type3D            string
	FrameRate         string
	BitRateV          int
	Width             int
	Height            int
	CodecA            string
	Audio             string
	Sampling          int
	BitRateA          int
	Subtitles         string
}

type MovieList struct {
	ID        int64
	Title     string
	Seen      bool
	DateAjout int32
}

func OpenDB(dbName string) {
	Db, err = sql.Open("sqlite3", dbName)
	if err != nil {
		panic(err.Error())
	}
}

func CreateMoviesTable() {
	// In the definition below, please note that BLOB is the only type we can use in sqlite for storing JSON.
	_, err = Db.Exec(reqCreateMovies)
	if err != nil {
		panic(err.Error())
	}
}

// InsertMovies : insertion d'un enregistrement dans la table.
func InsertMovies(movie *Movie) {
	stmt, err := Db.Prepare(sqlInsertMovies)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()
	res, err := stmt.Exec(
		movie.Title,
		movie.OriginalTitle,
		movie.Picture,
		movie.Trailer,
		movie.DateAjout,
		movie.DateModif,
		movie.DateSortie,
		movie.Directors,
		movie.Duration,
		movie.AgeMini,
		movie.Countries,
		movie.Genres,
		movie.Actors,
		movie.Synopsis,
		movie.Seen,
		movie.BadMovie,
		movie.Rating,
		movie.RatingPress,
		movie.Comment,
		movie.Control,
		movie.Replace,
		movie.ReplaceInProgress,
		movie.Deleted,
		movie.Missing,
		movie.ToReEncode,
		movie.TimeLag,
		movie.BADQuality,
		movie.TS,
		movie.MD,
		movie.Sound,
		movie.VFQ,
		movie.VOSTFR,
		movie.OtherPb,
		movie.RIPQuality,
		movie.EncQuality,
		movie.Source,
		movie.FileSize,
		movie.Container,
		movie.BitRateT,
		movie.CodecV,
		movie.Type3D,
		movie.FrameRate,
		movie.BitRateV,
		movie.Width,
		movie.Height,
		movie.CodecA,
		movie.Audio,
		movie.Sampling,
		movie.BitRateA,
		movie.Subtitles,
	)
	if err != nil {
		panic(err.Error())
	}
	id, err := res.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	movie.ID = id
}

// UpdateMovies : Mise à jour d'un enregistrement dans la table.
func UpdateMovies(movie *Movie) {
	stmt, err := Db.Prepare(sqlUpdateMovies)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		movie.Title,
		movie.OriginalTitle,
		movie.Picture,
		movie.Trailer,
		movie.DateAjout,
		movie.DateModif,
		movie.DateSortie,
		movie.Directors,
		movie.Duration,
		movie.AgeMini,
		movie.Countries,
		movie.Genres,
		movie.Actors,
		movie.Synopsis,
		movie.Seen,
		movie.BadMovie,
		movie.Rating,
		movie.RatingPress,
		movie.Comment,
		movie.Control,
		movie.Replace,
		movie.ReplaceInProgress,
		movie.Deleted,
		movie.Missing,
		movie.ToReEncode,
		movie.TimeLag,
		movie.BADQuality,
		movie.TS,
		movie.MD,
		movie.Sound,
		movie.VFQ,
		movie.VOSTFR,
		movie.OtherPb,
		movie.RIPQuality,
		movie.EncQuality,
		movie.Source,
		movie.FileSize,
		movie.Container,
		movie.BitRateT,
		movie.CodecV,
		movie.Type3D,
		movie.FrameRate,
		movie.BitRateV,
		movie.Width,
		movie.Height,
		movie.CodecA,
		movie.Audio,
		movie.Sampling,
		movie.BitRateA,
		movie.Subtitles,
		movie.ID,
	)
	if err != nil {
		panic(err.Error())
	}
}

// GetMoviesByID : Restitue l'enregistrement de la table.
func GetMoviesByID(id int64) Movie {
	rows, err := Db.Query(sqlGetMoviesByID, id)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var movie Movie
	if rows.Next() {
		err := rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.OriginalTitle,
			&movie.Picture,
			&movie.Trailer,
			&movie.DateAjout,
			&movie.DateModif,
			&movie.DateSortie,
			&movie.Directors,
			&movie.Duration,
			&movie.AgeMini,
			&movie.Countries,
			&movie.Genres,
			&movie.Actors,
			&movie.Synopsis,
			&movie.Seen,
			&movie.BadMovie,
			&movie.Rating,
			&movie.RatingPress,
			&movie.Comment,
			&movie.Control,
			&movie.Replace,
			&movie.ReplaceInProgress,
			&movie.Deleted,
			&movie.Missing,
			&movie.ToReEncode,
			&movie.TimeLag,
			&movie.BADQuality,
			&movie.TS,
			&movie.MD,
			&movie.Sound,
			&movie.VFQ,
			&movie.VOSTFR,
			&movie.OtherPb,
			&movie.RIPQuality,
			&movie.EncQuality,
			&movie.Source,
			&movie.FileSize,
			&movie.Container,
			&movie.BitRateT,
			&movie.CodecV,
			&movie.Type3D,
			&movie.FrameRate,
			&movie.BitRateV,
			&movie.Width,
			&movie.Height,
			&movie.CodecA,
			&movie.Audio,
			&movie.Sampling,
			&movie.BitRateA,
			&movie.Subtitles,
		)
		if err != nil {
			panic(err.Error())
		}
	}
	return movie
}

// DeleteMoviesByID : détruit un enregistrement, identifié par son ID, dans la table et retourne le nb d'enregistrement supprimé (en principe 1).
func DeleteMoviesByID(id int64) int64 {
	stmt, err := Db.Prepare(sqlDeleteMoviesByID)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	res, err := stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}
	affect, err := res.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	return affect
}

// GetMoviesList : Restitue une vue (4 colonnes) de tous les enregistrements de la table.
func GetMoviesList() []MovieList {
	var movies []MovieList
	rows, err := Db.Query(sqlGetMoviesList)
	if err != nil {
		//panic(err.Error())
		fmt.Println(err.Error())
		return movies
	}
	defer rows.Close()

	for rows.Next() {
		var movie MovieList
		err = rows.Scan(
			&movie.ID,
			&movie.Title,
			&movie.Seen,
			&movie.DateAjout,
		)
		if err != nil {
			panic(err.Error())
		}
		movies = append(movies, movie)
	}
	return movies
}

// GetContainerList : Restitue tous les containers de la table.
// []string{".avi", ".mkv", ".mp4", ".mpg", ".mv", ".ogm", ".wmv"}
func GetContainerList() []string {
	rows, err := Db.Query(sqlGetContainerList)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var result []string
	for rows.Next() {
		var res string
		err = rows.Scan(
			&res,
		)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, res)
	}
	return result
}

// GetCodecVideoList : Restitue tous les codecs vidéo de la table.
// []string{"DivX 3", "DivX 3 Low", "DivX 4", "DivX 5", "MPEG-1", "MPEG-4", "Microsoft", "X264", "X264 - 1.3", "X264 - 2.1", "X264 - 2.2", "X264 - 3.0", "X264 - 3.1", "X264 - 3.2", "X264 - 4.0", "X264 - 4.1", "X264 - 4.2", "X264 - 5.0", "X264 - 5.1", "X264 - 5.2", "X265", "XviD", "????"}
func GetCodecVideoList() []string {
	rows, err := Db.Query(sqlGetCodecVideoList)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var result []string
	for rows.Next() {
		var res string
		err = rows.Scan(
			&res,
		)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, res)
	}
	return result
}

// GetCodecAudioList : Restitue tous les codecs audio de la table.
// []string{"AAC", "AC-3", "DTS", "E-AC-3", "MP2", "MP3", "PCM", "Vorbis", "WMA"}
func GetCodecAudioList() []string {
	rows, err := Db.Query(sqlGetCodecAudioList)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var result []string
	for rows.Next() {
		var res string
		err = rows.Scan(
			&res,
		)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, res)
	}
	return result
}

// GetAudioList : Restitue tous les audio type de la table.
// []string{"1ch: Mono", "2ch: Stereo", "3ch: Stereo 2.1", "5ch: Surround", "6ch: Surround", "8ch: Surround +"}
func GetAudioList() []string {
	rows, err := Db.Query(sqlGetAudioList)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var result []string
	for rows.Next() {
		var res string
		err = rows.Scan(
			&res,
		)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, res)
	}
	return result
}

// GetDirectorsList : Restitue tous Directors de la table.
// []string{"ASS", "PGS", "UTF-8", ..."}
func GetDirectorsList(saisie string) []string {
	rows, err := Db.Query(sqlGetDirectorsList(saisie))
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	resultMap := make(map[string]int)

	for rows.Next() {
		var res string
		err = rows.Scan(
			&res,
		)
		if err != nil {
			panic(err.Error())
		}
		resultSlice := strings.Split(res, ", ")
		for _, resultElmt := range resultSlice {
			if strings.Contains(strings.ToLower(resultElmt), strings.ToLower(saisie)) {
				resultMap[strings.Trim(resultElmt, " ")] = 0
			}
		}
	}
	var result []string
	for resultElmt := range resultMap {
		result = append(result, resultElmt)
	}
	sort.Strings(result)
	return result
}

// GetCountriesList : Restitue tous les Countries de la table.
// []string{"ASS", "PGS", "UTF-8", ..."}
func GetCountriesList(saisie string) []string {
	rows, err := Db.Query(sqlGetCountriesList(saisie))
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	resultMap := make(map[string]int)

	for rows.Next() {
		var res string
		err = rows.Scan(
			&res,
		)
		if err != nil {
			panic(err.Error())
		}
		resultSlice := strings.Split(res, ", ")
		for _, resultElmt := range resultSlice {
			if strings.Contains(strings.ToLower(resultElmt), strings.ToLower(saisie)) {
				resultMap[strings.Trim(resultElmt, " ")] = 0
			}
		}
	}
	var result []string
	for resultElmt := range resultMap {
		result = append(result, resultElmt)
	}
	sort.Strings(result)
	return result
}

// GetGenresList : Restitue tous les Genres de la table.
// []string{"ASS", "PGS", "UTF-8", ..."}
func GetGenresList(saisie string) []string {
	rows, err := Db.Query(sqlGetGenresList(saisie))
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	resultMap := make(map[string]int)

	for rows.Next() {
		var res string
		err = rows.Scan(
			&res,
		)
		if err != nil {
			panic(err.Error())
		}
		resultSlice := strings.Split(res, ", ")
		for _, resultElmt := range resultSlice {
			if strings.Contains(strings.ToLower(resultElmt), strings.ToLower(saisie)) {
				resultMap[strings.Trim(resultElmt, " ")] = 0
			}
		}
	}
	var result []string
	for resultElmt := range resultMap {
		result = append(result, resultElmt)
	}
	sort.Strings(result)
	return result
}

// GetActorsList : Restitue tous les Actors de la table.
// []string{"ASS", "PGS", "UTF-8", ..."}
func GetActorsList(saisie string) []string {
	rows, err := Db.Query(sqlGetActorsList(saisie))
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	resultMap := make(map[string]int)

	for rows.Next() {
		var res string
		err = rows.Scan(
			&res,
		)
		if err != nil {
			panic(err.Error())
		}
		resultSlice := strings.Split(res, ", ")
		for _, resultElmt := range resultSlice {
			if strings.Contains(strings.ToLower(resultElmt), strings.ToLower(saisie)) {
				resultMap[strings.Trim(resultElmt, " ")] = 0
			}
		}
	}
	var result []string
	for resultElmt := range resultMap {
		result = append(result, resultElmt)
	}
	sort.Strings(result)
	return result
}

// GetSubtitleList : Restitue tous les Subtitle de la table.
// []string{"ASS", "PGS", "UTF-8", ..."}
func GetSubtitleList() []string {
	rows, err := Db.Query(sqlGetSubtitleList)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	resultMap := make(map[string]int)

	for rows.Next() {
		var res string
		err = rows.Scan(
			&res,
		)
		if err != nil {
			panic(err.Error())
		}
		resultSlice := strings.Split(res, ", ")
		for _, resultElmt := range resultSlice {
			resultMap[strings.Trim(resultElmt, " ")] = 0
		}
	}
	var result []string
	for resultElmt := range resultMap {
		result = append(result, resultElmt)
	}
	sort.Strings(result)
	return result
}
