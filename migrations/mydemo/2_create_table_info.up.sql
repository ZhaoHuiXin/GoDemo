create table if not exists info (id int not null primary key auto_increment,
idcard varchar(20) default null unique,
age tinyint not null default 0,
sex tinyint not null default 1 comment "1 male, 0 female",
address varchar(32) default null,
phone bigint not null default 0);