SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL,ALLOW_INVALID_DATES';
-- -----------------------------------------------------
-- Schema messenger
-- -----------------------------------------------------
DROP SCHEMA IF EXISTS `keikibook` ;

-- -----------------------------------------------------
-- Schema messenger
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `keikibook` DEFAULT CHARACTER SET utf8 ;
SHOW WARNINGS;
USE `keikibook`;
-- -----------------------------------------------------
-- Table `users`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `users` ;

SHOW WARNINGS;
CREATE TABLE IF NOT EXISTS `users` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `phone` VARCHAR(16) NOT NULL,
  `email` VARCHAR(255) NOT NULL,
  `password` VARCHAR(40) NOT NULL,
  `first_name` VARCHAR(20) NOT NULL DEFAULT '',
  `last_name` VARCHAR(20) NOT NULL DEFAULT '',
  `middle_name` VARCHAR(20) NOT NULL DEFAULT '',
  `is_active` TINYINT(1) NOT NULL DEFAULT 0,
  `is_reported` TINYINT(1) NOT NULL DEFAULT 0,
  `is_blocked` TINYINT(1) NOT NULL DEFAULT 0,
  `preferences` TEXT NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`)
)
ENGINE = InnoDB;

SHOW WARNINGS;
CREATE UNIQUE INDEX `phone_UNIQUE` ON `users` (`phone` ASC);

SHOW WARNINGS;
CREATE UNIQUE INDEX `email_UNIQUE` ON `users` (`email` ASC);

SHOW WARNINGS;

-- -----------------------------------------------------
-- Table `contacts`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `contacts` ;

SHOW WARNINGS;
CREATE TABLE IF NOT EXISTS `contacts` (
  `id` INT NOT NULL COMMENT 'Sync the contacts to this table',
  `first_name` VARCHAR(20) NOT NULL DEFAULT '',
  `middle_name` VARCHAR(20) NOT NULL DEFAULT '',
  `last_name` VARCHAR(20) NOT NULL DEFAULT '',
  `phone` VARCHAR(14) NOT NULL,
  `email` VARCHAR(255) NOT NULL,
  `created_at` VARCHAR(45) NOT NULL,
  PRIMARY KEY (`id`)
)
ENGINE = InnoDB;

SHOW WARNINGS;

-- -----------------------------------------------------
-- Table `conversation`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `conversation` ;

SHOW WARNINGS;
CREATE TABLE IF NOT EXISTS `conversation` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(40) NOT NULL,
  `creator_id` INT NOT NULL,
  `channel_id` VARCHAR(45) NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  `deleted_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`)
)
ENGINE = InnoDB;

SHOW WARNINGS;

-- -----------------------------------------------------
-- Table `messages`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `messages` ;

SHOW WARNINGS;
CREATE TABLE IF NOT EXISTS `messages` (
  `id` INT NOT NULL,
  `guid` VARCHAR(100) NOT NULL,
  `conversation_id` INT NOT NULL,
  `sender_id` INT NOT NULL,
  `message_type` ENUM('text', 'image', 'video', 'audio') NOT NULL,
  `message` VARCHAR(255) NOT NULL DEFAULT '',
  `attachment_thumb_url` VARCHAR(255) NOT NULL DEFAULT '',
  `attachment_url` VARCHAR(255) NOT NULL DEFAULT '',
  `created_at` DATETIME NOT NULL,
  `deleted_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`)
)
ENGINE = InnoDB;

SHOW WARNINGS;

-- -----------------------------------------------------
-- Table `participants`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `participants` ;

SHOW WARNINGS;
CREATE TABLE IF NOT EXISTS `participants` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `conversation_id` INT NOT NULL,
  `users_id` INT NOT NULL,
  `type` ENUM('single', 'group') NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`)
)
ENGINE = InnoDB;

SHOW WARNINGS;

-- -----------------------------------------------------
-- Table `reports`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `reports` ;

SHOW WARNINGS;
CREATE TABLE IF NOT EXISTS `reports` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `users_id` INT NOT NULL,
  `participants_id` INT NOT NULL,
  `report_type` VARCHAR(45) NOT NULL,
  `notes` TEXT NOT NULL,
  `status` ENUM('pending', 'resolved') NOT NULL DEFAULT 'pending',
  `created_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`)
)
ENGINE = InnoDB;

SHOW WARNINGS;

-- -----------------------------------------------------
-- Table `block_list`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `block_list` ;

SHOW WARNINGS;
CREATE TABLE IF NOT EXISTS `block_list` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `users_id` INT NOT NULL,
  `participants_id` INT NOT NULL,
  `created_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`)
)
ENGINE = InnoDB;

SHOW WARNINGS;

-- -----------------------------------------------------
-- Table `deleted_conversations`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `deleted_conversations` ;

SHOW WARNINGS;
CREATE TABLE IF NOT EXISTS `deleted_conversations` (
  `id` INT NOT NULL,
  `conversation_id` INT NOT NULL,
  `users_id` INT NOT NULL,
  `created_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`)
)
ENGINE = InnoDB;

SHOW WARNINGS;

-- -----------------------------------------------------
-- Table `devices`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `devices` ;

SHOW WARNINGS;
CREATE TABLE IF NOT EXISTS `devices` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `users_id` INT NOT NULL,
  `device_id` VARCHAR(120) NOT NULL,
  `type` ENUM('APPLE') NOT NULL,
  `device_token` VARCHAR(120) NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`)
)
ENGINE = InnoDB;

SHOW WARNINGS;

-- -----------------------------------------------------
-- Table `access`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `access` ;

SHOW WARNINGS;
CREATE TABLE IF NOT EXISTS `access` (
  `id` INT NOT NULL,
  `users_id` INT NOT NULL,
  `devices_id` INT NOT NULL,
  `token` VARCHAR(60) NOT NULL,
  `created_at` DATETIME NOT NULL,
  `deleted_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`)
)
ENGINE = InnoDB;

SHOW WARNINGS;

-- -----------------------------------------------------
-- Table `deleted_messages`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `deleted_messages` ;

SHOW WARNINGS;
CREATE TABLE IF NOT EXISTS `deleted_messages` (
  `id` INT NOT NULL AUTO_INCREMENT,
  `messages_id` INT NOT NULL,
  `users_id` INT NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`)
)
ENGINE = InnoDB;

SHOW WARNINGS;

-- -----------------------------------------------------
-- Table `user_contact`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `user_contact` ;

SHOW WARNINGS;
CREATE TABLE IF NOT EXISTS `user_contact` (
  `user_id` INT NOT NULL,
  `contact_id` INT NOT NULL,
  `first_name` VARCHAR(45) NOT NULL DEFAULT '',
  `last_name` VARCHAR(45) NOT NULL DEFAULT '',
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`user_id`, `contact_id`)
)
ENGINE = InnoDB;

SHOW WARNINGS;

-- -----------------------------------------------------
-- Table `activities`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `activities` ;

SHOW WARNINGS;
CREATE TABLE IF NOT EXISTS `activities` (
  `id` INT NOT NULL,
  `activity_type` VARCHAR(45) NOT NULL,
  `activity_id` INT NOT NULL,
  `title` VARCHAR(45) NOT NULL,
  `detail` TEXT NOT NULL,
  `created_at` DATETIME NOT NULL,
  `updated_at` DATETIME NOT NULL,
  PRIMARY KEY (`id`)
)
ENGINE = InnoDB;

SHOW WARNINGS;

-- -----------------------------------------------------
-- Table `user_verification`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `user_verification` ;

SHOW WARNINGS;
CREATE TABLE IF NOT EXISTS `user_verification` (
  `users_id` INT NOT NULL,
  `verification_code` VARCHAR(45) NOT NULL,
  `created_at` VARCHAR(45) NOT NULL
)
ENGINE = InnoDB;

SHOW WARNINGS;

SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
