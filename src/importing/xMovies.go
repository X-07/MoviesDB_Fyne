package importing

import (
	"encoding/xml"
	"fmt"
	"io"
	"moviesDB/modele"
	"moviesDB/utils"
	"os"
	"strconv"
	"strings"
)

type moviesStruct struct {
	XMLName     xml.Name `xml:"collection"`
	Type        string   `xml:"type,attr"`
	AttrItems   string   `xml:"items,attr"`
	Version     string   `xml:"version,attr"`
	Information struct {
		MaxID string `xml:"maxId"`
	} `xml:"information"`
	Items []itemStruct `xml:"item"`
}
type itemStruct struct {
	ID               string      `xml:"id,attr"`
	Title            string      `xml:"title,attr"`
	Date             string      `xml:"date,attr"`
	Time             string      `xml:"time,attr"`
	Director         string      `xml:"director,attr"`
	Image            string      `xml:"image,attr"`
	Backpic          string      `xml:"backpic,attr"`
	Original         string      `xml:"original,attr"`
	WebPage          string      `xml:"webPage,attr"`
	Seen             string      `xml:"seen,attr"`
	Added            string      `xml:"added,attr"`
	Region           string      `xml:"region,attr"`
	Format           string      `xml:"format,attr"`
	Number           string      `xml:"number,attr"`
	Identifier       string      `xml:"identifier,attr"`
	Place            string      `xml:"place,attr"`
	Rating           string      `xml:"rating,attr"`
	Ratingpress      string      `xml:"ratingpress,attr"`
	Age              string      `xml:"age,attr"`
	Video            string      `xml:"video,attr"`
	Serie            string      `xml:"serie,attr"`
	Rank             string      `xml:"rank,attr"`
	Trailer          string      `xml:"trailer,attr"`
	FControle        string      `xml:"f_Controle,attr"`
	FRemplacement    string      `xml:"f_Remplacement,attr"`
	FSupprime        string      `xml:"f_Supprime,attr"`
	FElodie          string      `xml:"f_Elodie,attr"`
	FElodie2         string      `xml:"f_Elodie2,attr"`
	FARemplacer      string      `xml:"f_ARemplacer,attr"`
	FAbsent          string      `xml:"f_Absent,attr"`
	FAReencoder      string      `xml:"f_AReencoder,attr"`
	FDecalage        string      `xml:"f_Decalage,attr"`
	FQualite         string      `xml:"f_Qualite,attr"`
	FTS              string      `xml:"f_TS,attr"`
	FVFQ             string      `xml:"f_VFQ,attr"`
	FMD              string      `xml:"f_MD,attr"`
	FSon             string      `xml:"f_Son,attr"`
	FVOSTFR          string      `xml:"f_VOSTFR,attr"`
	FAutrePB         string      `xml:"f_AutrePB,attr"`
	FRIPQualite      string      `xml:"f_RIPQualite,attr"`
	FVCD             string      `xml:"f_VCD,attr"`
	FVHS             string      `xml:"f_VHS,attr"`
	FVGA             string      `xml:"f_VGA,attr"`
	FDVD             string      `xml:"f_DVD,attr"`
	FSVGA            string      `xml:"f_sVGA,attr"`
	FXGA             string      `xml:"f_XGA,attr"`
	FHD              string      `xml:"f_HD,attr"`
	FFHD             string      `xml:"f_fHD,attr"`
	FWqHD            string      `xml:"f_wqHD,attr"`
	F4K              string      `xml:"f_4K,attr"`
	F8K              string      `xml:"f_8K,attr"`
	FEncodage        string      `xml:"f_Encodage,attr"`
	FLow             string      `xml:"f_Low,attr"`
	FGood            string      `xml:"f_Good,attr"`
	FHight           string      `xml:"f_Hight,attr"`
	FWEBDL           string      `xml:"f_WEB-DL,attr"`
	FAnimation       string      `xml:"f_Animation,attr"`
	FDocu            string      `xml:"f_Docu,attr"`
	FTaille          string      `xml:"f_Taille,attr"`
	FConteneur       string      `xml:"f_Conteneur,attr"`
	FDebitT          string      `xml:"f_DebitT,attr"`
	FCodecV          string      `xml:"f_CodecV,attr"`
	F3DType          string      `xml:"f_3DType,attr"`
	FCadence         string      `xml:"f_Cadence,attr"`
	FDebitV          string      `xml:"f_DebitV,attr"`
	FLargeur         string      `xml:"f_Largeur,attr"`
	FHauteur         string      `xml:"f_Hauteur,attr"`
	FCodecA          string      `xml:"f_CodecA,attr"`
	FAudio           string      `xml:"f_Audio,attr"`
	FEchantillonnage string      `xml:"f_Echantillonnage,attr"`
	FDebitA          string      `xml:"f_DebitA,attr"`
	FSousTitres      string      `xml:"f_SousTitres,attr"`
	Borrower         string      `xml:"borrower,attr"`
	LendDate         string      `xml:"lendDate,attr"`
	Borrowings       string      `xml:"borrowings,attr"`
	Favourite        string      `xml:"favourite,attr"`
	Tags             string      `xml:"tags,attr"`
	Synopsis         string      `xml:"synopsis"`
	Comment          string      `xml:"comment"`
	Country          *countryXML `xml:"country,omitempty"`
	Genre            *genreXML   `xml:"genre,omitempty"`
	Actors           *actorsXML  `xml:"actors,omitempty"`
	Audio            *audioXML   `xml:"audio,omitempty"`
	Subt             *subtXML    `xml:"subt,omitempty"`
}

// chargeBase() charge la collection movies/Cartoons/Humour et crée l'index
func (movies *moviesStruct) chargeBase(bddGCStarCollec string) {
	// Open collec
	xmlCollec, err := os.Open(bddGCStarCollec)
	if err != nil {
		panic(fmt.Sprint(" ouverture de la base : "+collec, err))
	}
	defer xmlCollec.Close()
	byteValue, _ := io.ReadAll(xmlCollec)

	// conversion du byteArray en struct
	xml.Unmarshal(byteValue, &movies)

}

// saveBase() : sauvegarde de la collection movies/Cartoons/Humour
func (movies *moviesStruct) saveBase() {
	idx := 0
	for _, movie := range movies.Items {
		var movieDB modele.Movie
		movieDB.ID = utils.AtoI64(movie.ID)
		movieDB.Title = movie.Title
		movieDB.DateSortie = movie.Date
		movieDB.Duration = utils.AtoI(movie.Time)
		movieDB.Directors = movie.Director
		image := strings.Split(movie.Image, "/")
		movieDB.Picture = image[len(image)-1]
		movieDB.OriginalTitle = movie.Original
		movieDB.Seen = utils.AtoB(movie.Seen)
		movieDB.DateAjout = movie.Added
		value, err := strconv.Atoi(movie.Ratingpress)
		if err != nil {
			movieDB.Rating = 0
		} else {
			if value < 11 {
				movieDB.Rating = value
			} else {
				movieDB.Rating = 0
			}
		}
		movieDB.Rating = utils.AtoI(movie.Rating)
		value, err = strconv.Atoi(movie.Ratingpress)
		if err != nil {
			movieDB.RatingPress = 0
		} else {
			if value < 11 {
				movieDB.RatingPress = value
			} else if value > 99 {
				movieDB.RatingPress = value / 100
			} else {
				movieDB.RatingPress = 0
			}
		}
		movieDB.RatingPress = utils.AtoI(movie.Ratingpress)
		if movie.Age == "1" {
			movieDB.AgeMini = "Inconnu"
		} else if movie.Age == "2" {
			movieDB.AgeMini = "Aucune restriction"
		} else if movie.Age == "5" {
			movieDB.AgeMini = "Accord parental"
		} else {
			movieDB.AgeMini = movie.Age
		}
		movieDB.Control = utils.AtoB(movie.FControle) // 1 = true, 0 = False
		movieDB.ReplaceInProgress = utils.AtoB(movie.FRemplacement)
		movieDB.Deleted = utils.AtoB(movie.FSupprime)
		movieDB.Replace = utils.AtoB(movie.FARemplacer)
		movieDB.Missing = utils.AtoB(movie.FAbsent)
		movieDB.ToReEncode = utils.AtoB(movie.FAReencoder)
		movieDB.TimeLag = utils.AtoB(movie.FDecalage)
		movieDB.BADQuality = utils.AtoB(movie.FQualite)
		movieDB.TS = utils.AtoB(movie.FTS)
		movieDB.VFQ = utils.AtoB(movie.FVFQ)
		movieDB.MD = utils.AtoB(movie.FMD)
		movieDB.Sound = utils.AtoB(movie.FSon)
		movieDB.VOSTFR = utils.AtoB(movie.FVOSTFR)
		movieDB.OtherPb = movie.FAutrePB
		if movie.FVCD == "1" {
			movieDB.RIPQuality = "VCD"
		}
		if movie.FVHS == "1" {
			movieDB.RIPQuality = "VHS"
		}
		if movie.FVGA == "1" {
			movieDB.RIPQuality = "VGA"
		}
		if movie.FDVD == "1" {
			movieDB.RIPQuality = "DVD"
		}
		if movie.FSVGA == "1" {
			movieDB.RIPQuality = "sVGA"
		}
		if movie.FXGA == "1" {
			movieDB.RIPQuality = "XGA"
		}
		if movie.FHD == "1" {
			movieDB.RIPQuality = "HD"
		}
		if movie.FFHD == "1" {
			movieDB.RIPQuality = "FHD"
		}
		if movie.FWqHD == "1" {
			movieDB.RIPQuality = "wqHD"
		}
		if movie.F4K == "1" {
			movieDB.RIPQuality = "4K"
		}
		if movie.F8K == "1" {
			movieDB.RIPQuality = "8K"
		}

		if movie.FLow == "1" {
			movieDB.EncQuality = "Light"
		}
		if movie.FGood == "1" {
			movieDB.EncQuality = "Good"
		}
		if movie.FHight == "1" {
			movieDB.EncQuality = "Hight"
		}

		if movie.FWEBDL == "1" {
			movieDB.Source = "WEB-DL"
		}
		movieDB.FileSize = utils.AtoF(movie.FTaille)
		movieDB.Container = movie.FConteneur
		movieDB.BitRateT = utils.AtoI(movie.FDebitT)
		movieDB.CodecV = movie.FCodecV
		movieDB.FrameRate = movie.FCadence
		movieDB.Type3D = movie.F3DType
		movieDB.BitRateV = utils.AtoI(movie.FDebitV)
		movieDB.Width = utils.AtoI(movie.FLargeur)
		movieDB.Height = utils.AtoI(movie.FHauteur)
		movieDB.CodecA = movie.FCodecA
		movieDB.Audio = movie.FAudio
		movieDB.Sampling = utils.AtoI(movie.FEchantillonnage)
		movieDB.BitRateA = utils.AtoI(movie.FDebitA)
		movieDB.Subtitles = movie.FSousTitres
		subtitles := strings.Split(strings.ReplaceAll(movie.FSousTitres, " ", ""), "/")
		movieDB.Subtitles = strings.Join(subtitles, ", ")
		movieDB.Synopsis = movie.Synopsis
		movieDB.Comment = movie.Comment
		if movie.Country != nil {
			for _, country := range movie.Country.Lines {
				if movieDB.Countries != "" {
					movieDB.Countries += ", "
				}
				movieDB.Countries += country.Col
			}
		}
		if movie.Genre != nil {
			for _, genre := range movie.Genre.Lines {
				if movieDB.Genres != "" {
					movieDB.Genres += ", "
				}
				movieDB.Genres += genre.Col
			}
		}
		if movie.FAnimation == "1" {
			if !strings.Contains(movieDB.Genres, "Animation") {
				movieDB.Genres += "Animation"
			}
		}
		if movie.FDocu == "1" {
			if !strings.Contains(movieDB.Genres, "Documentaire") {
				movieDB.Genres += "Documentaire"
			}
		}
		if movie.Actors != nil {
			for _, actor := range movie.Actors.Lines {
				if movieDB.Actors != "" {
					movieDB.Actors += ", "
				}
				movieDB.Actors += actor.Col[0]
			}
		}
		if idx == 0 {
			modele.InsertMovies(&movieDB)
		}
		idx++
		if idx == 14 {
			idx = 0
		}
	}
}
