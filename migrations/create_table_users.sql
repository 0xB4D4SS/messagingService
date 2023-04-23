create table users (
    id int auto_increment primary key,
    login varchar(30) not null unique,
    password varchar(255) not null,
    token varchar(255) null
)