package modele

// reqCreateMovies : requête SQL de création de la table
const reqCreateMovies = `CREATE TABLE IF NOT EXISTS movies (
		id            		INTEGER PRIMARY KEY,
		title         		TEXT,
		originalTitle 		TEXT,
		picture				TEXT,
		trailer				TEXT,
		dateAjout     		TEXT,
		dateModif     		TEXT,
		dateSortie    		TEXT,
		directors   		TEXT,
		duration        	INTEGER,
		ageMini       		TEXT,
		countries    		TEXT,
		genres        		TEXT,
		actors       		TEXT,
		synopsis      		TEXT,
		seen           		INTEGER,
		badMovie			INTEGER,
		rating         		INTEGER,
		ratingPress    		INTEGER,
		comment				TEXT,
		control  			INTEGER,
		replace				INTEGER,
		replaceInProgress	INTEGER,
		deleted				INTEGER,
		missing				INTEGER,
		toReEncode			INTEGER,
		timeLag				INTEGER,
		BADQuality			INTEGER,
		TS					INTEGER,
		MD					INTEGER,
		sound				INTEGER,
		VFQ					INTEGER,
		VOSTFR				INTEGER,
		otherPb				TEXT,
		RIPQuality			TEXT,
		encQuality			TEXT,
		source				TEXT,
		fileSize			REAL,
		container			TEXT,
		bitRateT			INTEGER,
		codecV				TEXT,
		type3D 				TEXT,
		frameRate			TEXT,
		bitRateV			INTEGER,
		width				INTEGER,
		height				INTEGER,
		codecA				TEXT,
		audio				TEXT,
		sampling			INTEGER,
		bitRateA			INTEGER,
		subtitles			TEXT
  	)`

// sqlInsertMovies : requête SQL pour insérer un enregistrement dans la table.
const sqlInsertMovies = `INSERT INTO movies 
		( 
			title, 
			originalTitle, 
			picture,
			trailer,
			dateAjout, 
			dateModif, 
			dateSortie, 
			directors, 
			duration, 
			ageMini, 
			countries, 
			genres, 
			actors, 
			synopsis, 
			seen, 
			badMovie, 
			rating, 
			ratingPress, 
			comment, 
			control, 
			replace, 
			replaceInProgress, 
			deleted, 
			missing, 
			toReEncode, 
			timeLag, 
			BADQuality, 
			TS, 
			MD, 
			sound, 
			VFQ, 
			VOSTFR, 
			otherPb, 
			RIPQuality, 
			encQuality, 
			source,
			fileSize, 
			container, 
			bitRateT, 
			codecV, 
			type3D,
			frameRate, 
			bitRateV, 
			width, 
			height, 
			codecA, 
			audio, 
			sampling, 
			bitRateA, 
			subtitles 
		) 
		VALUES 
		(
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?,
			?
		)`

// sqlUpdateMovies : requête SQL pour mettre à jour un enregistrement dans la table.
const sqlUpdateMovies = `UPDATE movies SET 
		title = ?, 
		originalTitle = ?, 
		picture = ?,
		trailer = ?,
		dateAjout = ?, 
		dateModif = ?, 
		dateSortie = ?, 
		directors = ?, 
		duration = ?, 
		ageMini = ?, 
		countries = ?, 
		genres = ?, 
		actors = ?, 
		synopsis = ?, 
		seen = ?, 
		badMovie = ?, 
		rating = ?, 
		ratingPress = ?, 
		comment = ?, 
		control = ?, 
		replace = ?, 
		replaceInProgress = ?, 
		deleted = ?, 
		missing = ?, 
		toReEncode = ?, 
		timeLag = ?, 
		BADQuality = ?, 
		TS = ?, 
		MD = ?, 
		sound = ?, 
		VFQ = ?, 
		VOSTFR = ?, 
		otherPb = ?, 
		RIPQuality = ?, 
		encQuality = ?, 
		source = ?,
		fileSize = ?, 
		container = ?, 
		bitRateT = ?, 
		codecV = ?, 
		type3D = ?,
		frameRate = ?, 
		bitRateV = ?, 
		width = ?, 
		height = ?, 
		codecA = ?, 
		audio = ?, 
		sampling = ?, 
		bitRateA = ?, 
		subtitles = ? 
		WHERE id = ?
		`

// sqlGetMoviesByID : requête SQL pour restituer l'enregistrement de la table.
const sqlGetMoviesByID = `SELECT * FROM movies WHERE id = ?`

// sqlDeleteMoviesByID : requête SQL pour détruire un enregistrement dans la table.
const sqlDeleteMoviesByID = `DELETE FROM movies WHERE id = ?`

// sqlGetMoviesList : requête SQL pour restituer une vue (4 colonnes) de tous les enregistrements de la table.
const sqlGetMoviesList = `` +
	`SELECT` +
	` ` + `id,` +
	` ` + `title,` +
	` ` + `seen,` +
	` ` + `substr(dateAjout,7)||substr(dateAjout,4,2)||substr(dateAjout,1,2) AS date` +
	` FROM movies` +
	` ORDER BY title ASC`

// sqlGetContainerList : requête SQL pour restituer tous les containers de la table.
const sqlGetContainerList = `SELECT DISTINCT container FROM movies ORDER BY container ASC`

// sqlGetCodecVideoList : requête SQL pour restituer tous les codecs vidéo de la table.
const sqlGetCodecVideoList = `SELECT DISTINCT codecV FROM movies ORDER BY codecV ASC`

// sqlGetCodecAudioList : requête SQL pour restituer tous les codecs audio de la table.
const sqlGetCodecAudioList = `SELECT DISTINCT codecA FROM movies ORDER BY codecA ASC`

// sqlGetAudioList : requête SQL pour restituer tous les audio type de la table.
const sqlGetAudioList = `SELECT DISTINCT audio FROM movies ORDER BY audio ASC`

// sqlGetDirectorsList : requête SQL pour restituer tous les Directors de la table contenant 'saisie'.
func sqlGetDirectorsList(saisie string) string {
	return `` +
		`SELECT DISTINCT directors ` +
		`FROM movies ` +
		`WHERE directors LIKE '%` + saisie + `%'`
}

// sqlGetCountriesList : requête SQL pour restituer tous les Countries de la table.
func sqlGetCountriesList(saisie string) string {
	return `` +
		`SELECT DISTINCT countries ` +
		`FROM movies ` +
		`WHERE countries LIKE '%` + saisie + `%'`
}

// sqlGetGenresList : requête SQL pour restituer tous les Genres de la table.
func sqlGetGenresList(saisie string) string {
	return `` +
		`SELECT DISTINCT genres ` +
		`FROM movies ` +
		`WHERE genres LIKE '%` + saisie + `%'`
}

// sqlGetActorsList : requête SQL pour restituer tous les Actors de la table.
func sqlGetActorsList(saisie string) string {
	return `` +
		`SELECT DISTINCT actors ` +
		`FROM movies ` +
		`WHERE actors LIKE '%` + saisie + `%'`
}

// sqlGetSubtitleList : requête SQL pour restituer tous les subtitles type de la table.
const sqlGetSubtitleList = `SELECT DISTINCT subtitles FROM movies WHERE subtitles != "" ORDER BY subtitles ASC`
