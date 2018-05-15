package dao

import "log"

// BookedTicket : Booked_Tickets table schema
type BookedTicket struct {
	ID            int
	TicketID      int
	PaymentStatus int
	BookingUser   string
	MovieName     string `json:"movie_name"`
}

// GetAllBookedTickets : Returns all BookedTickets form BookedTicketsDB
func GetAllBookedTickets() []BookedTicket {
	var bookedTickets []BookedTicket

	rows, err := db.Query("Select * from BookedTickets")

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var bookedTicket BookedTicket

		err := rows.Scan(
			&bookedTicket.ID,
			&bookedTicket.TicketID,
			&bookedTicket.PaymentStatus,
			&bookedTicket.BookingUser,
			&bookedTicket.MovieName)

		if err != nil {
			panic(err)
		}
		bookedTickets = append(bookedTickets, bookedTicket)
	}

	return bookedTickets
}

// GetBookedTicketByBookingID : Returns BookedTicket associated with TicketID
func GetBookedTicketByBookingID(TicketID int) BookedTicket {
	var bookedTicket BookedTicket

	query := "Select * from BookedTickets where booking_id=$1"

	err := db.QueryRow(query, TicketID).Scan(
		&bookedTicket.ID,
		&bookedTicket.TicketID,
		&bookedTicket.PaymentStatus,
		&bookedTicket.BookingUser,
		&bookedTicket.MovieName)

	if err != nil {
		return BookedTicket{}
	}

	return bookedTicket
}

// GetBookedTicketByUserID : Returns BookedTicket associated with TicketID
func GetBookedTicketByUserID(UserID int) []BookedTicket {
	var bookedTickets []BookedTicket

	query := "Select * from BookedTickets where booking_user=$1"

	rows, err := db.Query(query, UserID)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var bt BookedTicket
		err := rows.Scan(
			&bt.ID,
			&bt.TicketID,
			&bt.PaymentStatus,
			&bt.BookingUser,
			&bt.MovieName)

		if err != nil {
			panic(err)
		}

		bookedTickets = append(bookedTickets, bt)
	}

	return bookedTickets
}

// GetBookedTicketByTicketID : Returns BookedTicket associated with TicketID
func GetBookedTicketByTicketID(TicketID int) BookedTicket {
	var bookedTicket BookedTicket

	query := "Select * from BookedTickets where ticket_id=$1"

	err := db.QueryRow(query, TicketID).Scan(
		&bookedTicket.ID,
		&bookedTicket.TicketID,
		&bookedTicket.PaymentStatus,
		&bookedTicket.BookingUser,
		&bookedTicket.MovieName)

	if err != nil {
		return BookedTicket{}
	}

	return bookedTicket
}

// CreateBookedTicket : Inserts Data into BookedTicket database
func CreateBookedTicket(ticketID int, paymentStatus int,
	bookingUser int, movieName string) {

	query := `Insert into BookedTickets 
			(payment_status, booking_user, movie_name) 
			values ($1, $2, $3)`

	_, err := db.Exec(query, paymentStatus, bookingUser, movieName)

	if err != nil {
		log.Fatal("Error in inserting BookedTicket", err)
	}
}

// GetAllBookedTicketsOfShow : Returns All BookedTickets of show associated with showID
func GetAllBookedTicketsOfShow(showID int) []BookedTicket {
	var bookedTickets []BookedTicket

	rows, err := db.Query(
		`Select * from BookedTicket where ticket_id in
			(Select ticket_id where show_id=$1)`, showID)

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var bookedTicket BookedTicket

		err := rows.Scan(
			&bookedTicket.ID,
			&bookedTicket.TicketID,
			&bookedTicket.PaymentStatus,
			&bookedTicket.BookingUser,
			&bookedTicket.MovieName)

		if err != nil {
			panic(err)
		}
		bookedTickets = append(bookedTickets, bookedTicket)
	}

	return bookedTickets
}
