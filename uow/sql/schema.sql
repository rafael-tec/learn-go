CREATE TABLE IF NOT EXISTS categories (
    id int PRIMARY KEY AUTO_INCREMENT,
    name text NOT NULL
);

CREATE TABLE IF NOT EXISTS courses (
    id int PRIMARY KEY AUTO_INCREMENT,
    category_id varchar(36) NOT NULL,
    name text NOT NULL,
    category_id int NOT NULL,
    FOREIGN KEY (category_id) REFERENCES categories(id)
);