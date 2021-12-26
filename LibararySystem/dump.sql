-- MySQL dump 10.13  Distrib 8.0.26, for macos11 (arm64)
--
-- Host: 127.0.0.1    Database: LibraryManagement
-- ------------------------------------------------------
-- Server version	8.0.26

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `book`
--

DROP TABLE IF EXISTS `book`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `book` (
  `book_id` int NOT NULL AUTO_INCREMENT,
  `book_name` varchar(20) NOT NULL,
  `book_author` varchar(20) NOT NULL,
  `book_pub` varchar(20) NOT NULL,
  `book_num` int NOT NULL,
  `book_record` datetime DEFAULT NULL,
  PRIMARY KEY (`book_id`),
  UNIQUE KEY `book_book_id_uindex` (`book_id`)
) ENGINE=InnoDB AUTO_INCREMENT=19 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `book`
--

LOCK TABLES `book` WRITE;
/*!40000 ALTER TABLE `book` DISABLE KEYS */;
INSERT INTO `book` VALUES (1,'红楼梦','曹雪芹','中国人民出版社',1,'2021-12-20 18:49:15'),(2,'西游记','吴承恩','中国人民出版社',1,'2021-12-20 18:49:53'),(3,'水浒传','施耐庵','清华大学出版社',1,'2021-12-22 09:10:55'),(4,'三国传','罗贯中','中国人民出版社',1,'2021-12-23 22:37:02');
/*!40000 ALTER TABLE `book` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `borrow_return_tab`
--

DROP TABLE IF EXISTS `borrow_return_tab`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `borrow_return_tab` (
  `br_id` int NOT NULL AUTO_INCREMENT,
  `sid` int NOT NULL,
  `bid` int DEFAULT NULL,
  `borrow_date` datetime DEFAULT NULL,
  `agreed_return_date` datetime DEFAULT NULL,
  `return_date` datetime DEFAULT NULL,
  PRIMARY KEY (`br_id`),
  UNIQUE KEY `borrow_return_tab_br_id_uindex` (`br_id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `borrow_return_tab`
--

LOCK TABLES `borrow_return_tab` WRITE;
/*!40000 ALTER TABLE `borrow_return_tab` DISABLE KEYS */;
/*!40000 ALTER TABLE `borrow_return_tab` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `manager`
--

DROP TABLE IF EXISTS `manager`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `manager` (
  `mana_id` int NOT NULL AUTO_INCREMENT,
  `mana_name` varchar(20) NOT NULL,
  `mana_phone` varchar(20) NOT NULL,
  `mana_acc` varchar(20) NOT NULL,
  `mana_passwd` varchar(20) NOT NULL,
  PRIMARY KEY (`mana_id`),
  UNIQUE KEY `manager_mana_id_uindex` (`mana_id`),
  UNIQUE KEY `manager_mana_acc_uindex` (`mana_acc`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `manager`
--

LOCK TABLES `manager` WRITE;
/*!40000 ALTER TABLE `manager` DISABLE KEYS */;
INSERT INTO `manager` VALUES (1,'超级管理员','131000000000','admin','admin');
/*!40000 ALTER TABLE `manager` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `student`
--

DROP TABLE IF EXISTS `student`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `student` (
  `stu_id` int NOT NULL AUTO_INCREMENT,
  `stu_name` varchar(20) NOT NULL,
  `stu_gender` varchar(20) NOT NULL,
  `stu_age` varchar(20) NOT NULL,
  `stu_major` varchar(20) NOT NULL,
  `stu_able` int NOT NULL DEFAULT '1',
  `stu_acc` varchar(20) NOT NULL,
  `stu_passwd` varchar(20) NOT NULL,
  PRIMARY KEY (`stu_id`),
  UNIQUE KEY `student_stu_id_uindex` (`stu_id`),
  UNIQUE KEY `student_stu_acc_uindex` (`stu_acc`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `student`
--

LOCK TABLES `student` WRITE;
/*!40000 ALTER TABLE `student` DISABLE KEYS */;
INSERT INTO `student` VALUES (1,'小明','男','19','经济管理',1,'1001','1001'),(2,'小刚','女','19','区块链技术应用',1,'1002','1002'),(3,'张三','男','19','软件工程',1,'1003','1003'),(4,'李四','男','22','新媒体',1,'1004','1004'),(5,'王五','男','19','区块链技术',1,'1005','1005');
/*!40000 ALTER TABLE `student` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-12-25  0:18:03
