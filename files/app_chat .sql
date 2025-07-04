-- phpMyAdmin SQL Dump
-- version 5.2.1deb3
-- https://www.phpmyadmin.net/
--
-- Host: localhost:3306
-- Generation Time: Jul 04, 2025 at 01:15 PM
-- Server version: 8.0.42-0ubuntu0.24.04.1
-- PHP Version: 8.3.6

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `app_chat`
--

-- --------------------------------------------------------

--
-- Table structure for table `chats`
--

CREATE TABLE `chats` (
  `id` int NOT NULL,
  `text` varchar(255) DEFAULT NULL,
  `room_id` int NOT NULL,
  `user_id` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `friends`
--

CREATE TABLE `friends` (
  `id` int NOT NULL,
  `name` varchar(255) NOT NULL,
  `friend` int NOT NULL,
  `user_id` int NOT NULL,
  `room_id` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `rooms`
--

CREATE TABLE `rooms` (
  `id` int NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- --------------------------------------------------------

--
-- Table structure for table `schema_migrations`
--

CREATE TABLE `schema_migrations` (
  `version` bigint NOT NULL,
  `dirty` tinyint(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `schema_migrations`
--

INSERT INTO `schema_migrations` (`version`, `dirty`) VALUES
(20250302092517, 0);

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int NOT NULL,
  `unique_number` int DEFAULT NULL,
  `name` varchar(255) NOT NULL,
  `username` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `last_login` date DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `unique_number`, `name`, `username`, `password`, `last_login`) VALUES
(1, 77846809, 'John', 'John Cena', '$2a$10$N.LvmtyXMfHzYI4rAVC5e.LakWipqfc4yhin7VmCgrKzNIUwrwClS', NULL),
(2, 90800251, 'Mary', 'Mary Cena', '$2a$10$7DKmpoGG./09CIDUzcUFXepGI2b3mkGyD3eExm.uMhs1ZEeFXQlYa', NULL),
(3, 30974724, 'Jane', 'Jane Cena', '$2a$10$KYicQhK/qREmbKbdQHUfyeJYbDdArusryWagdphCu9M9zqOz8RHeq', NULL),
(4, 64071565, 'Billy', 'Billy Cena', '$2a$10$KOY9hE8EtjBrxoyU.mikaOhT/Ya1V7XHIRJlU22CjX9rBg5Vj83.G', NULL),
(5, 23621306, 'Jeremy', 'Jeremy Cena', '$2a$10$kuTfNA4p/FOfpN3hdK/9.OotLM8/VxQvactVd8cDAi56CGbeeIDPa', NULL),
(6, 97933905, 'Skibidi', 'Skibidi Cena', '$2a$10$wGk20Z7gq2GVKHZ.Q1tL3ufWPm3gCwNFqwgIGWNrl.9p1AzXsgRCG', NULL),
(7, 38192662, 'Minto', 'Minto Cena', '$2a$10$xGKDSUpsS3xv/A9ASis1m.AclN1u6xdwqdghrDuLzpQXn6FEicro.', NULL),
(8, 14499291, 'Sulo', 'Sulo Cena', '$2a$10$NzoChXqe013sa2z8qOM5NeHpYKVOveAkw0IbyaO3YL7Wk6ugU.Kr2', NULL),
(9, 77517951, 'Rumi', 'Rumi Cena', '$2a$10$2L72x9WpTir8ZNlFsuni1OfRw2Sn9aXmVpZBLb5sTmIYIMqyvFWRW', NULL),
(10, 46444345, 'Jack', 'Jack Cena', '$2a$10$QZlX2EEXFT5E2Dxgu8T/DeyDv1FGk5z3UEFhD5th2UmP.rpLR3KyK', NULL),
(11, 96148715, 'Grealish', 'Grealish Cena', '$2a$10$qR8WLcWqJLyM4nl3zSTnZe2dn2B1hoVPqYoW2aO3Q/J/Q/NC2T9gO', NULL),
(12, 72707093, 'Naruto', 'Naruto Cena', '$2a$10$aDgVUe0NoD9w156XOEzziOATZchgPstHK7uTNC5qVKnV5vZKY/xLa', NULL);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `chats`
--
ALTER TABLE `chats`
  ADD PRIMARY KEY (`id`),
  ADD KEY `room_id` (`room_id`),
  ADD KEY `user_id` (`user_id`);

--
-- Indexes for table `friends`
--
ALTER TABLE `friends`
  ADD PRIMARY KEY (`id`),
  ADD KEY `friend` (`friend`),
  ADD KEY `user_id` (`user_id`),
  ADD KEY `room_id` (`room_id`);

--
-- Indexes for table `rooms`
--
ALTER TABLE `rooms`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `schema_migrations`
--
ALTER TABLE `schema_migrations`
  ADD PRIMARY KEY (`version`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `unique_number` (`unique_number`),
  ADD UNIQUE KEY `username` (`username`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `chats`
--
ALTER TABLE `chats`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `friends`
--
ALTER TABLE `friends`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `rooms`
--
ALTER TABLE `rooms`
  MODIFY `id` int NOT NULL AUTO_INCREMENT;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `chats`
--
ALTER TABLE `chats`
  ADD CONSTRAINT `chats_ibfk_1` FOREIGN KEY (`room_id`) REFERENCES `rooms` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `chats_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE;

--
-- Constraints for table `friends`
--
ALTER TABLE `friends`
  ADD CONSTRAINT `friends_ibfk_1` FOREIGN KEY (`friend`) REFERENCES `users` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `friends_ibfk_2` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE,
  ADD CONSTRAINT `friends_ibfk_3` FOREIGN KEY (`room_id`) REFERENCES `rooms` (`id`) ON DELETE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
