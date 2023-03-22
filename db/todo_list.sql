-- Active: 1678672883589@@127.0.0.1@3306@todo_list

CREATE TABLE
    todo_list (
        id INT NOT NULL AUTO_INCREMENT,
        todo VARCHAR(100) NOT NULL,
        description TEXT,
        is_finished BOOLEAN NOT NULL DEFAULT false,
        created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY (id)
    ) ENGINE = InnoDB;

DESC todo_list;

SELECT * FROM todo_list;
