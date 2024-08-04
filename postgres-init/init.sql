--- Users table -----
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) NOT NULL,
    phone_number VARCHAR(100),
    password VARCHAR(255) NOT NULL
);


