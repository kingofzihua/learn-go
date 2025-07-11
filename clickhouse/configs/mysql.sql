/*
 Navicat Premium Dump SQL

 Source Server         : local-mysql
 Source Server Type    : MariaDB
 Source Server Version : 100434 (10.4.34-MariaDB-1:10.4.34+maria~ubu2004-log)
 Source Host           : 127.0.0.1:3306
 Source Schema         : learn_ck

 Target Server Type    : MariaDB
 Target Server Version : 100434 (10.4.34-MariaDB-1:10.4.34+maria~ubu2004-log)
 File Encoding         : 65001

 Date: 11/07/2025 19:58:03
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for orders
-- ----------------------------
DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `order_sn` bigint(20) NOT NULL,
  `user_id` bigint(20) DEFAULT NULL,
  `order_amount` decimal(10,2) NOT NULL COMMENT '订单金额',
  `is_refunded` tinyint(1) DEFAULT NULL COMMENT '是否退款, 0或1',
  `merchant_id` bigint(20) NOT NULL COMMENT '商户ID',
  `merchant_name` varchar(255) DEFAULT NULL COMMENT '商家名称',
  `source` varchar(255) NOT NULL COMMENT '订单来源',
  `payment_at` datetime DEFAULT NULL COMMENT '支付时间',
  `finish_at` datetime DEFAULT NULL COMMENT '完成时间',
  `order_status` tinyint(2) DEFAULT NULL COMMENT '订单状态',
  `created_at` datetime DEFAULT current_timestamp(),
  `updated_at` datetime DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of orders
-- ----------------------------
BEGIN;
INSERT INTO `orders` (`id`, `order_sn`, `user_id`, `order_amount`, `is_refunded`, `merchant_id`, `merchant_name`, `source`, `payment_at`, `finish_at`, `order_status`, `created_at`, `updated_at`) VALUES (1, 202310270001, 1001, 299.00, 0, 101, '潮流服饰店', 'index,good_list,detail:101,cart', '2023-10-20 10:05:00', '2023-10-25 15:30:00', 3, '2023-10-20 10:00:00', '2025-07-11 11:15:49');
INSERT INTO `orders` (`id`, `order_sn`, `user_id`, `order_amount`, `is_refunded`, `merchant_id`, `merchant_name`, `source`, `payment_at`, `finish_at`, `order_status`, `created_at`, `updated_at`) VALUES (2, 202310270002, 1002, 88.50, 0, 102, '数码先锋', 'banner:20,detail:121,cart', NULL, NULL, 0, '2023-10-27 11:30:00', '2025-07-11 11:16:32');
INSERT INTO `orders` (`id`, `order_sn`, `user_id`, `order_amount`, `is_refunded`, `merchant_id`, `merchant_name`, `source`, `payment_at`, `finish_at`, `order_status`, `created_at`, `updated_at`) VALUES (3, 202310270003, 1001, 1280.00, 1, 102, '数码先锋', 'index,detail:105,cart', '2023-10-15 09:12:00', '2023-10-18 11:00:00', 5, '2023-10-15 09:10:00', '2025-07-11 11:16:58');
INSERT INTO `orders` (`id`, `order_sn`, `user_id`, `order_amount`, `is_refunded`, `merchant_id`, `merchant_name`, `source`, `payment_at`, `finish_at`, `order_status`, `created_at`, `updated_at`) VALUES (4, 202310270004, 1003, 45.90, 0, 103, '居家生活馆', 'app,cart', '2023-10-26 23:15:00', NULL, 1, '2023-10-26 23:14:00', '2025-07-11 11:19:38');
INSERT INTO `orders` (`id`, `order_sn`, `user_id`, `order_amount`, `is_refunded`, `merchant_id`, `merchant_name`, `source`, `payment_at`, `finish_at`, `order_status`, `created_at`, `updated_at`) VALUES (5, 202310270005, 1004, 32.00, 0, 101, '潮流服饰店', 'search:kw=hat,cart', NULL, NULL, 4, '2023-10-25 14:00:00', '2025-07-11 11:17:06');
INSERT INTO `orders` (`id`, `order_sn`, `user_id`, `order_amount`, `is_refunded`, `merchant_id`, `merchant_name`, `source`, `payment_at`, `finish_at`, `order_status`, `created_at`, `updated_at`) VALUES (6, 202310270006, NULL, 599.00, 0, 104, '美妆个护', 'live:301,cart', '2023-09-30 20:30:00', '2023-10-03 16:45:00', 3, '2023-09-30 20:28:00', '2025-07-11 11:15:59');
INSERT INTO `orders` (`id`, `order_sn`, `user_id`, `order_amount`, `is_refunded`, `merchant_id`, `merchant_name`, `source`, `payment_at`, `finish_at`, `order_status`, `created_at`, `updated_at`) VALUES (7, 202310270007, 1005, 78.00, NULL, 103, '居家生活馆', 'app,cart', '2023-10-22 18:00:00', '2023-10-26 10:00:00', 3, '2023-10-22 17:55:00', '2025-07-11 11:19:41');
INSERT INTO `orders` (`id`, `order_sn`, `user_id`, `order_amount`, `is_refunded`, `merchant_id`, `merchant_name`, `source`, `payment_at`, `finish_at`, `order_status`, `created_at`, `updated_at`) VALUES (8, 202310270008, 1002, 2150.50, 0, 102, '数码先锋', 'banner:15,detail:110,cart', '2023-10-21 11:11:11', '2023-10-24 12:00:00', 3, '2023-10-21 11:10:00', '2025-07-11 11:15:03');
INSERT INTO `orders` (`id`, `order_sn`, `user_id`, `order_amount`, `is_refunded`, `merchant_id`, `merchant_name`, `source`, `payment_at`, `finish_at`, `order_status`, `created_at`, `updated_at`) VALUES (9, 202310270009, 1006, 19.90, 0, 103, '居家生活馆', 'search:kw=towel,detail:200,cart', '2023-10-27 08:30:00', NULL, 2, '2023-10-27 08:29:00', '2025-07-11 11:16:19');
INSERT INTO `orders` (`id`, `order_sn`, `user_id`, `order_amount`, `is_refunded`, `merchant_id`, `merchant_name`, `source`, `payment_at`, `finish_at`, `order_status`, `created_at`, `updated_at`) VALUES (10, 202310270010, 1001, 399.00, 0, 101, '潮流服饰店', 'detail:102,direct', '2023-10-19 14:20:00', '2023-10-22 19:00:00', 3, '2023-10-19 14:18:00', '2025-07-11 11:17:26');
INSERT INTO `orders` (`id`, `order_sn`, `user_id`, `order_amount`, `is_refunded`, `merchant_id`, `merchant_name`, `source`, `payment_at`, `finish_at`, `order_status`, `created_at`, `updated_at`) VALUES (11, 202310270011, 1007, 102.00, 0, 104, '美妆个护', 'app,doctor,cart', NULL, NULL, 0, '2023-10-27 12:00:00', '2025-07-11 11:19:51');
INSERT INTO `orders` (`id`, `order_sn`, `user_id`, `order_amount`, `is_refunded`, `merchant_id`, `merchant_name`, `source`, `payment_at`, `finish_at`, `order_status`, `created_at`, `updated_at`) VALUES (12, 202310270012, 1008, 8999.00, 0, 102, '数码先锋', 'user_share:125,detail:108,direct', '2023-08-18 10:00:00', '2023-08-21 17:00:00', 3, '2023-08-18 09:58:00', '2025-07-11 11:17:43');
INSERT INTO `orders` (`id`, `order_sn`, `user_id`, `order_amount`, `is_refunded`, `merchant_id`, `merchant_name`, `source`, `payment_at`, `finish_at`, `order_status`, `created_at`, `updated_at`) VALUES (13, 202310270013, NULL, 68.00, 0, 101, '潮流服饰店', 'live:88,direct', '2023-10-26 13:00:00', NULL, 1, '2023-10-26 12:59:00', '2025-07-11 11:18:00');
INSERT INTO `orders` (`id`, `order_sn`, `user_id`, `order_amount`, `is_refunded`, `merchant_id`, `merchant_name`, `source`, `payment_at`, `finish_at`, `order_status`, `created_at`, `updated_at`) VALUES (14, 202310270014, 1009, 25.50, 0, 103, '居家生活馆', 'index-feed,cart', '2023-10-25 16:40:00', '2023-10-27 09:10:00', 3, '2023-10-25 16:38:00', '2025-07-11 11:18:59');
INSERT INTO `orders` (`id`, `order_sn`, `user_id`, `order_amount`, `is_refunded`, `merchant_id`, `merchant_name`, `source`, `payment_at`, `finish_at`, `order_status`, `created_at`, `updated_at`) VALUES (15, 202310270015, NULL, 158.00, 0, 104, '美妆个护', 'search:kw=mask, cart', NULL, NULL, 4, '2023-10-24 21:05:00', '2025-07-11 11:19:08');
INSERT INTO `orders` (`id`, `order_sn`, `user_id`, `order_amount`, `is_refunded`, `merchant_id`, `merchant_name`, `source`, `payment_at`, `finish_at`, `order_status`, `created_at`, `updated_at`) VALUES (16, 202310270016, 1010, 55.00, 0, 101, '潮流服饰店', 'banner:30,direct', '2023-10-27 10:10:10', NULL, 1, '2023-10-27 10:10:00', '2025-07-11 11:20:15');
INSERT INTO `orders` (`id`, `order_sn`, `user_id`, `order_amount`, `is_refunded`, `merchant_id`, `merchant_name`, `source`, `payment_at`, `finish_at`, `order_status`, `created_at`, `updated_at`) VALUES (17, 202310270017, 1001, 128.00, 0, 103, '居家生活馆', 'index-feed,cart', '2022-11-11 00:15:00', '2022-11-15 14:00:00', 3, '2022-11-11 00:12:00', '2025-07-11 11:19:19');
INSERT INTO `orders` (`id`, `order_sn`, `user_id`, `order_amount`, `is_refunded`, `merchant_id`, `merchant_name`, `source`, `payment_at`, `finish_at`, `order_status`, `created_at`, `updated_at`) VALUES (18, 202310270018, 1002, 499.00, 1, 102, '数码先锋', 'detail:106,cart', '2023-10-20 15:00:00', '2023-10-23 18:00:00', 5, '2023-10-20 14:58:00', '2025-07-11 11:19:25');
INSERT INTO `orders` (`id`, `order_sn`, `user_id`, `order_amount`, `is_refunded`, `merchant_id`, `merchant_name`, `source`, `payment_at`, `finish_at`, `order_status`, `created_at`, `updated_at`) VALUES (19, 202310270019, 1011, 89.90, 0, 104, '美妆个护', 'search:kw=hat,adv:1254,detail:581,cart', '2023-10-26 09:00:00', NULL, 2, '2023-10-26 08:59:00', '2025-07-11 11:20:47');
INSERT INTO `orders` (`id`, `order_sn`, `user_id`, `order_amount`, `is_refunded`, `merchant_id`, `merchant_name`, `source`, `payment_at`, `finish_at`, `order_status`, `created_at`, `updated_at`) VALUES (20, 202310270020, 1012, 199.00, 0, 101, '潮流服饰店', 'live:99,cart', NULL, NULL, 0, '2023-10-27 14:00:00', '2025-07-11 11:15:03');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
