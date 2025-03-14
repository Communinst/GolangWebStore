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
    role_id int not null references roles(role_id)
);
create table wallets
(
    wallet_id serial primary key,
    user_id int not null references users(user_id),
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
    rating decimal(4, 1) not null check (rating <= 100.0)
);
-- create table games_genres
-- (
--     game_genre_id serial primary key,
--     game_id int not null references games(game_id),
--     genre_id int not null references genres(genre_id)
-- );
CREATE TABLE games_genres (
    game_genre_id SERIAL PRIMARY KEY,
    game_id INT NOT NULL REFERENCES games(game_id),
    genre_id INT NOT NULL REFERENCES genres(genre_id),
    count INT DEFAULT 0,
    UNIQUE (game_id, genre_id)
);
-- create table games_developed_by
-- (
--     game_developer_id serial primary key,
--     game_id int not null references games(game_id),
--     company_id int not null references companies(company_id),
--     good_percentage int not null check (good_percentage between 0 and 100),
--     good_percentage_latest int not null check (good_percentage_latest between 0 and 100)
-- );

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
    minutes_spent bigint not null check (minutes_spent >= 0),
    receipt_date timestamp not null
);
    -- create table achieved_by_users
    -- (
    --     user_achievement_id serial primary key,
    --     user_id int not null references users(user_id),
    --     game_id int not null references games(game_id)
    -- );
CREATE TABLE dumps (
    id SERIAL PRIMARY KEY,
    filename VARCHAR(255) NOT NULL,
    size BIGINT NOT NULL
);



CREATE OR REPLACE FUNCTION create_cart_for_new_user()
RETURNS TRIGGER AS $$
DECLARE 
    user_role TEXT;
BEGIN
    SELECT name INTO user_role
    FROM roles  
    WHERE role_id = NEW.role_id;
    IF user_role = 'User' THEN
        
        INSERT INTO carts (user_id)
        VALUES (NEW.user_id);

        
        INSERT INTO wallets (user_id, balance)
        VALUES (NEW.user_id, 0);
    END IF;
        
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER after_user_inserted
AFTER INSERT ON users
FOR EACH ROW
EXECUTE FUNCTION create_cart_for_new_user();


ALTER TABLE carts
ADD CONSTRAINT fk_user
FOREIGN KEY (user_id)
REFERENCES users(user_id)
ON DELETE CASCADE;

ALTER TABLE wallets
ADD CONSTRAINT fk_user_wallet
FOREIGN KEY (user_id)
REFERENCES users(user_id)
ON DELETE CASCADE;



CREATE OR REPLACE FUNCTION check_if_valid_count()
RETURNS TRIGGER AS $$
BEGIN
    IF NEW.count < 0 THEN
        NEW.count := 0;
    END IF;
    RETURN NEW;

END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER after_update_games_genres
AFTER UPDATE ON games_genres
FOR EACH ROW
EXECUTE FUNCTION check_if_valid_count();




CREATE OR REPLACE FUNCTION insert_dump(p_filePath VARCHAR)
    RETURNS VOID AS
$$
BEGIN
    INSERT INTO dumps (filePath)
    VALUES (p_filePath)
    ON CONFLICT (filePath) DO NOTHING;
END;
$$ LANGUAGE plpgsql;


CREATE OR REPLACE FUNCTION get_all_dumps()
    RETURNS TABLE
            (
                filePath VARCHAR
            )
AS
$$
BEGIN
    RETURN QUERY SELECT dumps.filePath FROM dumps;
END;
$$ LANGUAGE plpgsql;
