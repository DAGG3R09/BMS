package dao

import "log"

// Ticket : Ticket Table Schema
type Ticket struct {
	ID     int
	ShowID int
	Price  int
	Booked bool
}

// GetAllTickets : Returns all Tickets form TicketsDB
func GetAllTickets() []Ticket {
	var tickets []Ticket

	rows, err := db.Query("Select * from Ticket")

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var ticket Ticket

		err := rows.Scan(
			&ticket.ID,
			&ticket.ShowID,
			&ticket.Price,
			&ticket.Booked)

		if err != nil {
			panic(err)
		}
		tickets = append(tickets, ticket)
	}

	return tickets
}

// GetSingleTicket : Returns Ticket associated with TicketID
func GetSingleTicket(TicketID int) Ticket {
	var ticket Ticket

	query := "Select * from users where Ticket_id=$1"

	err := db.QueryRow(query, TicketID).Scan(
		&ticket.ID,
		&ticket.ShowID,
		&ticket.Price,
		&ticket.Booked)

	if err != nil {
		return Ticket{}
	}

	return ticket
}

// CreateTicket : Inserts Data into Ticket database
func CreateTicket(ShowID int, price int, booked bool) {

	query := `Insert into Tickets 
			(show_id, price, booked) 
			values ($1, $2, $3)`

	err := db.QueryRow(query, ShowID, price, booked)

	if err != nil {
		log.Fatal("Error in inserting Ticket", err)
	}
}

// GetAllTicketsOfShow : Returns All tickets of show associated with showID
func GetAllTicketsOfShow(showID int) []Ticket {
	var tickets []Ticket

	rows, err := db.Query("Select * from Ticket where show_id=$1")
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var ticket Ticket

		err := rows.Scan(
			&ticket.ID,
			&ticket.ShowID,
			&ticket.Price,
			&ticket.Booked)

		if err != nil {
			panic(err)
		}
		tickets = append(tickets, ticket)
	}

	return tickets
}
