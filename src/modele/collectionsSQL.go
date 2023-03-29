package modele

// reqCreateCollec : requête SQL de création de la table
const reqCreateCollec = `CREATE TABLE IF NOT EXISTS collections (
		id            		INTEGER PRIMARY KEY,
		name         		TEXT,
		type		 		TEXT 
  	)`

// sqlInsertCollec : requête SQL pour insérer un enregistrement dans la table.
const sqlInsertCollec = `INSERT INTO collections
		( 
			name, 
			type 
		) 
		VALUES 
		(
			?,
			? 
		)`

// sqlUpdateCollec : requête SQL pour mettre à jour un enregistrement dans la table.
const sqlUpdateCollec = `UPDATE collections SET 
		name = ?, 
		type = ? 
		WHERE id = ?
		`

// sqlGetCollecByID : requête SQL pour restituer l'enregistrement de la table.
const sqlGetCollecByID = `SELECT * FROM collections WHERE id = ?`

// sqlDeleteCollecByID : requête SQL pour détruire un enregistrement dans la table.
const sqlDeleteCollecByID = `DELETE FROM collections WHERE id = ?`

// sqlGetCollecList : requête SQL pour restituer une liste de tous les enregistrements de la table.
const sqlGetCollecList = `` +
	`SELECT` +
	` ` + `id,` +
	` ` + `name,` +
	` ` + `type` +
	` FROM collections` +
	` ORDER BY id ASC`
