CREATE DATABASE TODOLIST;

CREATE TABLE TODOS (
    `TodoId` int AUTO_INCREMENT, 
    `Text` VARCHAR(255) NOT NULL DEFAULT "",
    `Done` BOOLEAN NOT NULL DEFAULT FALSE,
    PRIMARY KEY (`TodoId`));