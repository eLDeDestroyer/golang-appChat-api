CREATE TABLE users
(
    id INT NOT NULL AUTO_INCREMENT,
    unique_number INT UNIQUE,
    name VARCHAR(255) NOT NULL,
    username varchar(255) unique,
    password varchar(255),
    last_login date,
    primary key (id)
) ENGINE = InnoDB