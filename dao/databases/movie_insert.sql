
insert into MOVIES(name, date_of_release, rating, description) 
values('No One killed Jessica', '2011-01-07', 3.0, 'No One Killed Jessica is a 2011 Indian biographical thriller film starring Rani Mukerji and Vidya Balan, produced by UTV Spotboy and directed by Rajkumar Gupta. This film is based on the Jessica Lal murder case.');
insert into MOVIES(name, date_of_release, rating, description) 
values('Dhobi Ghat', '2011-01-21', 4.0, 'Dhobi Ghat (also known as Mumbai Diaries) is a 2011 Indian drama film directed by Kiran Rao in her directorial debut. The film was produced by Aamir Khan Productions, Reliance Entertainment, and Shree Ashtavinayak Cine Vision Ltd, and stars Prateik Babbar, Monica Dogra, Kriti Malhotra and Aamir Khan in the lead roles.[1][4][5] Gustavo Santaolalla was signed to compose the score and soundtrack of the film, which includes a song by Ryuichi Sakamoto.');
insert into MOVIES(name, date_of_release, rating, description) 
values('7 Khoon Maaf', '2011-02-18', 3.0, '7 Khoon Maaf (read as Saat Khoon Maaf, released internationally as Seven Sins Forgiven) is a 2011 Indian black comedy film directed, co-written and co-produced by Vishal Bhardwaj. The film stars Priyanka Chopra in the lead role, with Vivaan Shah, Irrfan Khan, Annu Kapoor, Neil Nitin Mukesh, John Abraham, Aleksandr Dyachenko, Naseeruddin Shah and Usha Uthup in supporting roles. The film tells the story of a femme fatale, Susanna Anna-Marie Johannes, an Anglo-Indian woman who murders her seven husbands in an unending quest for love.');


-- multiplexes

insert into Multiplex (name, screens, address)
    values ('City Pride Kothrud', 9, 'Kothrud, Pune');

insert into Multiplex (name, screens, address)
values ('City Pride Satara', 3, 'Satara Road, Pune');

insert into Multiplex (name, screens, address)
    values ('City Pride Abhiruchi', 6, 'Sinhagad Road, Pune');


insert into SHOW (movie_id, Multiplex_id, showdate, showtime, screen_no, movie_name, multiplex_name)
values (1,1,'2018-05-16','09:00:00', 1, 'No One killed Jessica', 'City Pride Kothrud');
insert into SHOW (movie_id, Multiplex_id, showdate, showtime, screen_no, movie_name, multiplex_name)
values (1,1,'2018-05-16','12:00:00', 1, 'No One killed Jessica', 'City Pride Kothrud');
insert into SHOW (movie_id, Multiplex_id, showdate, showtime, screen_no, movie_name, multiplex_name)
values (1,1,'2018-05-16','14:00:00', 1, 'No One killed Jessica', 'City Pride Kothrud');
insert into SHOW (movie_id, Multiplex_id, showdate, showtime, screen_no, movie_name, multiplex_name)
values (1,1,'2018-05-16','17:00:00', 1, 'No One killed Jessica', 'City Pride Kothrud');

insert into SHOW (movie_id, Multiplex_id, showdate, showtime, screen_no, movie_name, multiplex_name)
values (2,1,'2018-05-16','09:00:00', 2, 'Dhobi Ghat', 'City Pride Kothrud');
insert into SHOW (movie_id, Multiplex_id, showdate, showtime, screen_no, movie_name, multiplex_name)
values (2,1,'2018-05-16','12:00:00', 2, 'Dhobi Ghat', 'City Pride Kothrud');
insert into SHOW (movie_id, Multiplex_id, showdate, showtime, screen_no, movie_name, multiplex_name)
values (2,1,'2018-05-16','14:00:00', 2, 'Dhobi Ghat', 'City Pride Kothrud');
insert into SHOW (movie_id, Multiplex_id, showdate, showtime, screen_no, movie_name, multiplex_name)
values (2,1,'2018-05-16','17:00:00', 2, 'Dhobi Ghat', 'City Pride Kothrud');


insert into SHOW (movie_id, Multiplex_id, showdate, showtime, screen_no, movie_name, multiplex_name)
values (2,2,'2018-05-16','09:00:00', 2,'Dhobi Ghat','City Pride Satara');
insert into SHOW (movie_id, Multiplex_id,Dhobi Ghat showdate, showtime, screen_no)
values (2,2,'2018-05-16','12:00:00', 2,'Dhobi Ghat','City Pride Satara');
insert into SHOW (movie_id, Multiplex_id, showdate, showtime, screen_no, movie_name, multiplex_name)
values (2,2,'2018-05-16','14:00:00', 2,'Dhobi Ghat','City Pride Satara');
insert into SHOW (movie_id, Multiplex_id, showdate, showtime, screen_no, movie_name, multiplex_name)
values (2,2,'2018-05-16','17:00:00', 2,'Dhobi Ghat','City Pride Satara');

insert into SHOW (movie_id, Multiplex_id, showdate, showtime, screen_no, movie_name, multiplex_name)
values (3,2,'2018-05-16','09:00:00', 2,'7 Khoon Maaf','City Pride Satara');
insert into SHOW (movie_id, Multiplex_id, showdate, showtime, screen_no, movie_name, multiplex_name)
values (3,2,'2018-05-16','12:00:00', 2,'7 Khoon Maaf','City Pride Satara');
insert into SHOW (movie_id, Multiplex_id, showdate, showtime, screen_no, movie_name, multiplex_name)
values (3,2,'2018-05-16','14:00:00', 2,'7 Khoon Maaf','City Pride Satara');
insert into SHOW (movie_id, Multiplex_id, showdate, showtime, screen_no, movie_name, multiplex_name)
values (3,2,'2018-05-16','17:00:00', 2,'7 Khoon Maaf','City Pride Satara');


insert into SHOW (movie_id, Multiplex_id, showdate, showtime, screen_no, movie_name, multiplex_name)
values (2,3,'2018-05-16','09:00:00', 2,'Dhobi Ghat','City Pride Abhiruchi');
insert into SHOW (movie_id, Multiplex_id, showdate, showtime, screen_no, movie_name, multiplex_name)
values (2,3,'2018-05-16','12:00:00', 2,'Dhobi Ghat','City Pride Abhiruchi');
insert into SHOW (movie_id, Multiplex_id, showdate, showtime, screen_no, movie_name, multiplex_name)
values (2,3,'2018-05-16','14:00:00', 2,'Dhobi Ghat','City Pride Abhiruchi');
insert into SHOW (movie_id, Multiplex_id, showdate, showtime, screen_no, movie_name, multiplex_name)
values (2,3,'2018-05-16','17:00:00', 2,'Dhobi Ghat','City Pride Abhiruchi');


insert into SHOW (movie_id, Multiplex_id, showdate, showtime, screen_no, movie_name, multiplex_name)
values (3,3,'2018-05-16','09:00:00', 2,'7 Khoon Maaf','City Pride Abhiruchi');
insert into SHOW (movie_id, Multiplex_id, showdate, showtime, screen_no, movie_name, multiplex_name)
values (1,3,'2018-05-16','12:00:00', 2,'No One killed Jessica','City Pride Abhiruchi');
insert into SHOW (movie_id, Multiplex_id, showdate, showtime, screen_no, movie_name, multiplex_name)
values (2,3,'2018-05-16','14:00:00', 2,'Dhobi Ghat','City Pride Abhiruchi');
insert into SHOW (movie_id, Multiplex_id, showdate, showtime, screen_no, movie_name, multiplex_name)
values (3,3,'2018-05-16','17:00:00', 2,'7 Khoon Maaf','City Pride Abhiruchi');