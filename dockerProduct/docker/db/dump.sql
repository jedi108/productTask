#CREATE DATABASE `ShopItems` CHARACTER SET utf8 COLLATE utf8_general_ci;
CREATE TABLE `Product` (
  `id` int NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL,
  `image` varchar(255) NOT NULL,
  `url` varchar(255) NOT NULL,
  `price` DECIMAL(12, 2) NOT NULL,
  `aviable` int(1) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;