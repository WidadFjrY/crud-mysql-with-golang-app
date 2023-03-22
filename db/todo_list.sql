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

INSERT INTO
    todo_list (todo, description)
VALUES (
        'Belajar Golang',
        'Golang Databases'
    );

INSERT INTO
    todo_list (todo, description)
VALUES (
        'Belajar Java',
        'Java Dasar, OOP'
    ), ('Belajar PHP', 'Koneksi.php'), (
        'Belajar Python',
        'Machine Learning Dasar, AI'
    ), ('Belajar HTTP', 'Restful API');

UPDATE todo_list
SET
    todo = 'Belajar Golang Updated',
    description = 'Golang Databases Updated'
WHERE id = 1;

SELECT
    id,
    todo,
    description,
    is_finished,
    created_at
FROM todo_list;

DELETE FROM todo_list WHERE id = 10;