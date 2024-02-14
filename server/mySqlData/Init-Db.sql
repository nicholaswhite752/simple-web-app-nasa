CREATE DATABASE testGoDb;

USE testGoDb;

CREATE TABLE `NasaData` (
  `id` varchar(45) NOT NULL,
  `name` varchar(45) DEFAULT NULL,
  `nameType` varchar(45) DEFAULT NULL,
  `recclass` varchar(45) DEFAULT NULL,
  `mass` int DEFAULT NULL,
  `fall` varchar(45) DEFAULT NULL,
  `year` datetime DEFAULT NULL,
  `reclat` varchar(45) DEFAULT NULL,
  `reclong` varchar(45) DEFAULT NULL,
  `geolocation` blob,
  PRIMARY KEY (`id`)
);

GRANT ALL PRIVILEGES ON testGoDb.* TO 'testUser'@'%';