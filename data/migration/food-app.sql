SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for Food
-- ----------------------------
DROP TABLE IF EXISTS `foods`;
CREATE TABLE `foods` (
                         `id` int(255) NOT NULL AUTO_INCREMENT,
                         `user_id` int(11) NOT NULL,
                         `title` varchar(255) DEFAULT NULL,
                         `description` varchar(500) DEFAULT NULL,
                         `food_image` text,
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- ----------------------------
-- Table structure for User
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
                         `id` int(11) NOT NULL AUTO_INCREMENT,
                         `first_name` varchar(255) DEFAULT NULL,
                         `last_name` varchar(255) DEFAULT NULL,
                         `email` varchar(255) DEFAULT NULL,
                         `password` varchar(255) DEFAULT NULL,
                         PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

SET FOREIGN_KEY_CHECKS = 1;
