CREATE DATABASE harpun;
USE harpun;

# Dump of table stocks
# ------------------------------------------------------------

CREATE TABLE `stocks` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `avanzaId` int(11),
  `ticker` text UNIQUE,
  `name` text,
  PRIMARY KEY (`id`),
  KEY `avanzaId` (`avanzaId`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

# Dump of table trades
# ------------------------------------------------------------

CREATE TABLE `trades` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `ticker` text,
  `buyer` text,
  `seller` text,
  `amount`int(11),
  `price` text,
  `time` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
