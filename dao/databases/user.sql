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
    date_of_release date NOT NULL,
    Rating float default 0.0,
    description TEXT
);

drop table Multiplex;
Create Table Multiplex(
    multiplex_id SERIAL PRIMARY KEY,
    name varchar(30) Unique,
    screens int,
    address TEXT
);

drop table Show;
Create Table Show (
    show_id SERIAL PRIMARY KEY,
    movie_id int REFERENCES Movies(movie_id),
    multiplex_id int REFERENCES Multiplex(multiplex_id),
    showdate date NOT NULL,
    showtime time NOT NULL,
    screen_no int,
    movie_name varchar,
    multiplex_name varchar    
);


drop table Ticket;
Create Table Ticket (
    ticket_id SERIAL PRIMARY KEY,
    show_id int REFERENCES Show(show_id),
    price int
    booked boolean
);

drop table BookedTickets;
Create table BookedTickets(
    booking_id SERIAL PRIMARY KEY,
    ticket_id Serial,
    Payment_status int,
    booking_user int REFERENCES Users(user_id),
    movie_name varchar
);


drop table Session;
Create Table Session(
    user_id int REFERENCES Users(user_id),
    token char(36) Unique,
    login_time TIMESTAMPTZ
);