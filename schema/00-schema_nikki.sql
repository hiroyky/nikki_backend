-- MySQL Script generated by MySQL Workbench
-- Fri May 29 23:29:41 2020
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema nikki_db
-- -----------------------------------------------------

-- -----------------------------------------------------
-- Schema nikki_db
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `nikki_db` DEFAULT CHARACTER SET utf8 ;
USE `nikki_db` ;

-- -----------------------------------------------------
-- Table `nikki_db`.`articles`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `nikki_db`.`articles` ;

CREATE TABLE IF NOT EXISTS `nikki_db`.`articles` (
  `article_id` INT NOT NULL AUTO_INCREMENT,
  `title` TEXT NOT NULL,
  `body` TEXT NOT NULL,
  `description` TEXT NOT NULL,
  `publish_status` INT NOT NULL,
  `thumbnail_image` TEXT NOT NULL,
  `posted_at` DATETIME NOT NULL,
  `created_at` TIMESTAMP NOT NULL,
  `updated_at` TIMESTAMP NOT NULL,
  PRIMARY KEY (`article_id`))
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
