drop table Users;
Create Table Users (
    user_id SERIAL PRIMARY KEY,
    email varchar(30) Unique,
    name varchar(30) NOT NULL,
    pwd varchar(100) NOT NULL,
    phone char(10) Unique
);
 
drop table Movies;
Create Table Movies (
    movie_id SERIAL PRIMARY KEY,
    name varchar(30) NOT NULL,
    date_of_release varchar NOT NULL,
    Rating float default 0.0
);

drop table Multiplex;
Create Table Multiplex(
    multiplex_id int PRIMARY KEY,
    name varchar(30) NOT NULL,
    screens int,
);

drop table Show;
Create Table Show (
    show_id int PRIMARY KEY,
    movie_id int REFERENCES Movie(movie_id),
    multiplex_id int REFERENCES Multiplex(multiplex_id),
    showtime time NOT NULL,
    screen_no int,
);


drop table Ticket;
Create Table Ticket (
    ticket_id int PRIMARY KEY,
    show_id int REFERENCES Show(show_id),
    price int
    booked boolean
);

drop table Booked_Tickets;
Create table Booked_Tickets(
    booking_id int PRIMARY KEY,
    ticket_id int,
    Payment_status int,
    booking_user varchar(30) REFERENCES Users(email)
);


drop table Session;
Create Table Session(
    user_id int REFERENCES Users(user_id)
    token varchar(50) Unique
);