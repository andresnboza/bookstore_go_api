CREATE SCHEMA `users_db` DEFAULT CHARACTER SET utf8 ;

CREATE DATABASE usersDB;

use usersDB;

CREATE TABLE `users_db`.`users` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `first_name` VARCHAR(45) NOT NULL,
  `last_name` VARCHAR(45) NOT NULL,
  `email` VARCHAR(45) NOT NULL,
  `date_created` VARCHAR(45) NOT NULL,
  `password` VARCHAR(45) NOT NULL,
  `status` tinyint(1) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `email_UNIQUE` (`email` ASC) VISIBLE);


-- INSERT INTO Users (id, first_name, last_name, email, date_created, 'status')
-- VALUES (1, 'Andres', 'Navarrete', 'andresnboza92@gmail.com', '5-5-2021', '')
