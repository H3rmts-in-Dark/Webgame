# created by PHP-Storm

-- Database
create table logs
(
    date datetime not null,
    id   int auto_increment
        primary key,
    constraint logs_id_uindex
        unique (id)
);


-- User 

create user server
    identified by 'server';

grant alter, delete, insert, select, show databases, show view, update on *.* to server;

-- 
