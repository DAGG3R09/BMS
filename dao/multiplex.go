package dao

import "log"

// Multiplex : The Muliplex Yable Schema
type Multiplex struct {
	ID      int
	Name    string
	Screens int
}

// GetAllMultiplex : Returns all Multiplex form MultiplexDB
func GetAllMultiplex() []Multiplex {
	var multiplexes []Multiplex

	rows, err := db.Query("Select * from Multiplex")

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var multiplex Multiplex

		err := rows.Scan(
			&multiplex.ID,
			&multiplex.Name,
			&multiplex.Screens)

		if err != nil {
			panic(err)
		}
		multiplexes = append(multiplexes, multiplex)
	}

	return multiplexes
}

// GetSingleMultiplex : Returns Multiplex associated with MultiplexID
func GetSingleMultiplex(multiplexID int) Multiplex {
	var multiplex Multiplex

	query := "Select * from users where Multiplex_id=$1"

	err := db.QueryRow(query, multiplexID).Scan(
		&multiplex.ID,
		&multiplex.Name,
		&multiplex.Screens)

	if err != nil {
		return Multiplex{}
	}

	return multiplex
}

// CreateMultiplex : Inserts Data into Multiplex database
func CreateMultiplex(name string, screens int) {

	query := "Insert into Multiplex (name, screens) values ($1, $2)"

	err := db.QueryRow(query, name, screens)

	if err != nil {
		log.Fatal("Error in inserting Multiplex", err)
	}
}
