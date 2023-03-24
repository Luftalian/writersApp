-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: mariadb:3306
-- Generation Time: Mar 20, 2023 at 01:52 PM
-- Server version: 10.6.4-MariaDB-1:10.6.4+maria~focal
-- PHP Version: 8.0.19


CREATE DATABASE `WRAPP` DEFAULT CHARACTER SET utf8mb4;

USE `WRAPP`;

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `WRAPP`
--

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `user_id` char(36) DEFAULT NULL,
  `name` varchar(300) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Table structure for table `texts`
--

CREATE TABLE `texts` (
  `text_id` char(36) DEFAULT NULL,
  `title` varchar(300) DEFAULT NULL,
  `content` varchar(30000) DEFAULT NULL,
  `user_name` varchar(300) DEFAULT NULL,
  `user_id` char(36) DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `changed_at` timestamp NULL DEFAULT NULL,
  `good_count` int(11) UNSIGNED NOT NULL,
  `bad_count` int(11) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Table structure for table `tags`
--

CREATE TABLE `tags` (
  `tag_id` char(36) DEFAULT NULL,
  `tag_name` varchar(300) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Table structure for table `userCreates`
--

CREATE TABLE `userCreates` (
  `user_id` char(36) DEFAULT NULL,
  `text_id` char(36) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Table structure for table `userLikes`
--

CREATE TABLE `userLikes` (
  `user_id` char(36) DEFAULT NULL,
  `text_id` char(36) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Table structure for table `tagLists`
--

CREATE TABLE `tagLists` (
  `tag_id` char(36) DEFAULT NULL,
  `name` varchar(300) DEFAULT NULL,
  `text_id` char(36) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
