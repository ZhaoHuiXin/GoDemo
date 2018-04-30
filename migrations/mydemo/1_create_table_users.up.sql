create table if not exists users (id int not null primary key auto_increment,
name varchar(16) default null unique,
password varchar(32) default null);