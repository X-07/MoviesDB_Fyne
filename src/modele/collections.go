package modele

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DbCollec *sql.DB

type Collections struct {
	ID   int64
	Name string
	Type string
}

func OpenDBCollec(dbName string) {
	DbCollec, err = sql.Open("sqlite3", dbName)
	if err != nil {
		panic(err.Error())
	}
}

func CreateCollecTable() {
	// In the definition below, please note that BLOB is the only type we can use in sqlite for storing JSON.
	_, err = DbCollec.Exec(reqCreateCollec)
	if err != nil {
		panic(err.Error())
	}
}

// InsertCollec : insertion d'un enregistrement dans la table.
func InsertCollec(collec *Collections) {
	stmt, err := DbCollec.Prepare(sqlInsertCollec)
	if err != nil {
		panic(err.Error())
	}

	defer stmt.Close()
	res, err := stmt.Exec(
		collec.Name,
		collec.Type,
	)
	if err != nil {
		panic(err.Error())
	}
	id, err := res.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	collec.ID = id
}

// UpdateCollec : Mise à jour d'un enregistrement dans la table.
func UpdateCollec(collec *Collections) {
	stmt, err := DbCollec.Prepare(sqlUpdateCollec)
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	_, err = stmt.Exec(
		collec.Name,
		collec.Type,
		collec.ID,
	)
	if err != nil {
		panic(err.Error())
	}
}

// GetCollecByID : Restitue l'enregistrement de la table.
func GetCollecByID(id int64) Collections {
	rows, err := DbCollec.Query(sqlGetCollecByID, id)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	var collec Collections
	if rows.Next() {
		err := rows.Scan(
			&collec.ID,
			&collec.Name,
			&collec.Type,
		)
		if err != nil {
			panic(err.Error())
		}
	}
	return collec
}

// DeleteCollecByID : détruit un enregistrement, identifié par son ID, dans la table et retourne le nb d'enregistrement supprimé (en principe 1).
func DeleteCollecByID(id int64) int64 {
	stmt, err := DbCollec.Prepare(sqlDeleteCollecByID)
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

// GetCollecList : Restitue la liste de tous les enregistrements de la table.
func GetCollecList() []Collections {
	var collecs []Collections
	rows, err := DbCollec.Query(sqlGetCollecList)
	if err != nil {
		//panic(err.Error())
		fmt.Println(err.Error())
		return collecs
	}
	defer rows.Close()

	for rows.Next() {
		var collec Collections
		err = rows.Scan(
			&collec.ID,
			&collec.Name,
			&collec.Type,
		)
		if err != nil {
			panic(err.Error())
		}
		collecs = append(collecs, collec)
	}
	return collecs
}
