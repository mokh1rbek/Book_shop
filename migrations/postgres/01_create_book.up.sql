
CREATE TABLE books (
    book_id UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    title VARCHAR NOT NULL,
    author VARCHAR(255),
    price NUMERIC NOT NULL,
    date DATE NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE users (
    user_id UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    first_name VARCHAR(45) NOT NULL,
    last_name VARCHAR(45) NOT NULL,
    phone_number VARCHAR(12) NOT NULL,
    balance NUMERIC NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE orders (
    order_id UUID NOT NULL,
    user_id UUID NOT NULL REFERENCES users(user_id),
    book_id UUID NOT NULL REFERENCES books(book_id),
    payed NUMERIC NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

SELECT 
    users.first_name || ' ' || users.last_name as fullname,
    SUM(orders.payed)
FROM
    orders
JOIN users ON orders.user_id = users.user_id
GROUP BY fullname;