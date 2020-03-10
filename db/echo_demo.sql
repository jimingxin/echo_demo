/*
 Navicat Premium Data Transfer

 Source Server         : LocalHost
 Source Server Type    : MySQL
 Source Server Version : 50720
 Source Host           : localhost:3306
 Source Schema         : echo_demo

 Target Server Type    : MySQL
 Target Server Version : 50720
 File Encoding         : 65001

 Date: 10/03/2020 20:46:27
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for class
-- ----------------------------
DROP TABLE IF EXISTS `class`;
CREATE TABLE `class` (
  `id` int(12) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) DEFAULT NULL,
  `desc` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of class
-- ----------------------------
BEGIN;
INSERT INTO `class` VALUES (1, '唐僧', '西天取经人');
INSERT INTO `class` VALUES (2, '宋江江', '喜欢屁股翘起来跪拜的那位');
INSERT INTO `class` VALUES (3, '猪无能', '天蓬元帅，高老庄那个');
INSERT INTO `class` VALUES (4, '西孟庆', '水浒有名的人物');
INSERT INTO `class` VALUES (5, '嵇蛋花', '喜欢讲故事的那个小朋友');
INSERT INTO `class` VALUES (6, '周淑华', '就是喜欢骂人的那位');
INSERT INTO `class` VALUES (7, '贾宝玉', '喜欢林妹妹那个混世魔王');
INSERT INTO `class` VALUES (12, '如来佛祖', '手指被猴子撒尿的');
INSERT INTO `class` VALUES (13, '观音菩萨', '手里拿个瓶子的男人');
INSERT INTO `class` VALUES (14, '小龙人', '一直找妈妈的那个娃');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
