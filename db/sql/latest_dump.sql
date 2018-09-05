-- MySQL dump 10.13  Distrib 5.7.21, for Linux (x86_64)
--
-- Host: 127.0.0.1    Database: homework
-- ------------------------------------------------------
-- Server version	5.7.21

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `company`
--
DROP TABLE IF EXISTS `company`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `company` (
  `auto_id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'primary key',
  `user` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `money` double(255,2) DEFAULT NULL,
  `time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`auto_id`),
  KEY `user` (`user`)
) ENGINE=InnoDB AUTO_INCREMENT=46 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `company`
--

LOCK TABLES `company` WRITE;
/*!40000 ALTER TABLE `company` DISABLE KEYS */;
INSERT INTO `company` VALUES (19,'vimi',50.00,'2018-03-25 05:54:56'),(20,'vimi',32.00,'2018-03-25 05:56:18'),(21,'vimi',45.50,'2018-03-25 14:37:38'),(22,'aki',50.00,'2018-03-25 14:42:53'),(23,'aki',50.00,'2018-03-25 14:45:10'),(24,'vimi',50.00,'2018-03-25 14:50:39'),(25,'vimi',50.00,'2018-03-25 14:51:31'),(26,'vimi',50.00,'2018-03-25 14:53:26'),(27,'vimi',50.00,'2018-03-25 15:01:08'),(28,'test',50.00,'2018-03-25 15:02:21'),(29,'jjj',50.00,'2018-03-25 15:37:52'),(30,'sadf',50.00,'2018-03-25 16:30:50'),(31,'23',32.00,'2018-03-25 16:31:15'),(32,'22',50.00,'2018-03-25 16:31:53'),(33,'vimi',50.00,'2018-03-26 05:22:47'),(34,'jack',50.00,'2018-03-26 05:34:00'),(35,'vimi',50.00,'2018-03-26 05:41:30'),(36,'vimi',50.00,'2018-03-26 05:45:22'),(37,'asdf',0.50,'2018-03-26 05:45:29'),(38,'vimi',34.00,'2018-03-26 05:50:45'),(39,'asdf',50.00,'2018-03-26 05:52:13'),(40,'testq',50.00,'2018-03-26 05:53:04'),(41,'vimi',23.00,'2018-03-26 06:06:00'),(42,'vimi',2.00,'2018-03-26 06:06:35'),(43,'q',2.00,'2018-03-26 06:11:02'),(44,'vpn',50.00,'2018-03-26 07:28:05'),(45,'vpn',50.00,'2018-03-26 07:28:13');
/*!40000 ALTER TABLE `company` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `self`
--

DROP TABLE IF EXISTS `self`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `self` (
  `auto_id` int(11) NOT NULL AUTO_INCREMENT COMMENT 'autoIncreasment\r\n',
  `user` varchar(255) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
  `money` double(255,2) DEFAULT NULL,
  `time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`auto_id`)
) ENGINE=InnoDB AUTO_INCREMENT=23 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `self`
--

LOCK TABLES `self` WRITE;
/*!40000 ALTER TABLE `self` DISABLE KEYS */;

/*!40000 ALTER TABLE `self` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-03-26 15:33:38
