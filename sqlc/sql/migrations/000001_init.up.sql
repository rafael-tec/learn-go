CREATE TABLE categories (
    id varchar(36) NOT NULL PRIMARY KEY,
    name text NOT NULL,
    description text
);

CREATE TABLE courses (
    id varchar(36) NOT NULL PRIMARY KEY,
    category_id varchar(36) NOT NULL,
    name text NOT NULL,
    description text,
    price decimal(10,2) NOT NULL,
    FOREIGN KEY (category_id) REFERENCES categories(id)
);

INSERT INTO categories (
    id,
    name,
    description
) VALUES (
    '5314ff79-d20b-46e9-a9cf',
    'Technology',
    'Data Science, Machine Learning, Software Development'
);

INSERT INTO courses (
    id,
    category_id,
    name,
    description,
    price
) VALUES (
    'bc9ec010-68b8-49ed-a431',
    '5314ff79-d20b-46e9-a9cf',
    'Basic to advanced with Golang',
    'Learn the basic concepts and main features of the Golang language',
    99.99
);