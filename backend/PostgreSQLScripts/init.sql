create table roles
(
    role_id serial primary key,
    name varchar(31) not null unique,
    description varchar(127) not null,
    significance_order int not null
);
create table users
(
    user_id serial primary key,
    login varchar(63) not null unique,
    password varchar(63) NOT NULL,
    nickname varchar(63) not null,
    email varchar(127) not null unique,
    sign_up_date timestamp not null,
);
create table user_roles
(
    user_role_id serial not null primary key,
    user_id int not null references users(user_id),
    role_id int not null references roles(role_id)
);
create table wallets
(
    user_id int primary key not null references users(user_id),
    balance bigint not null check (balance >= 0)
);



create table companies
(
    company_id serial primary key,
    name varchar(127) not null
);

create table genres
(
    genre_id serial primary key,
    name varchar(31) not null unique,
    description varchar(63) not null unique
);

create table games 
(
    game_id serial primary key,
    publisher_id int not null references companies(company_id),
    name varchar(63) not null unique,
    description text not null,
    price int not null check (price >= 0),
    release_date timestamp not null,
);
create table games_genres
(
    game_genre_id serial primary key,
    game_id int not null references games(game_id),
    genre_id int not null references genres(genre_id)
);
create table games_developed_by
(
    game_developer_id serial primary key,
    game_id int not null references games(game_id),
    company_id int not null references companies(company_id)
    good_percentage int not null check (good_percentage between 0 and 100),
    good_percentage_latest int not null check (good_percentage_latest between 0 and 100)
);

create table discounts
(
    discount_id serial not null primary key,
    game_id int not null references games(game_id),
    discount_value int not null check (discount_value between 0 and 100),
    start_date timestamp not null,
    cease_date timestamp not null
);
-- create table achievements
-- (
--     achievement_id serial primary key,
--     name varchar(63) not null,
--     description varchar(255) not null,
--     game_id int not null references games(game_id)
--     frequency_percentage decimal(4, 1) not null check (frequency <= 100.0)
-- );


create table reviews
(
    review_id serial primary key,
    recommended boolean not null,
    message text not null,
    user_id int not null references users(user_id),
    game_id int not null references games(game_id),
    date timestamp not null
);
-- create table wishings
-- (
--     wishing_id serial primary key,
--     user_id int not null references users(user_id),
--     game_id int not null references games(game_id)
-- );
create table carts
(
    cart_id serial primary key,
    user_id int not null references users(user_id)
);
create table cart_games 
(
    cart_games_id serial primary key,
    cart_id int not null references carts(cart_id),
    game_id int not null references games(game_id)
);
create table ownerships
(
    ownership_id serial primary key,
    user_id int not null references users(user_id),
    game_id int not null references games(game_id),
    minutes_spent bigint not null check (time_spent >= 0),
    receipt_date timestamp not null
);
    -- create table achieved_by_users
    -- (
    --     user_achievement_id serial primary key,
    --     user_id int not null references users(user_id),
    --     game_id int not null references games(game_id)
    -- );