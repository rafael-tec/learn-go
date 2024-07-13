CREATE TABLE orders (
    id varchar(255) NOT NULL,
    price float NOT NULL,
    tax float NOT NULL,
    final_price float NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO orders (id, price, tax, final_price)
VALUES ('a3c7g', 999, 333, 11);

INSERT INTO orders (id, price, tax, final_price)
VALUES ('h9f4z', 777, 888, 33);

INSERT INTO orders (id, price, tax, final_price)
VALUES ('u6i5w', 111, 5555, 44);