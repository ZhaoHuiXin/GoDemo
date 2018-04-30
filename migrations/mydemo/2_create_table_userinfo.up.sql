create table if not exists user (id int not null primary key auto_increment,
idcard varchar(20) default null unique,
age int not null default 0,
sex char(1) default null,
address varchar(32) default null,
phone int not null default 0);