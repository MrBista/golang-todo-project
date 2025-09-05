create datbase if not exists todo_app;
use todo_app;

create table if not exists users (
    id int primary key auto_increment,
    username varchar(255) not null,
    email varchar(255) not null,
    password varchar(255) not null,
    full_name varchar(255) not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    unique (email),
    unique (username)
);

create table if not exists user_otp (
    id int primary key auto_increment,
    user_id int not null,
    otp_code varchar(6) not null,
    expires_at timestamp not null,
    created_at timestamp default current_timestamp,
    foreign key (user_id) references users(id) on delete cascade
);

create table if not exists todos (
    id int primary key auto_increment,
    user_id int not null,
    title varchar(255) not null,
    description text,
    status int comment '0: pending, 1: in-progress, 2: complete, 3: remaining ' default 0,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp on update current_timestamp,
    foreign key (user_id) references users(id) on delete cascade
);