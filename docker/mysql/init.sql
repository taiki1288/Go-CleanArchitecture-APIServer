CREATE DATABASE go_cleanarchitecture_apiserver;
USE go_cleanarchitecture_apiserver;

CREATE TABLE users (
    id         INT NOT NULL AUTO_INCREMENT,
    name       VARCHAR(256) NOT NULL,
    email      VARCHAR(256) NOT NULL,
    age        INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON CURRENT_TIMESTAMP,
    update_at  TIMESTAMP NULL ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY(id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

