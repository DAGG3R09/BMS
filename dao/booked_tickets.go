package dao

import "log"

// BookedTicket : Booked_Tickets table schema
type BookedTicket struct {
	ID            int
	TicketID      int
	PaymentStatus int
	BookingUser   string
}

// GetAllBookedTickets : Returns all BookedTickets form BookedTicketsDB
func GetAllBookedTickets() []BookedTicket {
	var bookedTickets []BookedTicket

	rows, err := db.Query("Select * from BookedTicket")

	if err != nil {
		panic(err)
	}

	for rows.Next() {
		var bookedTicket BookedTicket

		err := rows.Scan(
			&bookedTicket.ID,
			&bookedTicket.TicketID,
			&bookedTicket.PaymentStatus,
			&bookedTicket.BookingUser)

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

	query := "Select * from users where booking_id=$1"

	err := db.QueryRow(query, TicketID).Scan(
		&bookedTicket.ID,
		&bookedTicket.TicketID,
		&bookedTicket.PaymentStatus,
		&bookedTicket.BookingUser)

	if err != nil {
		return BookedTicket{}
	}

	return bookedTicket
}

// GetBookedTicketByTicketID : Returns BookedTicket associated with TicketID
func GetBookedTicketByTicketID(TicketID int) BookedTicket {
	var bookedTicket BookedTicket

	query := "Select * from users where ticket_id=$1"

	err := db.QueryRow(query, TicketID).Scan(
		&bookedTicket.ID,
		&bookedTicket.TicketID,
		&bookedTicket.PaymentStatus,
		&bookedTicket.BookingUser)

	if err != nil {
		return BookedTicket{}
	}

	return bookedTicket
}

// CreateBookedTicket : Inserts Data into BookedTicket database
func CreateBookedTicket(ticketID int, paymentStatus int, bookingUser string) {

	query := `Insert into BookedTickets 
			(ticket_id, payment_status, booking_user) 
			values ($1, $2, $3)`

	err := db.QueryRow(query, ticketID, paymentStatus, bookingUser)

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
			&bookedTicket.BookingUser)

		if err != nil {
			panic(err)
		}
		bookedTickets = append(bookedTickets, bookedTicket)
	}

	return bookedTickets
}
