create table messages (
    id int auto_increment primary key,
    user_id int not null,
    data varchar(255) not null
);