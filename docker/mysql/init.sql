CREATE DATABASE IF NOT EXISTS go_cleanarchitecture_apiserver;
USE go_cleanarchitecture_apiserver;

CREATE TABLE users (
    id         INT NOT NULL AUTO_INCREMENT,
    name       VARCHAR(256) NOT NULL,
    email      VARCHAR(256) NOT NULL,
    age        INT NOT NULL,
    PRIMARY KEY(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

