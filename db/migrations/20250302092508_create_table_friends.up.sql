CREATE TABLE friends
(
    id int NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    friend int not null,
    user_id int not null,
    room_id int not null ,
    primary key (id),
    foreign key (friend) references users(id) ON DELETE CASCADE,
    foreign key (user_id) references users(id) ON DELETE CASCADE,
    foreign key (room_id) references rooms(id) ON DELETE CASCADE
) ENGINE = InnoDB