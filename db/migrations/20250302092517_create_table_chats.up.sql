CREATE TABLE chats
(
    id INT NOT NULL AUTO_INCREMENT,
    text varchar(255),
    room_id INT NOT NULL,
    user_id INT NOT NULL,
    primary key (id),
    foreign key (room_id) references rooms(id) ON DELETE CASCADE,
    foreign key (user_id) references users(id) ON DELETE CASCADE
) ENGINE = InnoDB