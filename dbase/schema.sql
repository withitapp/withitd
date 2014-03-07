CREATE TABLE `users` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`created_at` DATETIME NOT NULL,
	`updated_at` DATETIME NOT NULL,
	`username` VARCHAR(99) NOT NULL,
	`email` VARCHAR(99) NOT NULL,
	`first_name` VARCHAR(99) NOT NULL,
	`last_name` VARCHAR(99) NOT NULL,
	`fb_id` VARCHAR(99) NOT NULL,
	`fb_token` VARCHAR(99) NOT NULL,
	`fb_synect_at` DATETIME NOT NULL,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ;

CREATE TABLE `polls` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`created_at` DATETIME NOT NULL,
	`updated_at` DATETIME NOT NULL,
	`title` VARCHAR(99) NOT NULL,
	`description` TEXT NOT NULL,
	`user_id` INT NOT NULL,
	`ends_at` DATETIME NOT NULL,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ;

CREATE TABLE `memberships` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`created_at` DATETIME NOT NULL,
	`updated_at` DATETIME NOT NULL,
	`user_id` INT NOT NULL,
	`poll_id` INT NOT NULL,
	`response` bit NOT NULL, 
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ;

CREATE TABLE `friendships` (
	`id` INT NOT NULL AUTO_INCREMENT,
	`created_at` DATETIME NOT NULL,
	`updated_at` DATETIME NOT NULL,
	`alpha_id` INT NOT NULL,
	`beta_id` INT NOT NULL,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 ;
