-- Create table for restaurant_types
CREATE TABLE restaurant_types (
    id INTEGER PRIMARY KEY,
    name VARCHAR(255)
);

-- Create table for restaurants
CREATE TABLE restaurants (
    id INTEGER PRIMARY KEY,
    name VARCHAR(255),
    address TEXT,
    restaurant_type_id INTEGER,
    FOREIGN KEY (restaurant_type_id) REFERENCES restaurant_types(id)
);

-- Create table for menu_types
CREATE TABLE menu_types (
    id INTEGER PRIMARY KEY,
    name VARCHAR(255)
);

-- Create table for menus
CREATE TABLE menus (
    id INTEGER PRIMARY KEY,
    menu_type_id INTEGER,
    restaurant_id INTEGER,
    name VARCHAR(255),
    description TEXT,
    price INTEGER,
    FOREIGN KEY (menu_type_id) REFERENCES menu_types(id),
    FOREIGN KEY (restaurant_id) REFERENCES restaurants(id)
);

-- Create table for users
CREATE TABLE users (
    id INTEGER PRIMARY KEY,
    username VARCHAR(255),
    email VARCHAR(255),
    password VARCHAR(255)
);

-- Create table for user_reviews
CREATE TABLE user_reviews (
    id INTEGER PRIMARY KEY,
    restaurant_id INTEGER,
    rating FLOAT,
    description TEXT,
    user_id INTEGER,
    FOREIGN KEY (restaurant_id) REFERENCES restaurants(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);