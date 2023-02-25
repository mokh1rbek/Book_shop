
CREATE TABLE books (
    book_id UUID PRIMARY KEY NOT NULL,
    book_name VARCHAR(50) NOT NULL,
    author_name VARCHAR(50) NOT NULL,
    prise NUMERIC NOT NULL,
    date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE users (,
    user_id UUID PRIMARY KEY NOT NULL,
    first_name VARCHAR (50) NOT NULL,
    last_name VARCHAR (50) NOT NULL,
    phone_number VARCHAR(9) NOT NULL,
    balance NUMERIC DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE orders (
    order_id UUID PRIMARY KEY NOT NULL,
    user_id UUID NOT NULL REFERENCES users(user_id),
    book_id UUID NOT NULL REFERENCES books(book_id),
    payed NUMERIC NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL 
);
