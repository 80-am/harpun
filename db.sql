CREATE DATABASE harpun;
USE harpun;

# Dump of table stocks
# ------------------------------------------------------------

CREATE TABLE `stocks` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `avanzaId` int(11),
  `ticker` text UNIQUE,
  `name` text,
  `avgTradeAmount` int(11),
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

# Dump of table alerts
# ------------------------------------------------------------

CREATE TABLE `alerts` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `ticker` text,
  `amount`int(11),
  `price` text,
  `totalPrice` text,
  `time` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;