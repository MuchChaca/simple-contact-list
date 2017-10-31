-- Valentina Studio --
-- MySQL dump --
-- ---------------------------------------------------------


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
-- ---------------------------------------------------------


-- CREATE DATABASE "go_training_contactList" ---------------
CREATE DATABASE IF NOT EXISTS `go_training_contactList` CHARACTER SET utf8 COLLATE utf8_general_ci;
USE `go_training_contactList`;
-- ---------------------------------------------------------


-- CREATE TABLE "contacts" ---------------------------------
-- CREATE TABLE "contacts" -------------------------------------
CREATE TABLE `contacts` ( 
	`id` Int( 11 ) AUTO_INCREMENT NOT NULL,
	`lName` VarChar( 255 ) CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
	`fName` VarChar( 255 ) CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
	`email` VarChar( 255 ) CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
	`phone` VarChar( 18 ) CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
	`cPhone` VarChar( 18 ) CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
	PRIMARY KEY ( `id` ) )
CHARACTER SET = utf8
COLLATE = utf8_general_ci
ENGINE = InnoDB
AUTO_INCREMENT = 39;
-- -------------------------------------------------------------
-- ---------------------------------------------------------


-- Dump data of "contacts" ---------------------------------
INSERT INTO `contacts`(`id`,`lName`,`fName`,`email`,`phone`,`cPhone`) VALUES 
( '26', 'Williams', 'Nina', 'n.williams@gmail.com', '667766776677', '776677667766' ),
( '27', 'Scarlet', 'Erza', 'erza.scarlet@gmail.com', '1243245353', '2343234325' ),
( '28', 'Stark', 'Tony', 't.stark@stark.com', '444444444', '444444444' ),
( '29', 'Banner', 'Bruce', 'b.banner@hotmail.com', '1245673423', '1245673423' ),
( '30', 'Draagnir', 'Natsu', 'natsu.happy@fairytail.com', '1122443322', '3322442233' ),
( '35', 'Kazama', 'Jin', 'j.kazama@gmail.com', '2288339922', '2288339922' ),
( '36', 'Dragonborn', 'Dovahkin', 'fus@roh.da', 'yoleeeee', 'fus roh da' ),
( '37', 'Danvers', 'Carol', 'ms.marvel@stark.com', '3366774411', '7766443322' ),
( '38', 'Marvel', 'Blue', 'b.marvel@marvel.com', '2200993388', '3300994488' );
-- ---------------------------------------------------------


/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
-- ---------------------------------------------------------


