-- phpMyAdmin SQL Dump
-- version 5.0.4
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Apr 30, 2021 at 01:03 PM
-- Server version: 10.4.17-MariaDB
-- PHP Version: 7.4.14

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+05:30";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `orders`
--

-- --------------------------------------------------------

--
-- Table structure for table `deliveries`
--

--
-- Dumping data for table `deliveries`
--

INSERT INTO `deliveries` (`delivery_id`, `location`, `status_id`, `order_id`, `otp`, `created_at`) VALUES
(1, 3, 1, 1, 8081, '2021-04-30 09:06:42'),
(2, 1, 3, 2, 7887, '2021-04-30 09:07:11'),
(3, 3, 1, 3, 1847, '2021-04-30 09:07:30'),
(4, 3, 1, 4, 4059, '2021-04-30 09:07:36');

-- --------------------------------------------------------

--
-- Table structure for table `items`
--

--
-- Dumping data for table `items`
--

INSERT INTO `items` (`item_id`, `item_code`, `description`, `quantity`, `order_id`) VALUES
(1, '123', 'oranges', 1, 1),
(2, '123', 'oranges', 1, 2),
(3, '123', 'oranges', 1, 3),
(4, '123', 'oranges', 1, 4);

-- --------------------------------------------------------

--
-- Table structure for table `orders`
--


--
-- Dumping data for table `orders`
--

INSERT INTO `orders` (`order_id`, `customer_name`, `mobile_number`, `ordered_at`) VALUES
(1, 'jerry', '7004995813', '0000-00-00 00:00:00'),
(2, 'jerry', '7004995813', '0000-00-00 00:00:00'),
(3, 'Harry', '7004995813', '0000-00-00 00:00:00'),
(4, 'Tom', '7004995813', '0000-00-00 00:00:00');

-- --------------------------------------------------------

--
-- Table structure for table `statuses`
--


--
-- Dumping data for table `statuses`
--

INSERT INTO `statuses` (`status_id`, `value`, `created_at`) VALUES
(1, 'Processing', '2021-04-27 00:00:00'),
(2, 'PickedUp', '2021-04-27 00:00:00'),
(3, 'Delivered', '2021-04-27 00:00:00');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `deliveries`
--

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `deliveries`
--
ALTER TABLE `deliveries`
  MODIFY `delivery_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `items`
--
ALTER TABLE `items`
  MODIFY `item_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `orders`
--
ALTER TABLE `orders`
  MODIFY `order_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;

--
-- AUTO_INCREMENT for table `statuses`
--
ALTER TABLE `statuses`
  MODIFY `status_id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
