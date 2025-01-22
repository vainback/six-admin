/*
 Navicat Premium Dump SQL

 Source Server         : six
 Source Server Type    : MySQL
 Source Server Version : 50744 (5.7.44-log)
 Source Host           : 103.91.209.176:3306
 Source Schema         : six

 Target Server Type    : MySQL
 Target Server Version : 50744 (5.7.44-log)
 File Encoding         : 65001

 Date: 22/01/2025 19:19:42
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for auth_relation
-- ----------------------------
DROP TABLE IF EXISTS `auth_relation`;
CREATE TABLE `auth_relation`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime(3) NULL DEFAULT NULL,
  `update_time` datetime(3) NULL DEFAULT NULL,
  `tenant_id` bigint(20) NULL DEFAULT NULL COMMENT '租户id',
  `role_id` bigint(20) NULL DEFAULT NULL,
  `rule_id` bigint(20) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_auth_relation_tenant_id`(`tenant_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 148 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of auth_relation
-- ----------------------------
INSERT INTO `auth_relation` VALUES (1, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 12);
INSERT INTO `auth_relation` VALUES (2, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 1203);
INSERT INTO `auth_relation` VALUES (3, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 120306);
INSERT INTO `auth_relation` VALUES (4, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 120307);
INSERT INTO `auth_relation` VALUES (5, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 120304);
INSERT INTO `auth_relation` VALUES (6, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 120303);
INSERT INTO `auth_relation` VALUES (7, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 120302);
INSERT INTO `auth_relation` VALUES (8, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 120301);
INSERT INTO `auth_relation` VALUES (9, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 120305);
INSERT INTO `auth_relation` VALUES (10, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 1202);
INSERT INTO `auth_relation` VALUES (11, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 120205);
INSERT INTO `auth_relation` VALUES (12, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 120201);
INSERT INTO `auth_relation` VALUES (13, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 120207);
INSERT INTO `auth_relation` VALUES (14, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 1201);
INSERT INTO `auth_relation` VALUES (15, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 120106);
INSERT INTO `auth_relation` VALUES (16, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 120107);
INSERT INTO `auth_relation` VALUES (17, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 120104);
INSERT INTO `auth_relation` VALUES (18, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 120103);
INSERT INTO `auth_relation` VALUES (19, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 120102);
INSERT INTO `auth_relation` VALUES (20, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 120101);
INSERT INTO `auth_relation` VALUES (21, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 120105);
INSERT INTO `auth_relation` VALUES (22, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 11);
INSERT INTO `auth_relation` VALUES (23, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 1101);
INSERT INTO `auth_relation` VALUES (24, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110106);
INSERT INTO `auth_relation` VALUES (25, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110107);
INSERT INTO `auth_relation` VALUES (26, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110104);
INSERT INTO `auth_relation` VALUES (27, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110103);
INSERT INTO `auth_relation` VALUES (28, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110102);
INSERT INTO `auth_relation` VALUES (29, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110101);
INSERT INTO `auth_relation` VALUES (30, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110105);
INSERT INTO `auth_relation` VALUES (31, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 1102);
INSERT INTO `auth_relation` VALUES (32, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110206);
INSERT INTO `auth_relation` VALUES (33, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110207);
INSERT INTO `auth_relation` VALUES (34, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110208);
INSERT INTO `auth_relation` VALUES (35, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110209);
INSERT INTO `auth_relation` VALUES (36, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110204);
INSERT INTO `auth_relation` VALUES (37, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110203);
INSERT INTO `auth_relation` VALUES (38, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110202);
INSERT INTO `auth_relation` VALUES (39, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110201);
INSERT INTO `auth_relation` VALUES (40, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110205);
INSERT INTO `auth_relation` VALUES (41, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 1103);
INSERT INTO `auth_relation` VALUES (42, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110306);
INSERT INTO `auth_relation` VALUES (43, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110307);
INSERT INTO `auth_relation` VALUES (44, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110304);
INSERT INTO `auth_relation` VALUES (45, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110303);
INSERT INTO `auth_relation` VALUES (46, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110302);
INSERT INTO `auth_relation` VALUES (47, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110301);
INSERT INTO `auth_relation` VALUES (48, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110305);
INSERT INTO `auth_relation` VALUES (49, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 1104);
INSERT INTO `auth_relation` VALUES (50, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110406);
INSERT INTO `auth_relation` VALUES (51, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110407);
INSERT INTO `auth_relation` VALUES (52, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110404);
INSERT INTO `auth_relation` VALUES (53, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110403);
INSERT INTO `auth_relation` VALUES (54, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110402);
INSERT INTO `auth_relation` VALUES (55, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110401);
INSERT INTO `auth_relation` VALUES (56, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110405);
INSERT INTO `auth_relation` VALUES (57, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 1105);
INSERT INTO `auth_relation` VALUES (58, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110506);
INSERT INTO `auth_relation` VALUES (59, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110507);
INSERT INTO `auth_relation` VALUES (60, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110504);
INSERT INTO `auth_relation` VALUES (61, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110503);
INSERT INTO `auth_relation` VALUES (62, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110502);
INSERT INTO `auth_relation` VALUES (63, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110501);
INSERT INTO `auth_relation` VALUES (64, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110505);
INSERT INTO `auth_relation` VALUES (65, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 1106);
INSERT INTO `auth_relation` VALUES (66, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110606);
INSERT INTO `auth_relation` VALUES (67, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110607);
INSERT INTO `auth_relation` VALUES (68, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110604);
INSERT INTO `auth_relation` VALUES (69, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110603);
INSERT INTO `auth_relation` VALUES (70, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110602);
INSERT INTO `auth_relation` VALUES (71, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110601);
INSERT INTO `auth_relation` VALUES (72, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110605);
INSERT INTO `auth_relation` VALUES (73, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 1107);
INSERT INTO `auth_relation` VALUES (74, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110704);
INSERT INTO `auth_relation` VALUES (75, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110703);
INSERT INTO `auth_relation` VALUES (76, '2025-01-15 15:04:00.852', '2025-01-15 15:04:00.852', 1, 1, 110701);
INSERT INTO `auth_relation` VALUES (77, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 120301);
INSERT INTO `auth_relation` VALUES (78, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 120302);
INSERT INTO `auth_relation` VALUES (79, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 120303);
INSERT INTO `auth_relation` VALUES (80, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 120304);
INSERT INTO `auth_relation` VALUES (81, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 120305);
INSERT INTO `auth_relation` VALUES (82, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 120201);
INSERT INTO `auth_relation` VALUES (83, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 1201);
INSERT INTO `auth_relation` VALUES (84, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 110104);
INSERT INTO `auth_relation` VALUES (85, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 110103);
INSERT INTO `auth_relation` VALUES (86, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 110102);
INSERT INTO `auth_relation` VALUES (87, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 110101);
INSERT INTO `auth_relation` VALUES (88, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 110204);
INSERT INTO `auth_relation` VALUES (89, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 110203);
INSERT INTO `auth_relation` VALUES (90, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 110202);
INSERT INTO `auth_relation` VALUES (91, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 110201);
INSERT INTO `auth_relation` VALUES (92, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 110303);
INSERT INTO `auth_relation` VALUES (93, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 110302);
INSERT INTO `auth_relation` VALUES (94, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 110301);
INSERT INTO `auth_relation` VALUES (95, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 110304);
INSERT INTO `auth_relation` VALUES (96, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 110404);
INSERT INTO `auth_relation` VALUES (97, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 110403);
INSERT INTO `auth_relation` VALUES (98, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 110402);
INSERT INTO `auth_relation` VALUES (99, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 110401);
INSERT INTO `auth_relation` VALUES (100, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 110405);
INSERT INTO `auth_relation` VALUES (101, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 110504);
INSERT INTO `auth_relation` VALUES (102, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 110503);
INSERT INTO `auth_relation` VALUES (103, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 110502);
INSERT INTO `auth_relation` VALUES (104, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 110501);
INSERT INTO `auth_relation` VALUES (105, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 110505);
INSERT INTO `auth_relation` VALUES (106, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 110506);
INSERT INTO `auth_relation` VALUES (107, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 110507);
INSERT INTO `auth_relation` VALUES (108, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 1105);
INSERT INTO `auth_relation` VALUES (109, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 1106);
INSERT INTO `auth_relation` VALUES (110, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 120106);
INSERT INTO `auth_relation` VALUES (111, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 120107);
INSERT INTO `auth_relation` VALUES (112, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 120104);
INSERT INTO `auth_relation` VALUES (113, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 120103);
INSERT INTO `auth_relation` VALUES (114, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 120102);
INSERT INTO `auth_relation` VALUES (115, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 120101);
INSERT INTO `auth_relation` VALUES (116, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 120105);
INSERT INTO `auth_relation` VALUES (117, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 110606);
INSERT INTO `auth_relation` VALUES (118, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 110607);
INSERT INTO `auth_relation` VALUES (119, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 110604);
INSERT INTO `auth_relation` VALUES (120, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 110603);
INSERT INTO `auth_relation` VALUES (121, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 110602);
INSERT INTO `auth_relation` VALUES (122, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 110601);
INSERT INTO `auth_relation` VALUES (123, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 110605);
INSERT INTO `auth_relation` VALUES (124, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 1107);
INSERT INTO `auth_relation` VALUES (125, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 110704);
INSERT INTO `auth_relation` VALUES (126, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 110703);
INSERT INTO `auth_relation` VALUES (127, '2025-01-15 15:06:40.088', '2025-01-15 15:06:40.088', 1, 2, 110701);
INSERT INTO `auth_relation` VALUES (128, '2025-01-15 15:32:11.146', '2025-01-15 15:32:11.146', 1, 2, 11);
INSERT INTO `auth_relation` VALUES (129, '2025-01-15 15:32:11.146', '2025-01-15 15:32:11.146', 1, 2, 1101);
INSERT INTO `auth_relation` VALUES (130, '2025-01-15 15:32:11.146', '2025-01-15 15:32:11.146', 1, 2, 1102);
INSERT INTO `auth_relation` VALUES (131, '2025-01-15 15:32:11.146', '2025-01-15 15:32:11.146', 1, 2, 1103);
INSERT INTO `auth_relation` VALUES (132, '2025-01-15 15:32:11.146', '2025-01-15 15:32:11.146', 1, 2, 1104);
INSERT INTO `auth_relation` VALUES (133, '2025-01-15 15:32:11.146', '2025-01-15 15:32:11.146', 1, 2, 12);
INSERT INTO `auth_relation` VALUES (134, '2025-01-15 15:32:11.146', '2025-01-15 15:32:11.146', 1, 2, 1202);
INSERT INTO `auth_relation` VALUES (135, '2025-01-15 15:32:11.146', '2025-01-15 15:32:11.146', 1, 2, 1203);
INSERT INTO `auth_relation` VALUES (136, '2025-01-18 19:10:14.395', '2025-01-18 19:10:14.395', 1, 2, 1299);
INSERT INTO `auth_relation` VALUES (137, '2025-01-18 19:10:14.395', '2025-01-18 19:10:14.395', 1, 2, 129901);
INSERT INTO `auth_relation` VALUES (138, '2025-01-18 19:10:14.395', '2025-01-18 19:10:14.395', 1, 2, 129905);
INSERT INTO `auth_relation` VALUES (139, '2025-01-18 19:10:14.395', '2025-01-18 19:10:14.395', 1, 2, 129906);
INSERT INTO `auth_relation` VALUES (141, '2025-01-18 19:10:14.395', '2025-01-18 19:10:14.395', 1, 2, 99);
INSERT INTO `auth_relation` VALUES (142, '2025-01-18 19:10:14.395', '2025-01-18 19:10:14.395', 1, 2, 9901);
INSERT INTO `auth_relation` VALUES (143, '2025-01-18 19:10:14.395', '2025-01-18 19:10:14.395', 1, 2, 990101);
INSERT INTO `auth_relation` VALUES (144, '2025-01-18 19:10:14.395', '2025-01-18 19:10:14.395', 1, 2, 990103);
INSERT INTO `auth_relation` VALUES (145, '2025-01-18 19:10:14.395', '2025-01-18 19:10:14.395', 1, 2, 990104);
INSERT INTO `auth_relation` VALUES (146, '2025-01-18 19:10:14.395', '2025-01-18 19:10:14.395', 1, 2, 990105);
INSERT INTO `auth_relation` VALUES (147, '2025-01-18 19:10:14.395', '2025-01-18 19:10:14.395', 1, 2, 990106);

-- ----------------------------
-- Table structure for auth_role
-- ----------------------------
DROP TABLE IF EXISTS `auth_role`;
CREATE TABLE `auth_role`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime(3) NULL DEFAULT NULL,
  `update_time` datetime(3) NULL DEFAULT NULL,
  `delete_time` datetime(3) NULL DEFAULT NULL,
  `tenant_id` bigint(20) NULL DEFAULT NULL COMMENT '租户id',
  `parent_id` bigint(20) NULL DEFAULT NULL COMMENT '父级id',
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '角色名',
  `sign` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '角色签名',
  `status` tinyint(1) NULL DEFAULT NULL COMMENT '状态 -1 禁用 1启用',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_auth_role_sign`(`sign`) USING BTREE,
  INDEX `idx_auth_role_delete_time`(`delete_time`) USING BTREE,
  INDEX `idx_auth_role_tenant_id`(`tenant_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of auth_role
-- ----------------------------
INSERT INTO `auth_role` VALUES (1, '2025-01-15 14:54:30.980', '2025-01-15 14:54:30.980', NULL, 1, 0, '超级管理员', 'super', 1);
INSERT INTO `auth_role` VALUES (2, '2025-01-15 15:05:05.330', '2025-01-15 15:05:05.330', NULL, 1, 0, '游客', 'guest', 1);

-- ----------------------------
-- Table structure for auth_rule
-- ----------------------------
DROP TABLE IF EXISTS `auth_rule`;
CREATE TABLE `auth_rule`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime(3) NULL DEFAULT NULL,
  `update_time` datetime(3) NULL DEFAULT NULL,
  `delete_time` datetime(3) NULL DEFAULT NULL,
  `parent_id` bigint(20) NOT NULL DEFAULT 0 COMMENT '父级ID',
  `type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '菜单类型',
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '组件路径',
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '路由name',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '路由path',
  `api_route` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'API路由路径',
  `auth` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '权限标识',
  `locale` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '菜单名称',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '菜单图标',
  `order` bigint(20) NULL DEFAULT NULL COMMENT '显示排序 数字越小越靠前',
  `status` tinyint(4) NULL DEFAULT NULL COMMENT '菜单状态 -1 禁用 1 正常',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_auth_rule_delete_time`(`delete_time`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 990108 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of auth_rule
-- ----------------------------
INSERT INTO `auth_rule` VALUES (11, '0000-00-00 00:00:00.000', '2025-01-14 19:48:42.573', NULL, 0, 'menu', '', 'System', '/system', '', 'system', '系统管理', 'icon-settings', 98, 1);
INSERT INTO `auth_rule` VALUES (12, '2025-01-10 18:25:06.320', '2025-01-10 19:17:28.895', NULL, 0, 'menu', '', 'Basic', '/basic', '', 'basic', '基础设置', 'icon-desktop', 99, 1);
INSERT INTO `auth_rule` VALUES (99, '2025-01-18 10:47:49.124', '2025-01-18 10:47:49.124', NULL, 0, 'menu', 'Layout', 'Testss', '/testss', '', 'testss', '测试代码生成', 'icon-translate', 99999, 1);
INSERT INTO `auth_rule` VALUES (1101, '0000-00-00 00:00:00.000', '2025-01-10 18:40:09.956', NULL, 11, 'menu', ' ', 'AuthRule', 'authrule', '', 'system:rule', '菜单管理', 'icon-menu', 1, 1);
INSERT INTO `auth_rule` VALUES (1102, '0000-00-00 00:00:00.000', '2025-01-10 18:41:04.765', NULL, 11, 'menu', '', 'AuthRole', 'authrole', '', 'system:role', '角色管理', 'icon-user-group', 2, 1);
INSERT INTO `auth_rule` VALUES (1103, '0000-00-00 00:00:00.000', '2025-01-10 18:41:13.974', NULL, 11, 'menu', '', 'AuthUser', 'authuser', '', 'system:user', '用户管理', 'icon-user', 3, 1);
INSERT INTO `auth_rule` VALUES (1104, '0000-00-00 00:00:00.000', '2025-01-10 18:42:14.924', NULL, 11, 'menu', '', 'Tenant', 'tenant', '', 'system:tenant', '租户管理', 'icon-stamp', 4, 1);
INSERT INTO `auth_rule` VALUES (1105, '0000-00-00 00:00:00.000', '2025-01-10 18:42:48.915', NULL, 11, 'menu', '', 'Organization', 'organization', '', 'system:organization', '组织架构', 'icon-branch', 5, 1);
INSERT INTO `auth_rule` VALUES (1106, '0000-00-00 00:00:00.000', '2025-01-10 18:44:59.893', NULL, 11, 'menu', '', 'Job', 'job', '', 'system:job', '职位管理', 'icon-idcard', 6, 1);
INSERT INTO `auth_rule` VALUES (1107, '0000-00-00 00:00:00.000', '2025-01-10 18:45:34.307', NULL, 11, 'menu', '', 'Logs', 'logs', '', 'system:logs', '日志记录', 'icon-storage', 7, 1);
INSERT INTO `auth_rule` VALUES (1201, '2025-01-10 21:27:09.708', '2025-01-10 21:27:09.708', NULL, 12, 'menu', '', 'Dict', 'dict', '', 'basic:dict', '字典管理', 'icon-book', 1, 1);
INSERT INTO `auth_rule` VALUES (1202, '2025-01-10 21:28:47.501', '2025-01-10 21:34:46.868', NULL, 12, 'menu', '', 'Files', 'files', '', 'basic:files', '文件管理', 'icon-drive-file', 2, 1);
INSERT INTO `auth_rule` VALUES (1203, '2025-01-10 21:31:33.211', '2025-01-14 16:18:05.789', NULL, 12, 'menu', '', 'Cron', 'cron', '', 'basic:cron', '定时任务', 'icon-clock-circle', 3, 1);
INSERT INTO `auth_rule` VALUES (1299, '2025-01-10 21:33:21.921', '2025-01-18 08:40:42.728', NULL, 12, 'menu', '', 'Code', 'code', '', 'basic:code', '代码生成', 'icon-code', 99, 1);
INSERT INTO `auth_rule` VALUES (9901, '2025-01-18 17:37:00.455', '2025-01-18 17:37:00.455', NULL, 99, 'menu', '', 'Tests', 'tests', '', 'testss:tests', '测试模块', '', 44, 1);
INSERT INTO `auth_rule` VALUES (110101, '0000-00-00 00:00:00.000', '2025-01-10 18:34:28.379', NULL, 1101, 'page', '', '', '', '/auth/rule/list', 'system:rule:list', '列表（分页）', '', 1, 1);
INSERT INTO `auth_rule` VALUES (110102, '2025-01-10 17:09:53.779', '2025-01-10 18:34:34.103', NULL, 1101, 'page', '', '', '', '/auth/rule/tree-select', 'system:rule:tree-select', '列表（树）', '', 2, 1);
INSERT INTO `auth_rule` VALUES (110103, '2025-01-10 17:10:35.810', '2025-01-10 18:34:40.013', NULL, 1101, 'page', '', '', '', '/auth/rule/select', 'system:rule:select', '列表（选择）', '', 3, 1);
INSERT INTO `auth_rule` VALUES (110104, '2025-01-10 17:33:07.763', '2025-01-10 18:34:46.269', NULL, 1101, 'button', '', '', '', '/auth/rule/get', 'system:rule:get', '详情', '', 4, 1);
INSERT INTO `auth_rule` VALUES (110105, '0000-00-00 00:00:00.000', '2025-01-10 18:34:50.972', NULL, 1101, 'button', '', '', '', '/auth/rule/add', 'system:rule:add', '添加', '', 5, 1);
INSERT INTO `auth_rule` VALUES (110106, '2025-01-10 18:25:06.320', '2025-01-10 18:33:23.289', NULL, 1101, 'button', '', '', '', '/auth/rule/save', 'system:rule:save', '编辑', '', 6, 1);
INSERT INTO `auth_rule` VALUES (110107, '2025-01-10 18:25:06.320', '2025-01-10 18:25:06.320', NULL, 1101, 'button', '', '', '', '/auth/rule/del', 'system:rule:del', '删除', '', 7, 1);
INSERT INTO `auth_rule` VALUES (110201, '0000-00-00 00:00:00.000', '2025-01-10 18:34:28.379', NULL, 1102, 'page', '', '', '', '/auth/role/list', 'system:role:list', '列表（分页）', '', 1, 1);
INSERT INTO `auth_rule` VALUES (110202, '2025-01-10 17:09:53.779', '2025-01-10 18:34:34.103', NULL, 1102, 'page', '', '', '', '/auth/role/tree-select', 'system:role:tree-select', '列表（树）', '', 2, 1);
INSERT INTO `auth_rule` VALUES (110203, '2025-01-10 17:10:35.810', '2025-01-10 18:34:40.013', NULL, 1102, 'page', '', '', '', '/auth/role/select', 'system:role:select', '列表（选择）', '', 3, 1);
INSERT INTO `auth_rule` VALUES (110204, '2025-01-10 17:33:07.763', '2025-01-10 18:34:46.269', NULL, 1102, 'button', '', '', '', '/auth/role/get', 'system:role:get', '详情', '', 4, 1);
INSERT INTO `auth_rule` VALUES (110205, '0000-00-00 00:00:00.000', '2025-01-10 18:34:50.972', NULL, 1102, 'button', '', '', '', '/auth/role/add', 'system:role:add', '添加', '', 5, 1);
INSERT INTO `auth_rule` VALUES (110206, '2025-01-10 18:25:06.320', '2025-01-10 19:02:40.217', NULL, 1102, 'button', '', '', '', '/auth/role/save', 'system:role:save', '编辑', '', 6, 1);
INSERT INTO `auth_rule` VALUES (110207, '2025-01-10 18:25:06.320', '2025-01-10 18:25:06.320', NULL, 1102, 'button', '', '', '', '/auth/role/del', 'system:role:del', '删除', '', 7, 1);
INSERT INTO `auth_rule` VALUES (110208, '2025-01-10 18:25:06.320', '2025-01-10 18:33:23.289', NULL, 1102, 'button', '', '', '', '/auth/relation/select/rule', 'system:role:select-rule', '设置权限-获取权限列表', '', 8, 1);
INSERT INTO `auth_rule` VALUES (110209, '2025-01-10 18:25:06.320', '2025-01-10 18:33:23.289', NULL, 1102, 'button', '', '', '', '/auth/relation/set', 'system:role:set-rule', '设置权限-保存', '', 9, 1);
INSERT INTO `auth_rule` VALUES (110301, '0000-00-00 00:00:00.000', '2025-01-10 18:34:28.379', NULL, 1103, 'page', '', '', '', '/auth/user/list', 'system:user:list', '列表（分页）', '', 1, 1);
INSERT INTO `auth_rule` VALUES (110302, '2025-01-10 17:09:53.779', '2025-01-10 18:34:34.103', NULL, 1103, 'page', '', '', '', '/auth/user/tree-select', 'system:user:tree-select', '列表（树）', '', 2, 1);
INSERT INTO `auth_rule` VALUES (110303, '2025-01-10 17:10:35.810', '2025-01-10 18:34:40.013', NULL, 1103, 'page', '', '', '', '/auth/user/select', 'system:user:select', '列表（选择）', '', 3, 1);
INSERT INTO `auth_rule` VALUES (110304, '2025-01-10 17:33:07.763', '2025-01-10 18:34:46.269', NULL, 1103, 'button', '', '', '', '/auth/user/get', 'system:user:get', '详情', '', 4, 1);
INSERT INTO `auth_rule` VALUES (110305, '0000-00-00 00:00:00.000', '2025-01-10 18:34:50.972', NULL, 1103, 'button', '', '', '', '/auth/user/add', 'system:user:add', '添加', '', 5, 1);
INSERT INTO `auth_rule` VALUES (110306, '2025-01-10 18:25:06.320', '2025-01-10 18:33:23.289', NULL, 1103, 'button', '', '', '', '/auth/user/edit', 'system:user:edit', '编辑', '', 6, 1);
INSERT INTO `auth_rule` VALUES (110307, '2025-01-10 18:25:06.320', '2025-01-10 18:25:06.320', NULL, 1103, 'button', '', '', '', '/auth/user/del', 'system:user:del', '删除', '', 7, 1);
INSERT INTO `auth_rule` VALUES (110401, '0000-00-00 00:00:00.000', '2025-01-10 18:34:28.379', NULL, 1104, 'page', '', '', '', '/tenant/list', 'system:tenant:list', '列表（分页）', '', 1, 1);
INSERT INTO `auth_rule` VALUES (110402, '2025-01-10 17:09:53.779', '2025-01-10 18:34:34.103', NULL, 1104, 'page', '', '', '', '/tenant/tree-select', 'system:tenant:tree-select', '列表（树）', '', 2, 1);
INSERT INTO `auth_rule` VALUES (110403, '2025-01-10 17:10:35.810', '2025-01-10 18:34:40.013', NULL, 1104, 'page', '', '', '', '/tenant/select', 'system:tenant:select', '列表（选择）', '', 3, 1);
INSERT INTO `auth_rule` VALUES (110404, '2025-01-10 17:33:07.763', '2025-01-10 18:34:46.269', NULL, 1104, 'button', '', '', '', '/tenant/get', 'system:tenant:get', '详情', '', 4, 1);
INSERT INTO `auth_rule` VALUES (110405, '0000-00-00 00:00:00.000', '2025-01-10 18:34:50.972', NULL, 1104, 'button', '', '', '', '/tenant/add', 'system:tenant:add', '添加', '', 5, 1);
INSERT INTO `auth_rule` VALUES (110406, '2025-01-10 18:25:06.320', '2025-01-10 18:33:23.289', NULL, 1104, 'button', '', '', '', '/tenant/save', 'system:tenant:save', '编辑', '', 6, 1);
INSERT INTO `auth_rule` VALUES (110407, '2025-01-10 18:25:06.320', '2025-01-10 18:25:06.320', NULL, 1104, 'button', '', '', '', '/tenant/del', 'system:tenant:del', '删除', '', 7, 1);
INSERT INTO `auth_rule` VALUES (110501, '0000-00-00 00:00:00.000', '2025-01-10 18:34:28.379', NULL, 1105, 'page', '', '', '', '/organization/list', 'system:organization:list', '列表（分页）', '', 1, 1);
INSERT INTO `auth_rule` VALUES (110502, '2025-01-10 17:09:53.779', '2025-01-10 18:34:34.103', NULL, 1105, 'page', '', '', '', '/organization/tree-select', 'system:organization:tree-select', '列表（树）', '', 2, 1);
INSERT INTO `auth_rule` VALUES (110503, '2025-01-10 17:10:35.810', '2025-01-10 18:34:40.013', NULL, 1105, 'page', '', '', '', '/organization/select', 'system:organization:select', '列表（选择）', '', 3, 1);
INSERT INTO `auth_rule` VALUES (110504, '2025-01-10 17:33:07.763', '2025-01-10 18:34:46.269', NULL, 1105, 'button', '', '', '', '/organization/get', 'system:organization:get', '详情', '', 4, 1);
INSERT INTO `auth_rule` VALUES (110505, '0000-00-00 00:00:00.000', '2025-01-10 18:34:50.972', NULL, 1105, 'button', '', '', '', '/organization/add', 'system:organization:add', '添加', '', 5, 1);
INSERT INTO `auth_rule` VALUES (110506, '2025-01-10 18:25:06.320', '2025-01-10 18:33:23.289', NULL, 1105, 'button', '', '', '', '/organization/save', 'system:organization:save', '编辑', '', 6, 1);
INSERT INTO `auth_rule` VALUES (110507, '2025-01-10 18:25:06.320', '2025-01-10 18:25:06.320', NULL, 1105, 'button', '', '', '', '/organization/del', 'system:organization:del', '删除', '', 7, 1);
INSERT INTO `auth_rule` VALUES (110601, '0000-00-00 00:00:00.000', '2025-01-10 18:34:28.379', NULL, 1106, 'page', '', '', '', '/job/list', 'system:job:list', '列表（分页）', '', 1, 1);
INSERT INTO `auth_rule` VALUES (110602, '2025-01-10 17:09:53.779', '2025-01-10 18:34:34.103', NULL, 1106, 'page', '', '', '', '/job/tree-select', 'system:job:tree-select', '列表（树）', '', 2, 1);
INSERT INTO `auth_rule` VALUES (110603, '2025-01-10 17:10:35.810', '2025-01-10 18:34:40.013', NULL, 1106, 'page', '', '', '', '/job/select', 'system:job:select', '列表（选择）', '', 3, 1);
INSERT INTO `auth_rule` VALUES (110604, '2025-01-10 17:33:07.763', '2025-01-10 18:34:46.269', NULL, 1106, 'button', '', '', '', '/job/get', 'system:job:get', '详情', '', 4, 1);
INSERT INTO `auth_rule` VALUES (110605, '0000-00-00 00:00:00.000', '2025-01-10 18:34:50.972', NULL, 1106, 'button', '', '', '', '/job/add', 'system:job:add', '添加', '', 5, 1);
INSERT INTO `auth_rule` VALUES (110606, '2025-01-10 18:25:06.320', '2025-01-10 18:33:23.289', NULL, 1106, 'button', '', '', '', '/job/save', 'system:job:save', '编辑', '', 6, 1);
INSERT INTO `auth_rule` VALUES (110607, '2025-01-10 18:25:06.320', '2025-01-10 18:25:06.320', NULL, 1106, 'button', '', '', '', '/job/del', 'system:job:del', '删除', '', 7, 1);
INSERT INTO `auth_rule` VALUES (110701, '0000-00-00 00:00:00.000', '2025-01-10 18:34:28.379', NULL, 1107, 'page', '', '', '', '/logs/list', 'system:logs:list', '列表（分页）', '', 1, 1);
INSERT INTO `auth_rule` VALUES (110703, '2025-01-10 17:10:35.810', '2025-01-10 18:34:40.013', NULL, 1107, 'page', '', '', '', '/logs/select', 'system:logs:select', '列表（选择）', '', 3, 1);
INSERT INTO `auth_rule` VALUES (110704, '2025-01-10 17:33:07.763', '2025-01-10 18:57:04.777', NULL, 1107, 'button', '', '', '', '/logs/get', 'system:logs:get', '详情', '', 4, 1);
INSERT INTO `auth_rule` VALUES (120101, '0000-00-00 00:00:00.000', '2025-01-10 18:34:28.379', NULL, 1201, 'page', '', '', '', '/dict/list', 'basic:dict:list', '列表（分页）', '', 1, 1);
INSERT INTO `auth_rule` VALUES (120102, '2025-01-10 17:09:53.779', '2025-01-10 18:34:34.103', NULL, 1201, 'page', '', '', '', '/dict/tree-select', 'basic:dict:tree-select', '列表（树）', '', 2, 1);
INSERT INTO `auth_rule` VALUES (120103, '2025-01-10 17:10:35.810', '2025-01-10 18:34:40.013', NULL, 1201, 'page', '', '', '', '/dict/select', 'basic:dict:select', '列表（选择）', '', 3, 1);
INSERT INTO `auth_rule` VALUES (120104, '2025-01-10 17:33:07.763', '2025-01-10 18:34:46.269', NULL, 1201, 'button', '', '', '', '/dict/get', 'basic:dict:get', '详情', '', 4, 1);
INSERT INTO `auth_rule` VALUES (120105, '0000-00-00 00:00:00.000', '2025-01-10 18:34:50.972', NULL, 1201, 'button', '', '', '', '/dict/add', 'basic:dict:add', '添加', '', 5, 1);
INSERT INTO `auth_rule` VALUES (120106, '2025-01-10 18:25:06.320', '2025-01-10 18:33:23.289', NULL, 1201, 'button', '', '', '', '/dict/save', 'basic:dict:save', '编辑', '', 6, 1);
INSERT INTO `auth_rule` VALUES (120107, '2025-01-10 18:25:06.320', '2025-01-10 18:25:06.320', NULL, 1201, 'button', '', '', '', '/dict/del', 'basic:dict:del', '删除', '', 7, 1);
INSERT INTO `auth_rule` VALUES (120201, '2025-01-15 09:21:49.789', '2025-01-15 09:23:23.491', NULL, 1202, 'page', '', '', '', '/files/list', 'basic:files:list', '列表（分页）', '', 1, 1);
INSERT INTO `auth_rule` VALUES (120205, '2025-01-15 09:23:18.456', '2025-01-15 09:25:02.323', NULL, 1202, 'button', '', '', '', '/user/single/upload', 'basic:files:add', '上传', '', 5, 1);
INSERT INTO `auth_rule` VALUES (120207, '2025-01-15 09:21:49.789', '2025-01-15 09:21:49.789', NULL, 1202, 'button', '', '', '', '/files/del', 'basic:files:del', '删除', '', 7, 1);
INSERT INTO `auth_rule` VALUES (120301, '0000-00-00 00:00:00.000', '2025-01-10 18:34:28.379', NULL, 1203, 'page', '', '', '', '/cron/list', 'basic:cron:list', '列表（分页）', '', 1, 1);
INSERT INTO `auth_rule` VALUES (120302, '2025-01-10 17:09:53.779', '2025-01-10 18:34:34.103', NULL, 1203, 'page', '', '', '', '/cron/tree-select', 'basic:cron:tree-select', '列表（树）', '', 2, 1);
INSERT INTO `auth_rule` VALUES (120303, '2025-01-10 17:10:35.810', '2025-01-10 18:34:40.013', NULL, 1203, 'page', '', '', '', '/cron/select', 'basic:cron:select', '列表（选择）', '', 3, 1);
INSERT INTO `auth_rule` VALUES (120304, '2025-01-10 17:33:07.763', '2025-01-10 18:34:46.269', NULL, 1203, 'button', '', '', '', '/cron/get', 'basic:cron:get', '详情', '', 4, 1);
INSERT INTO `auth_rule` VALUES (120305, '0000-00-00 00:00:00.000', '2025-01-10 18:34:50.972', NULL, 1203, 'button', '', '', '', '/cron/add', 'basic:cron:add', '添加', '', 5, 1);
INSERT INTO `auth_rule` VALUES (120306, '2025-01-10 18:25:06.320', '2025-01-10 18:33:23.289', NULL, 1203, 'button', '', '', '', '/cron/save', 'basic:cron:save', '编辑', '', 6, 1);
INSERT INTO `auth_rule` VALUES (120307, '2025-01-10 18:25:06.320', '2025-01-10 18:25:06.320', NULL, 1203, 'button', '', '', '', '/cron/del', 'basic:cron:del', '删除', '', 7, 1);
INSERT INTO `auth_rule` VALUES (129901, '2025-01-18 18:53:01.656', '2025-01-18 18:53:01.656', NULL, 1299, 'page', '', '', '', '/codegen/list', 'basic:code:list', '列表（分页）', '', 1, 1);
INSERT INTO `auth_rule` VALUES (129905, '2025-01-18 18:53:51.458', '2025-01-18 18:54:16.657', NULL, 1299, 'button', '', '', '', '/codegen/add', 'basic:code:add', '添加', '', 5, 1);
INSERT INTO `auth_rule` VALUES (129906, '2025-01-18 18:54:11.301', '2025-01-18 18:54:11.301', NULL, 1299, 'button', '', '', '', '/codegen/save', 'basic:code:save', '编辑', '', 6, 1);
INSERT INTO `auth_rule` VALUES (129907, '2025-01-18 18:53:51.458', '2025-01-18 18:53:51.458', NULL, 1299, 'button', '', '', '', '/codegen/del', 'basic:code:del', '删除', '', 7, 1);
INSERT INTO `auth_rule` VALUES (129908, '2025-01-18 18:53:51.458', '2025-01-18 18:53:51.458', NULL, 1299, 'button', '', '', '', '/codegen/generator', 'basic:code:gen', '生成', '', 8, 1);
INSERT INTO `auth_rule` VALUES (990101, '2025-01-18 17:37:00.455', '2025-01-18 17:37:00.455', NULL, 9901, 'page', '', '', '', '/tests/list', 'testss:tests:list', '列表（分页）', '', 1, 1);
INSERT INTO `auth_rule` VALUES (990103, '2025-01-18 17:37:00.455', '2025-01-18 17:37:00.455', NULL, 9901, 'page', '', '', '', '/tests/select', 'testss:tests:select', '列表（选择）', '', 3, 1);
INSERT INTO `auth_rule` VALUES (990104, '2025-01-18 17:37:00.455', '2025-01-18 17:37:00.455', NULL, 9901, 'button', '', '', '', '/tests/get', 'testss:tests:get', '详情', '', 4, 1);
INSERT INTO `auth_rule` VALUES (990105, '2025-01-18 17:37:00.455', '2025-01-18 17:37:00.455', NULL, 9901, 'button', '', '', '', '/tests/add', 'testss:tests:add', '添加', '', 5, 1);
INSERT INTO `auth_rule` VALUES (990106, '2025-01-18 17:37:00.455', '2025-01-18 17:37:00.455', NULL, 9901, 'button', '', '', '', '/tests/save', 'testss:tests:save', '编辑', '', 6, 1);
INSERT INTO `auth_rule` VALUES (990107, '2025-01-18 17:37:00.455', '2025-01-18 17:37:00.455', NULL, 9901, 'button', '', '', '', '/tests/del', 'testss:tests:del', '删除', '', 7, 1);

-- ----------------------------
-- Table structure for auth_user
-- ----------------------------
DROP TABLE IF EXISTS `auth_user`;
CREATE TABLE `auth_user`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime(3) NULL DEFAULT NULL,
  `update_time` datetime(3) NULL DEFAULT NULL,
  `delete_time` datetime(3) NULL DEFAULT NULL,
  `tenant_id` bigint(20) NULL DEFAULT NULL COMMENT '租户id',
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户名 可以以手机号作为用户名',
  `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '密码',
  `nickname` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '昵称',
  `status` tinyint(1) NULL DEFAULT NULL COMMENT '状态 -1 禁用 1启用',
  `role_ids` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '角色ids',
  `is_root` tinyint(1) NULL DEFAULT 0 COMMENT '是否是root用户 1 是 0 否',
  `organization_id` bigint(20) NULL DEFAULT NULL,
  `job_id` bigint(20) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_auth_user_username`(`username`) USING BTREE,
  INDEX `idx_auth_user_delete_time`(`delete_time`) USING BTREE,
  INDEX `idx_auth_user_tenant_id`(`tenant_id`) USING BTREE,
  INDEX `idx_auth_user_organization_id`(`organization_id`) USING BTREE,
  INDEX `idx_auth_user_job_id`(`job_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of auth_user
-- ----------------------------
INSERT INTO `auth_user` VALUES (1, '2025-01-06 17:45:14.310', '2025-01-15 14:53:40.498', NULL, 1, 'root', '$2a$12$Gg4SVPU1eGNf9IgphPlWUeawyQJP1JeNH9QXPCmoS.Xifo9ddib6O', '六子', 1, '[-20240929]', 1, 1, 2);
INSERT INTO `auth_user` VALUES (2, '2025-01-15 15:08:56.133', '2025-01-15 15:08:56.133', NULL, 1, 'admin', '$2a$12$dDf1TbihY/636aVX8iHbQO5aMXOLniIdbHhs5eDTLk.m8lD3rF/AC', '游客', 1, '[2]', 0, 1, 1);

-- ----------------------------
-- Table structure for auth_user_login_storage
-- ----------------------------
DROP TABLE IF EXISTS `auth_user_login_storage`;
CREATE TABLE `auth_user_login_storage`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime(3) NULL DEFAULT NULL,
  `update_time` datetime(3) NULL DEFAULT NULL,
  `username` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'username',
  `token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'token',
  `expire` bigint(20) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of auth_user_login_storage
-- ----------------------------

-- ----------------------------
-- Table structure for code_generator
-- ----------------------------
DROP TABLE IF EXISTS `code_generator`;
CREATE TABLE `code_generator`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime(3) NULL DEFAULT NULL,
  `update_time` datetime(3) NULL DEFAULT NULL,
  `delete_time` datetime(3) NULL DEFAULT NULL,
  `table` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `fields` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `parent_module` bigint(20) NULL DEFAULT NULL,
  `is_soft_delete` tinyint(1) NULL DEFAULT NULL,
  `is_tenant` tinyint(1) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_code_generator_delete_time`(`delete_time`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of code_generator
-- ----------------------------
INSERT INTO `code_generator` VALUES (1, '2025-01-18 11:24:09.565', '2025-01-18 11:24:09.565', NULL, 'tests', '[{\"title\":\"测试普通字符串\",\"name\":\"te_str\",\"type\":\"varchar(255)\",\"default\":\"\",\"index\":\"index\",\"comment\":\"测试普通字符串\",\"is_keyword\":true,\"is_required\":true,\"dict_type\":\"\"},{\"title\":\"测试普通整形\",\"name\":\"te_int\",\"type\":\"int(11)\",\"default\":\"\",\"index\":\"index\",\"comment\":\"测试普通整形\",\"is_keyword\":true,\"is_required\":true,\"dict_type\":\"\"},{\"title\":\"测试大整形\",\"name\":\"te_bigint\",\"type\":\"bigint(20)\",\"default\":\"\",\"index\":\"index\",\"comment\":\"测试大整形\",\"is_keyword\":true,\"is_required\":true,\"dict_type\":\"\"},{\"title\":\"测试浮点数\",\"name\":\"te_float\",\"type\":\"float\",\"default\":\"0\",\"index\":\"\",\"comment\":\"测试浮点数\",\"is_keyword\":false,\"is_required\":true,\"dict_type\":\"\"},{\"title\":\"测试高精度浮点\",\"name\":\"te_decimal\",\"type\":\"decimal(10,2)\",\"default\":\"0\",\"index\":\"\",\"comment\":\"测试高精度浮点\",\"is_keyword\":false,\"is_required\":true,\"dict_type\":\"\"},{\"title\":\"测试单选下拉框\",\"name\":\"te_select\",\"type\":\"下拉框\",\"default\":\"\",\"index\":\"\",\"comment\":\"测试单选下拉框\",\"is_keyword\":false,\"is_required\":true,\"dict_type\":\"test\"},{\"title\":\"测试多选下拉框\",\"name\":\"te_select_many\",\"type\":\"下拉框多选\",\"default\":\"\",\"index\":\"\",\"comment\":\"测试多选下拉框\",\"is_keyword\":false,\"is_required\":true,\"dict_type\":\"test2\"},{\"title\":\"单选框\",\"name\":\"te_radio\",\"type\":\"单选框\",\"default\":\"\",\"index\":\"\",\"comment\":\"单选框\",\"is_keyword\":false,\"is_required\":true,\"dict_type\":\"test\"},{\"title\":\"复选框\",\"name\":\"te_checkbox\",\"type\":\"复选框\",\"default\":\"\",\"index\":\"\",\"comment\":\"复选框\",\"is_keyword\":false,\"is_required\":true,\"dict_type\":\"test2\"},{\"title\":\"开关\",\"name\":\"te_switch\",\"type\":\"开关\",\"default\":\"0\",\"index\":\"\",\"comment\":\"开关\",\"is_keyword\":false,\"is_required\":true,\"dict_type\":\"\"},{\"title\":\"时间选择器\",\"name\":\"te_timepicker\",\"type\":\"时间选择器\",\"default\":\"\",\"index\":\"\",\"comment\":\"时间选择器\",\"is_keyword\":false,\"is_required\":true,\"dict_type\":\"\"},{\"title\":\"日期选择器\",\"name\":\"te_datepicker\",\"type\":\"日期选择器\",\"default\":\"\",\"index\":\"\",\"comment\":\"日期选择器\",\"is_keyword\":false,\"is_required\":true,\"dict_type\":\"\"},{\"title\":\"日期时间选择器\",\"name\":\"te_datetimepicker\",\"type\":\"日期时间选择器\",\"default\":\"\",\"index\":\"\",\"comment\":\"日期时间选择器\",\"is_keyword\":false,\"is_required\":true,\"dict_type\":\"\"},{\"title\":\"上传图片\",\"name\":\"te_image_one\",\"type\":\"上传图片\",\"default\":\"\",\"index\":\"\",\"comment\":\"上传图片\",\"is_keyword\":false,\"is_required\":true,\"dict_type\":\"\"},{\"title\":\"上传图片（多图）\",\"name\":\"te_image_many\",\"type\":\"上传图片（多图）\",\"default\":\"\",\"index\":\"\",\"comment\":\"上传图片（多图）\",\"is_keyword\":false,\"is_required\":true,\"dict_type\":\"\"},{\"title\":\"上传视频\",\"name\":\"te_video\",\"type\":\"上传视频\",\"default\":\"\",\"index\":\"\",\"comment\":\"上传视频\",\"is_keyword\":false,\"is_required\":false,\"dict_type\":\"\"},{\"title\":\"上传文件\",\"name\":\"te_file\",\"type\":\"上传文件\",\"default\":\"\",\"index\":\"\",\"comment\":\"上传文件\",\"is_keyword\":false,\"is_required\":true,\"dict_type\":\"\"},{\"title\":\"文本域\",\"name\":\"te_text\",\"type\":\"text\",\"default\":\"\",\"index\":\"\",\"comment\":\"文本域\",\"is_keyword\":false,\"is_required\":true,\"dict_type\":\"\"},{\"title\":\"富文本\",\"name\":\"te_editor\",\"type\":\"富文本\",\"default\":\"\",\"index\":\"\",\"comment\":\"富文本\",\"is_keyword\":false,\"is_required\":true,\"dict_type\":\"\"}]', '测试模块', 99, 1, 1);
INSERT INTO `code_generator` VALUES (2, '2025-01-21 09:17:34.587', '2025-01-21 09:17:34.587', NULL, 'tests', '[{\"title\":\"测试普通字符串\",\"name\":\"te_str\",\"type\":\"varchar(255)\",\"default\":\"\",\"index\":\"index\",\"comment\":\"测试普通字符串\",\"is_keyword\":true,\"is_required\":true,\"dict_type\":\"\"},{\"title\":\"测试普通整形\",\"name\":\"te_int\",\"type\":\"int(11)\",\"default\":\"\",\"index\":\"index\",\"comment\":\"测试普通整形\",\"is_keyword\":true,\"is_required\":true,\"dict_type\":\"\"},{\"title\":\"测试大整形\",\"name\":\"te_bigint\",\"type\":\"bigint(20)\",\"default\":\"\",\"index\":\"index\",\"comment\":\"测试大整形\",\"is_keyword\":true,\"is_required\":true,\"dict_type\":\"\"},{\"title\":\"测试浮点数\",\"name\":\"te_float\",\"type\":\"float\",\"default\":\"0\",\"index\":\"\",\"comment\":\"测试浮点数\",\"is_keyword\":false,\"is_required\":true,\"dict_type\":\"\"},{\"title\":\"测试高精度浮点\",\"name\":\"te_decimal\",\"type\":\"decimal(10,2)\",\"default\":\"0\",\"index\":\"\",\"comment\":\"测试高精度浮点\",\"is_keyword\":false,\"is_required\":true,\"dict_type\":\"\"},{\"title\":\"测试单选下拉框\",\"name\":\"te_select\",\"type\":\"下拉框\",\"default\":\"\",\"index\":\"\",\"comment\":\"测试单选下拉框\",\"is_keyword\":false,\"is_required\":true,\"dict_type\":\"test\"},{\"title\":\"测试多选下拉框\",\"name\":\"te_select_many\",\"type\":\"下拉框多选\",\"default\":\"\",\"index\":\"\",\"comment\":\"测试多选下拉框\",\"is_keyword\":false,\"is_required\":true,\"dict_type\":\"test2\"},{\"title\":\"单选框\",\"name\":\"te_radio\",\"type\":\"单选框\",\"default\":\"\",\"index\":\"\",\"comment\":\"单选框\",\"is_keyword\":false,\"is_required\":true,\"dict_type\":\"\"},{\"title\":\"复选框\",\"name\":\"te_checkbox\",\"type\":\"复选框\",\"default\":\"\",\"index\":\"\",\"comment\":\"复选框\",\"is_keyword\":false,\"is_required\":true,\"dict_type\":\"\"},{\"title\":\"开关\",\"name\":\"te_switch\",\"type\":\"开关\",\"default\":\"0\",\"index\":\"\",\"comment\":\"开关\",\"is_keyword\":false,\"is_required\":true,\"dict_type\":\"\"},{\"title\":\"时间选择器\",\"name\":\"te_timepicker\",\"type\":\"时间选择器\",\"default\":\"\",\"index\":\"\",\"comment\":\"时间选择器\",\"is_keyword\":false,\"is_required\":true,\"dict_type\":\"\"},{\"title\":\"日期选择器\",\"name\":\"te_datepicker\",\"type\":\"日期选择器\",\"default\":\"\",\"index\":\"\",\"comment\":\"日期选择器\",\"is_keyword\":false,\"is_required\":true,\"dict_type\":\"\"},{\"title\":\"日期时间选择器\",\"name\":\"te_datetimepicker\",\"type\":\"日期时间选择器\",\"default\":\"\",\"index\":\"\",\"comment\":\"日期时间选择器\",\"is_keyword\":false,\"is_required\":true,\"dict_type\":\"\"},{\"title\":\"上传图片\",\"name\":\"te_image_one\",\"type\":\"上传图片\",\"default\":\"\",\"index\":\"\",\"comment\":\"上传图片\",\"is_keyword\":false,\"is_required\":true,\"dict_type\":\"\"},{\"title\":\"上传图片（多图）\",\"name\":\"te_image_many\",\"type\":\"上传图片（多图）\",\"default\":\"\",\"index\":\"\",\"comment\":\"上传图片（多图）\",\"is_keyword\":false,\"is_required\":true,\"dict_type\":\"\"},{\"title\":\"上传视频\",\"name\":\"te_video\",\"type\":\"上传视频\",\"default\":\"\",\"index\":\"\",\"comment\":\"上传视频\",\"is_keyword\":false,\"is_required\":true,\"dict_type\":\"\"},{\"title\":\"上传文件\",\"name\":\"te_file\",\"type\":\"上传文件\",\"default\":\"\",\"index\":\"\",\"comment\":\"上传文件\",\"is_keyword\":false,\"is_required\":true,\"dict_type\":\"\"},{\"title\":\"文本域\",\"name\":\"te_text\",\"type\":\"text\",\"default\":\"\",\"index\":\"\",\"comment\":\"文本域\",\"is_keyword\":false,\"is_required\":true,\"dict_type\":\"\"},{\"title\":\"富文本\",\"name\":\"te_editor\",\"type\":\"富文本\",\"default\":\"\",\"index\":\"\",\"comment\":\"富文本\",\"is_keyword\":false,\"is_required\":true,\"dict_type\":\"\"}]', '测试模块', 99, 1, 1);

-- ----------------------------
-- Table structure for cron_jobs
-- ----------------------------
DROP TABLE IF EXISTS `cron_jobs`;
CREATE TABLE `cron_jobs`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime(3) NULL DEFAULT NULL,
  `update_time` datetime(3) NULL DEFAULT NULL,
  `delete_time` datetime(3) NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `times` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `status` tinyint(1) NULL DEFAULT NULL COMMENT '状态 -1 禁用 1启用',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_cron_jobs_delete_time`(`delete_time`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of cron_jobs
-- ----------------------------
INSERT INTO `cron_jobs` VALUES (1, '2025-01-15 17:11:25.096', '2025-01-15 17:11:25.096', NULL, 'example', 'example', '*/15 * * * * *', 1);
INSERT INTO `cron_jobs` VALUES (2, '2025-01-17 13:24:23.917', '2025-01-17 13:24:23.917', NULL, 'example', 'example', '*/15 * * * * *', 1);
INSERT INTO `cron_jobs` VALUES (3, '2025-01-20 14:15:57.001', '2025-01-20 14:15:57.001', NULL, 'example1', 'example1', '10 * * * * *', 1);

-- ----------------------------
-- Table structure for dict
-- ----------------------------
DROP TABLE IF EXISTS `dict`;
CREATE TABLE `dict`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime(3) NULL DEFAULT NULL,
  `update_time` datetime(3) NULL DEFAULT NULL,
  `delete_time` datetime(3) NULL DEFAULT NULL,
  `tenant_id` bigint(20) NULL DEFAULT NULL COMMENT '租户id',
  `type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `label` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `value` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `color` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `is_sync` tinyint(1) NULL DEFAULT -1 COMMENT '状态 -1 不同步 1同步',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_dict_delete_time`(`delete_time`) USING BTREE,
  INDEX `idx_dict_tenant_id`(`tenant_id`) USING BTREE,
  INDEX `idx_dict_type`(`type`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of dict
-- ----------------------------
INSERT INTO `dict` VALUES (1, '2025-01-17 10:22:13.154', '2025-01-17 10:22:13.154', '2025-01-20 00:54:26.294', 1, 'root', '测试配置', 'test', '', -1);
INSERT INTO `dict` VALUES (2, '2025-01-17 10:22:23.437', '2025-01-18 10:43:51.299', NULL, 1, 'test', '测试1', '1', '#D8FF00', -1);
INSERT INTO `dict` VALUES (3, '2025-01-17 10:22:29.429', '2025-01-18 10:43:35.111', NULL, 1, 'test', '测试2', '2', '#0014FF', -1);
INSERT INTO `dict` VALUES (4, '2025-01-18 10:42:44.405', '2025-01-18 10:42:44.405', '2025-01-20 00:54:19.128', 1, 'root', '测试配置2', 'test2', '', -1);
INSERT INTO `dict` VALUES (5, '2025-01-18 10:43:10.743', '2025-01-18 10:43:10.743', NULL, 1, 'test2', 'xxx', 'xxx', '#40B551', -1);
INSERT INTO `dict` VALUES (6, '2025-01-18 10:43:21.313', '2025-01-18 10:43:21.313', NULL, 1, 'test2', 'sss', 'sss', '', -1);
INSERT INTO `dict` VALUES (7, '2025-01-22 14:50:08.684', '2025-01-22 14:50:08.684', NULL, 1, 'root', '大师的', '大师', '', 1);
INSERT INTO `dict` VALUES (8, '2025-01-22 15:35:07.469', '2025-01-22 15:35:07.469', NULL, 1, '大师', 'test', 'test', '#5A2E2E', 1);
INSERT INTO `dict` VALUES (9, '2025-01-22 15:35:25.476', '2025-01-22 15:35:25.476', NULL, 1, 'root', '测试', 'test', '', 1);
INSERT INTO `dict` VALUES (10, '2025-01-22 15:35:46.801', '2025-01-22 15:35:46.801', NULL, 1, 'root', 'ccc', 'ccc', '', -1);

-- ----------------------------
-- Table structure for files
-- ----------------------------
DROP TABLE IF EXISTS `files`;
CREATE TABLE `files`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime(3) NULL DEFAULT NULL,
  `update_time` datetime(3) NULL DEFAULT NULL,
  `delete_time` datetime(3) NULL DEFAULT NULL,
  `tenant_id` bigint(20) NULL DEFAULT NULL COMMENT '租户id',
  `mime` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `type` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `url` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_files_delete_time`(`delete_time`) USING BTREE,
  INDEX `idx_files_tenant_id`(`tenant_id`) USING BTREE,
  INDEX `idx_files_mime`(`mime`) USING BTREE,
  INDEX `idx_files_type`(`type`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 124 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of files
-- ----------------------------
INSERT INTO `files` VALUES (1, '2025-01-17 22:47:33.443', '2025-01-17 22:47:33.443', NULL, 1, 'image/jpeg', 'local', 'uploads/850235ece509ff435a644c230db367b6.jpg');
INSERT INTO `files` VALUES (2, '2025-01-17 22:55:47.520', '2025-01-17 22:55:47.520', NULL, 1, 'image/jpeg', 'local', 'uploads/23e7fc78ef2cc6886915c867f64d4203.jpg');
INSERT INTO `files` VALUES (3, '2025-01-17 22:55:47.533', '2025-01-17 22:55:47.533', NULL, 1, 'image/png', 'local', 'uploads/7a5ef5baa48be2256c6f6ac64789b56a.png');
INSERT INTO `files` VALUES (4, '2025-01-17 22:55:47.597', '2025-01-17 22:55:47.597', NULL, 1, 'image/png', 'local', 'uploads/b955b6eb175e4818948022e26f7cc71c.png');
INSERT INTO `files` VALUES (5, '2025-01-18 13:56:02.533', '2025-01-18 13:56:02.533', NULL, 1, 'image/png', 'local', 'uploads/0016c95f21ba94b3c50623c23cc8bec0.png');
INSERT INTO `files` VALUES (6, '2025-01-18 13:56:09.128', '2025-01-18 13:56:09.128', NULL, 1, 'image/jpeg', 'local', 'uploads/416b1e585d59b677aa9d693d94aacbfb.jpg');
INSERT INTO `files` VALUES (7, '2025-01-18 13:56:09.135', '2025-01-18 13:56:09.135', NULL, 1, 'image/jpeg', 'local', 'uploads/d91d1d809293180c31731071c987e455.jpg');
INSERT INTO `files` VALUES (8, '2025-01-18 13:56:12.123', '2025-01-18 13:56:12.123', NULL, 1, 'image/jpeg', 'local', 'uploads/416b1e585d59b677aa9d693d94aacbfb.jpg');
INSERT INTO `files` VALUES (9, '2025-01-18 14:01:18.192', '2025-01-18 14:01:18.192', NULL, 1, 'image/jpeg', 'local', 'uploads/d91d1d809293180c31731071c987e455.jpg');
INSERT INTO `files` VALUES (10, '2025-01-18 14:03:28.918', '2025-01-18 14:03:28.918', NULL, 1, 'image/jpeg', 'local', 'uploads/416b1e585d59b677aa9d693d94aacbfb.jpg');
INSERT INTO `files` VALUES (11, '2025-01-18 14:03:46.695', '2025-01-18 14:03:46.695', NULL, 1, 'image/jpeg', 'local', 'uploads/d91d1d809293180c31731071c987e455.jpg');
INSERT INTO `files` VALUES (12, '2025-01-18 14:04:08.781', '2025-01-18 14:04:08.781', NULL, 1, 'image/jpeg', 'local', 'uploads/416b1e585d59b677aa9d693d94aacbfb.jpg');
INSERT INTO `files` VALUES (13, '2025-01-18 14:05:03.413', '2025-01-18 14:05:03.413', NULL, 1, 'image/jpeg', 'local', 'uploads/416b1e585d59b677aa9d693d94aacbfb.jpg');
INSERT INTO `files` VALUES (14, '2025-01-18 14:05:26.064', '2025-01-18 14:05:26.064', NULL, 1, 'image/jpeg', 'local', 'uploads/416b1e585d59b677aa9d693d94aacbfb.jpg');
INSERT INTO `files` VALUES (15, '2025-01-18 14:07:38.245', '2025-01-18 14:07:38.245', NULL, 1, 'image/jpeg', 'local', 'uploads/d91d1d809293180c31731071c987e455.jpg');
INSERT INTO `files` VALUES (16, '2025-01-18 14:07:58.106', '2025-01-18 14:07:58.106', NULL, 1, 'image/jpeg', 'local', 'uploads/416b1e585d59b677aa9d693d94aacbfb.jpg');
INSERT INTO `files` VALUES (17, '2025-01-18 14:08:13.725', '2025-01-18 14:08:13.725', NULL, 1, 'image/jpeg', 'local', 'uploads/416b1e585d59b677aa9d693d94aacbfb.jpg');
INSERT INTO `files` VALUES (18, '2025-01-18 14:08:25.534', '2025-01-18 14:08:25.534', NULL, 1, 'image/jpeg', 'local', 'uploads/416b1e585d59b677aa9d693d94aacbfb.jpg');
INSERT INTO `files` VALUES (19, '2025-01-18 14:09:28.536', '2025-01-18 14:09:28.536', NULL, 1, 'image/jpeg', 'local', 'uploads/d91d1d809293180c31731071c987e455.jpg');
INSERT INTO `files` VALUES (20, '2025-01-18 14:09:33.028', '2025-01-18 14:09:33.028', NULL, 1, 'image/png', 'local', 'uploads/7a5ef5baa48be2256c6f6ac64789b56a.png');
INSERT INTO `files` VALUES (21, '2025-01-18 14:09:36.569', '2025-01-18 14:09:36.569', NULL, 1, 'image/png', 'local', 'uploads/7a5ef5baa48be2256c6f6ac64789b56a.png');
INSERT INTO `files` VALUES (22, '2025-01-18 14:09:36.559', '2025-01-18 14:09:36.559', NULL, 1, 'image/jpeg', 'local', 'uploads/23e7fc78ef2cc6886915c867f64d4203.jpg');
INSERT INTO `files` VALUES (23, '2025-01-18 14:09:36.579', '2025-01-18 14:09:36.579', NULL, 1, 'image/png', 'local', 'uploads/b955b6eb175e4818948022e26f7cc71c.png');
INSERT INTO `files` VALUES (24, '2025-01-18 14:09:42.244', '2025-01-18 14:09:42.244', NULL, 1, 'image/png', 'local', 'uploads/b955b6eb175e4818948022e26f7cc71c.png');
INSERT INTO `files` VALUES (25, '2025-01-18 14:11:54.838', '2025-01-18 14:11:54.838', NULL, 1, 'image/jpeg', 'local', 'uploads/416b1e585d59b677aa9d693d94aacbfb.jpg');
INSERT INTO `files` VALUES (26, '2025-01-18 14:11:57.140', '2025-01-18 14:11:57.140', NULL, 1, 'image/png', 'local', 'uploads/0016c95f21ba94b3c50623c23cc8bec0.png');
INSERT INTO `files` VALUES (27, '2025-01-18 14:12:10.127', '2025-01-18 14:12:10.127', NULL, 1, 'image/png', 'local', 'uploads/0016c95f21ba94b3c50623c23cc8bec0.png');
INSERT INTO `files` VALUES (28, '2025-01-18 14:12:12.173', '2025-01-18 14:12:12.173', NULL, 1, 'image/png', 'local', 'uploads/b955b6eb175e4818948022e26f7cc71c.png');
INSERT INTO `files` VALUES (29, '2025-01-18 14:12:14.567', '2025-01-18 14:12:14.567', NULL, 1, 'image/jpeg', 'local', 'uploads/64cbee4779a053d12be63513baee6501.jpg');
INSERT INTO `files` VALUES (30, '2025-01-18 14:13:01.006', '2025-01-18 14:13:01.006', NULL, 1, 'image/jpeg', 'local', 'uploads/416b1e585d59b677aa9d693d94aacbfb.jpg');
INSERT INTO `files` VALUES (31, '2025-01-18 14:13:03.807', '2025-01-18 14:13:03.807', NULL, 1, 'image/png', 'local', 'uploads/b955b6eb175e4818948022e26f7cc71c.png');
INSERT INTO `files` VALUES (32, '2025-01-18 14:13:19.782', '2025-01-18 14:13:19.782', NULL, 1, 'image/png', 'local', 'uploads/7a5ef5baa48be2256c6f6ac64789b56a.png');
INSERT INTO `files` VALUES (33, '2025-01-18 14:13:51.238', '2025-01-18 14:13:51.238', NULL, 1, 'image/png', 'local', 'uploads/7a5ef5baa48be2256c6f6ac64789b56a.png');
INSERT INTO `files` VALUES (34, '2025-01-18 14:13:53.816', '2025-01-18 14:13:53.816', NULL, 1, 'image/jpeg', 'local', 'uploads/23e7fc78ef2cc6886915c867f64d4203.jpg');
INSERT INTO `files` VALUES (35, '2025-01-18 14:14:17.377', '2025-01-18 14:14:17.377', NULL, 1, 'image/png', 'local', 'uploads/0016c95f21ba94b3c50623c23cc8bec0.png');
INSERT INTO `files` VALUES (36, '2025-01-18 14:14:24.083', '2025-01-18 14:14:24.083', NULL, 1, 'image/png', 'local', 'uploads/b955b6eb175e4818948022e26f7cc71c.png');
INSERT INTO `files` VALUES (37, '2025-01-18 14:15:15.807', '2025-01-18 14:15:15.807', NULL, 1, 'image/jpeg', 'local', 'uploads/416b1e585d59b677aa9d693d94aacbfb.jpg');
INSERT INTO `files` VALUES (38, '2025-01-18 14:15:18.516', '2025-01-18 14:15:18.516', NULL, 1, 'image/png', 'local', 'uploads/7a5ef5baa48be2256c6f6ac64789b56a.png');
INSERT INTO `files` VALUES (39, '2025-01-18 14:15:23.467', '2025-01-18 14:15:23.467', NULL, 1, 'image/jpeg', 'local', 'uploads/850235ece509ff435a644c230db367b6.jpg');
INSERT INTO `files` VALUES (40, '2025-01-18 14:15:42.412', '2025-01-18 14:15:42.412', NULL, 1, 'image/jpeg', 'local', 'uploads/a7aff36347460a71c342728d444e9f27.jpg');
INSERT INTO `files` VALUES (41, '2025-01-18 14:17:31.494', '2025-01-18 14:17:31.494', NULL, 1, 'image/jpeg', 'local', 'uploads/850235ece509ff435a644c230db367b6.jpg');
INSERT INTO `files` VALUES (42, '2025-01-18 14:17:37.781', '2025-01-18 14:17:37.781', NULL, 1, 'image/png', 'local', 'uploads/b955b6eb175e4818948022e26f7cc71c.png');
INSERT INTO `files` VALUES (43, '2025-01-18 14:17:57.643', '2025-01-18 14:17:57.643', NULL, 1, 'image/png', 'local', 'uploads/0016c95f21ba94b3c50623c23cc8bec0.png');
INSERT INTO `files` VALUES (44, '2025-01-18 14:18:00.822', '2025-01-18 14:18:00.822', NULL, 1, 'image/jpeg', 'local', 'uploads/23e7fc78ef2cc6886915c867f64d4203.jpg');
INSERT INTO `files` VALUES (45, '2025-01-18 14:20:18.054', '2025-01-18 14:20:18.054', NULL, 1, 'image/png', 'local', 'uploads/7a5ef5baa48be2256c6f6ac64789b56a.png');
INSERT INTO `files` VALUES (46, '2025-01-18 14:20:20.870', '2025-01-18 14:20:20.870', NULL, 1, 'image/png', 'local', 'uploads/7a5ef5baa48be2256c6f6ac64789b56a.png');
INSERT INTO `files` VALUES (47, '2025-01-18 14:20:53.584', '2025-01-18 14:20:53.584', NULL, 1, 'image/png', 'local', 'uploads/b955b6eb175e4818948022e26f7cc71c.png');
INSERT INTO `files` VALUES (48, '2025-01-18 14:22:36.901', '2025-01-18 14:22:36.901', NULL, 1, 'image/png', 'local', 'uploads/0016c95f21ba94b3c50623c23cc8bec0.png');
INSERT INTO `files` VALUES (49, '2025-01-18 14:22:42.305', '2025-01-18 14:22:42.305', NULL, 1, 'image/jpeg', 'local', 'uploads/850235ece509ff435a644c230db367b6.jpg');
INSERT INTO `files` VALUES (50, '2025-01-18 14:23:47.646', '2025-01-18 14:23:47.646', NULL, 1, 'image/jpeg', 'local', 'uploads/416b1e585d59b677aa9d693d94aacbfb.jpg');
INSERT INTO `files` VALUES (51, '2025-01-18 14:23:49.923', '2025-01-18 14:23:49.923', NULL, 1, 'image/jpeg', 'local', 'uploads/850235ece509ff435a644c230db367b6.jpg');
INSERT INTO `files` VALUES (52, '2025-01-18 14:28:08.314', '2025-01-18 14:28:08.314', NULL, 1, 'image/png', 'local', 'uploads/b955b6eb175e4818948022e26f7cc71c.png');
INSERT INTO `files` VALUES (53, '2025-01-18 14:28:10.318', '2025-01-18 14:28:10.318', NULL, 1, 'image/png', 'local', 'uploads/7a5ef5baa48be2256c6f6ac64789b56a.png');
INSERT INTO `files` VALUES (54, '2025-01-18 14:28:44.557', '2025-01-18 14:28:44.557', NULL, 1, 'image/png', 'local', 'uploads/7a5ef5baa48be2256c6f6ac64789b56a.png');
INSERT INTO `files` VALUES (55, '2025-01-18 14:28:44.552', '2025-01-18 14:28:44.552', NULL, 1, 'image/jpeg', 'local', 'uploads/23e7fc78ef2cc6886915c867f64d4203.jpg');
INSERT INTO `files` VALUES (56, '2025-01-18 14:28:44.565', '2025-01-18 14:28:44.565', NULL, 1, 'image/png', 'local', 'uploads/b955b6eb175e4818948022e26f7cc71c.png');
INSERT INTO `files` VALUES (57, '2025-01-18 14:29:05.388', '2025-01-18 14:29:05.388', NULL, 1, 'image/png', 'local', 'uploads/7a5ef5baa48be2256c6f6ac64789b56a.png');
INSERT INTO `files` VALUES (58, '2025-01-18 14:29:05.389', '2025-01-18 14:29:05.389', NULL, 1, 'image/png', 'local', 'uploads/0016c95f21ba94b3c50623c23cc8bec0.png');
INSERT INTO `files` VALUES (59, '2025-01-18 14:29:05.394', '2025-01-18 14:29:05.394', NULL, 1, 'image/jpeg', 'local', 'uploads/416b1e585d59b677aa9d693d94aacbfb.jpg');
INSERT INTO `files` VALUES (60, '2025-01-18 14:29:28.284', '2025-01-18 14:29:28.284', NULL, 1, 'image/jpeg', 'local', 'uploads/23e7fc78ef2cc6886915c867f64d4203.jpg');
INSERT INTO `files` VALUES (61, '2025-01-18 14:29:28.307', '2025-01-18 14:29:28.307', NULL, 1, 'image/png', 'local', 'uploads/b955b6eb175e4818948022e26f7cc71c.png');
INSERT INTO `files` VALUES (62, '2025-01-18 14:29:28.310', '2025-01-18 14:29:28.310', NULL, 1, 'image/png', 'local', 'uploads/7a5ef5baa48be2256c6f6ac64789b56a.png');
INSERT INTO `files` VALUES (63, '2025-01-18 14:32:58.047', '2025-01-18 14:32:58.047', NULL, 1, 'image/jpeg', 'local', 'uploads/416b1e585d59b677aa9d693d94aacbfb.jpg');
INSERT INTO `files` VALUES (64, '2025-01-18 14:33:00.374', '2025-01-18 14:33:00.374', NULL, 1, 'image/png', 'local', 'uploads/b955b6eb175e4818948022e26f7cc71c.png');
INSERT INTO `files` VALUES (65, '2025-01-18 14:34:07.738', '2025-01-18 14:34:07.738', NULL, 1, 'image/png', 'local', 'uploads/b955b6eb175e4818948022e26f7cc71c.png');
INSERT INTO `files` VALUES (66, '2025-01-18 14:34:07.744', '2025-01-18 14:34:07.744', NULL, 1, 'image/png', 'local', 'uploads/7a5ef5baa48be2256c6f6ac64789b56a.png');
INSERT INTO `files` VALUES (67, '2025-01-18 14:34:07.736', '2025-01-18 14:34:07.736', NULL, 1, 'image/jpeg', 'local', 'uploads/d91d1d809293180c31731071c987e455.jpg');
INSERT INTO `files` VALUES (68, '2025-01-18 14:34:07.747', '2025-01-18 14:34:07.747', NULL, 1, 'image/png', 'local', 'uploads/7a5ef5baa48be2256c6f6ac64789b56a.png');
INSERT INTO `files` VALUES (69, '2025-01-18 14:35:21.406', '2025-01-18 14:35:21.406', NULL, 1, 'image/png', 'local', 'uploads/b955b6eb175e4818948022e26f7cc71c.png');
INSERT INTO `files` VALUES (70, '2025-01-18 14:35:21.425', '2025-01-18 14:35:21.425', NULL, 1, 'image/png', 'local', 'uploads/7a5ef5baa48be2256c6f6ac64789b56a.png');
INSERT INTO `files` VALUES (71, '2025-01-18 14:35:21.447', '2025-01-18 14:35:21.447', NULL, 1, 'image/png', 'local', 'uploads/7a5ef5baa48be2256c6f6ac64789b56a.png');
INSERT INTO `files` VALUES (72, '2025-01-18 14:35:50.595', '2025-01-18 14:35:50.595', NULL, 1, 'image/jpeg', 'local', 'uploads/23e7fc78ef2cc6886915c867f64d4203.jpg');
INSERT INTO `files` VALUES (73, '2025-01-18 14:35:50.600', '2025-01-18 14:35:50.600', NULL, 1, 'image/jpeg', 'local', 'uploads/d91d1d809293180c31731071c987e455.jpg');
INSERT INTO `files` VALUES (74, '2025-01-18 14:35:50.612', '2025-01-18 14:35:50.612', NULL, 1, 'image/png', 'local', 'uploads/b955b6eb175e4818948022e26f7cc71c.png');
INSERT INTO `files` VALUES (75, '2025-01-18 14:36:23.661', '2025-01-18 14:36:23.661', NULL, 1, 'image/png', 'local', 'uploads/7a5ef5baa48be2256c6f6ac64789b56a.png');
INSERT INTO `files` VALUES (76, '2025-01-18 14:36:23.661', '2025-01-18 14:36:23.661', NULL, 1, 'image/png', 'local', 'uploads/b955b6eb175e4818948022e26f7cc71c.png');
INSERT INTO `files` VALUES (77, '2025-01-18 14:36:23.661', '2025-01-18 14:36:23.661', NULL, 1, 'image/png', 'local', 'uploads/7a5ef5baa48be2256c6f6ac64789b56a.png');
INSERT INTO `files` VALUES (78, '2025-01-18 14:37:35.198', '2025-01-18 14:37:35.198', NULL, 1, 'image/png', 'local', 'uploads/b955b6eb175e4818948022e26f7cc71c.png');
INSERT INTO `files` VALUES (79, '2025-01-18 14:37:35.192', '2025-01-18 14:37:35.192', NULL, 1, 'image/jpeg', 'local', 'uploads/d91d1d809293180c31731071c987e455.jpg');
INSERT INTO `files` VALUES (80, '2025-01-18 14:37:35.195', '2025-01-18 14:37:35.195', NULL, 1, 'image/jpeg', 'local', 'uploads/a7aff36347460a71c342728d444e9f27.jpg');
INSERT INTO `files` VALUES (81, '2025-01-18 14:38:30.263', '2025-01-18 14:38:30.263', NULL, 1, 'image/png', 'local', 'uploads/b955b6eb175e4818948022e26f7cc71c.png');
INSERT INTO `files` VALUES (82, '2025-01-18 14:38:30.255', '2025-01-18 14:38:30.255', NULL, 1, 'image/jpeg', 'local', 'uploads/23e7fc78ef2cc6886915c867f64d4203.jpg');
INSERT INTO `files` VALUES (83, '2025-01-18 14:38:30.261', '2025-01-18 14:38:30.261', NULL, 1, 'image/png', 'local', 'uploads/b955b6eb175e4818948022e26f7cc71c.png');
INSERT INTO `files` VALUES (84, '2025-01-18 14:40:53.863', '2025-01-18 14:40:53.863', NULL, 1, 'image/png', 'local', 'uploads/7a5ef5baa48be2256c6f6ac64789b56a.png');
INSERT INTO `files` VALUES (85, '2025-01-18 14:40:53.866', '2025-01-18 14:40:53.866', NULL, 1, 'image/png', 'local', 'uploads/b955b6eb175e4818948022e26f7cc71c.png');
INSERT INTO `files` VALUES (86, '2025-01-18 14:40:53.867', '2025-01-18 14:40:53.867', NULL, 1, 'image/png', 'local', 'uploads/b955b6eb175e4818948022e26f7cc71c.png');
INSERT INTO `files` VALUES (87, '2025-01-18 14:47:18.667', '2025-01-18 14:47:18.667', NULL, 1, 'image/png', 'local', 'uploads/b955b6eb175e4818948022e26f7cc71c.png');
INSERT INTO `files` VALUES (88, '2025-01-18 14:47:28.193', '2025-01-18 14:47:28.193', NULL, 1, 'image/jpeg', 'local', 'uploads/d91d1d809293180c31731071c987e455.jpg');
INSERT INTO `files` VALUES (89, '2025-01-18 14:47:28.197', '2025-01-18 14:47:28.197', NULL, 1, 'image/jpeg', 'local', 'uploads/23e7fc78ef2cc6886915c867f64d4203.jpg');
INSERT INTO `files` VALUES (90, '2025-01-18 14:47:28.212', '2025-01-18 14:47:28.212', NULL, 1, 'image/png', 'local', 'uploads/b955b6eb175e4818948022e26f7cc71c.png');
INSERT INTO `files` VALUES (91, '2025-01-18 15:32:59.302', '2025-01-18 15:32:59.302', NULL, 1, 'image/jpeg', 'local', 'uploads/d91d1d809293180c31731071c987e455.jpg');
INSERT INTO `files` VALUES (92, '2025-01-18 15:33:01.924', '2025-01-18 15:33:01.924', NULL, 1, 'image/png', 'local', 'uploads/b955b6eb175e4818948022e26f7cc71c.png');
INSERT INTO `files` VALUES (93, '2025-01-18 15:34:15.570', '2025-01-18 15:34:15.570', NULL, 1, 'image/jpeg', 'local', 'uploads/850235ece509ff435a644c230db367b6.jpg');
INSERT INTO `files` VALUES (94, '2025-01-18 15:37:00.557', '2025-01-18 15:37:00.557', NULL, 1, 'image/jpeg', 'local', 'uploads/850235ece509ff435a644c230db367b6.jpg');
INSERT INTO `files` VALUES (95, '2025-01-18 15:37:04.316', '2025-01-18 15:37:04.316', NULL, 1, 'image/jpeg', 'local', 'uploads/64cbee4779a053d12be63513baee6501.jpg');
INSERT INTO `files` VALUES (96, '2025-01-18 15:38:33.041', '2025-01-18 15:38:33.041', NULL, 1, 'image/png', 'local', 'uploads/7a5ef5baa48be2256c6f6ac64789b56a.png');
INSERT INTO `files` VALUES (97, '2025-01-18 15:38:35.180', '2025-01-18 15:38:35.180', NULL, 1, 'image/jpeg', 'local', 'uploads/23e7fc78ef2cc6886915c867f64d4203.jpg');
INSERT INTO `files` VALUES (98, '2025-01-18 15:38:51.248', '2025-01-18 15:38:51.248', NULL, 1, 'image/jpeg', 'local', 'uploads/850235ece509ff435a644c230db367b6.jpg');
INSERT INTO `files` VALUES (99, '2025-01-18 17:05:30.646', '2025-01-18 17:05:30.646', NULL, 1, 'image/jpeg', 'local', 'uploads/d91d1d809293180c31731071c987e455.jpg');
INSERT INTO `files` VALUES (100, '2025-01-18 17:05:33.831', '2025-01-18 17:05:33.831', NULL, 1, 'image/png', 'local', 'uploads/b955b6eb175e4818948022e26f7cc71c.png');
INSERT INTO `files` VALUES (101, '2025-01-18 17:05:33.831', '2025-01-18 17:05:33.831', NULL, 1, 'image/jpeg', 'local', 'uploads/23e7fc78ef2cc6886915c867f64d4203.jpg');
INSERT INTO `files` VALUES (102, '2025-01-18 17:05:36.693', '2025-01-18 17:05:36.693', NULL, 1, 'image/jpeg', 'local', 'uploads/64cbee4779a053d12be63513baee6501.jpg');
INSERT INTO `files` VALUES (103, '2025-01-18 17:07:51.877', '2025-01-18 17:07:51.877', NULL, 1, 'image/jpeg', 'local', 'uploads/d91d1d809293180c31731071c987e455.jpg');
INSERT INTO `files` VALUES (104, '2025-01-18 17:07:55.522', '2025-01-18 17:07:55.522', NULL, 1, 'image/jpeg', 'local', 'uploads/23e7fc78ef2cc6886915c867f64d4203.jpg');
INSERT INTO `files` VALUES (105, '2025-01-18 17:07:55.538', '2025-01-18 17:07:55.538', NULL, 1, 'image/png', 'local', 'uploads/b955b6eb175e4818948022e26f7cc71c.png');
INSERT INTO `files` VALUES (106, '2025-01-18 17:07:59.668', '2025-01-18 17:07:59.668', NULL, 1, 'image/jpeg', 'local', 'uploads/23e7fc78ef2cc6886915c867f64d4203.jpg');
INSERT INTO `files` VALUES (107, '2025-01-18 17:07:59.671', '2025-01-18 17:07:59.671', NULL, 1, 'image/jpeg', 'local', 'uploads/850235ece509ff435a644c230db367b6.jpg');
INSERT INTO `files` VALUES (108, '2025-01-18 17:10:47.719', '2025-01-18 17:10:47.719', NULL, 1, 'image/jpeg', 'local', 'uploads/850235ece509ff435a644c230db367b6.jpg');
INSERT INTO `files` VALUES (109, '2025-01-18 17:10:49.623', '2025-01-18 17:10:49.623', NULL, 1, 'image/png', 'local', 'uploads/b955b6eb175e4818948022e26f7cc71c.png');
INSERT INTO `files` VALUES (110, '2025-01-18 17:10:54.554', '2025-01-18 17:10:54.554', NULL, 1, 'image/jpeg', 'local', 'uploads/23e7fc78ef2cc6886915c867f64d4203.jpg');
INSERT INTO `files` VALUES (111, '2025-01-18 17:12:42.553', '2025-01-18 17:12:42.553', NULL, 1, 'image/png', 'local', 'uploads/b955b6eb175e4818948022e26f7cc71c.png');
INSERT INTO `files` VALUES (112, '2025-01-18 17:12:44.178', '2025-01-18 17:12:44.178', NULL, 1, 'image/png', 'local', 'uploads/7a5ef5baa48be2256c6f6ac64789b56a.png');
INSERT INTO `files` VALUES (113, '2025-01-18 17:12:49.202', '2025-01-18 17:12:49.202', NULL, 1, 'image/png', 'local', 'uploads/b955b6eb175e4818948022e26f7cc71c.png');
INSERT INTO `files` VALUES (114, '2025-01-18 17:14:34.189', '2025-01-18 17:14:34.189', NULL, 1, 'image/png', 'local', 'uploads/b955b6eb175e4818948022e26f7cc71c.png');
INSERT INTO `files` VALUES (115, '2025-01-18 17:14:37.800', '2025-01-18 17:14:37.800', NULL, 1, 'image/png', 'local', 'uploads/b955b6eb175e4818948022e26f7cc71c.png');
INSERT INTO `files` VALUES (116, '2025-01-18 17:14:37.804', '2025-01-18 17:14:37.804', NULL, 1, 'image/jpeg', 'local', 'uploads/23e7fc78ef2cc6886915c867f64d4203.jpg');
INSERT INTO `files` VALUES (117, '2025-01-18 17:14:40.166', '2025-01-18 17:14:40.166', NULL, 1, 'image/jpeg', 'local', 'uploads/64cbee4779a053d12be63513baee6501.jpg');
INSERT INTO `files` VALUES (118, '2025-01-18 17:37:32.504', '2025-01-18 17:37:32.504', NULL, 1, 'image/png', 'local', 'uploads/b955b6eb175e4818948022e26f7cc71c.png');
INSERT INTO `files` VALUES (119, '2025-01-18 17:37:36.854', '2025-01-18 17:37:36.854', NULL, 1, 'image/png', 'local', 'uploads/b955b6eb175e4818948022e26f7cc71c.png');
INSERT INTO `files` VALUES (120, '2025-01-18 17:37:36.855', '2025-01-18 17:37:36.855', NULL, 1, 'image/png', 'local', 'uploads/b955b6eb175e4818948022e26f7cc71c.png');
INSERT INTO `files` VALUES (121, '2025-01-18 17:37:36.854', '2025-01-18 17:37:36.854', NULL, 1, 'image/png', 'local', 'uploads/7a5ef5baa48be2256c6f6ac64789b56a.png');
INSERT INTO `files` VALUES (122, '2025-01-18 17:37:40.933', '2025-01-18 17:37:40.933', NULL, 1, 'image/jpeg', 'local', 'uploads/23e7fc78ef2cc6886915c867f64d4203.jpg');
INSERT INTO `files` VALUES (123, '2025-01-18 17:37:40.942', '2025-01-18 17:37:40.942', NULL, 1, 'image/png', 'local', 'uploads/b955b6eb175e4818948022e26f7cc71c.png');

-- ----------------------------
-- Table structure for jobs
-- ----------------------------
DROP TABLE IF EXISTS `jobs`;
CREATE TABLE `jobs`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime(3) NULL DEFAULT NULL,
  `update_time` datetime(3) NULL DEFAULT NULL,
  `delete_time` datetime(3) NULL DEFAULT NULL,
  `tenant_id` bigint(20) NULL DEFAULT NULL COMMENT '租户id',
  `parent_id` bigint(20) NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_jobs_delete_time`(`delete_time`) USING BTREE,
  INDEX `idx_jobs_tenant_id`(`tenant_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of jobs
-- ----------------------------
INSERT INTO `jobs` VALUES (1, '2025-01-15 15:08:09.620', '2025-01-15 15:08:09.620', '2025-01-20 01:03:13.842', 1, 0, '董事长', '');
INSERT INTO `jobs` VALUES (2, '2025-01-17 09:53:09.744', '2025-01-17 09:53:09.744', '2025-01-17 09:53:12.346', 1, 0, '销售', '');
INSERT INTO `jobs` VALUES (3, '2025-01-22 08:49:41.930', '2025-01-22 08:49:41.930', NULL, 1, 0, 'ceshi', '');

-- ----------------------------
-- Table structure for operate_logs
-- ----------------------------
DROP TABLE IF EXISTS `operate_logs`;
CREATE TABLE `operate_logs`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime(3) NULL DEFAULT NULL,
  `update_time` datetime(3) NULL DEFAULT NULL,
  `delete_time` datetime(3) NULL DEFAULT NULL,
  `tenant_id` bigint(20) NULL DEFAULT NULL COMMENT '租户id',
  `ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `uid` bigint(20) NULL DEFAULT NULL,
  `route` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `route_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `request_body` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `response_body` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `latency` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `agent` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_operate_logs_delete_time`(`delete_time`) USING BTREE,
  INDEX `idx_operate_logs_tenant_id`(`tenant_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4562 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of operate_logs
-- ----------------------------
INSERT INTO `operate_logs` VALUES (4559, '2025-01-22 19:19:31.673', '2025-01-22 19:19:31.673', NULL, 1, '114.252.153.14', 2, '/auth/rule/tree-select', '系统管理--菜单管理--列表（树）', '{\"page\":1,\"limit\":15,\"start_time\":\"\",\"end_time\":\"\",\"keyword\":\"\",\"order_by\":[{\"field\":\"`order`\",\"is_desc\":false}],\"is_delete\":false,\"model\":{\"order\":0,\"status\":0,\"id\":0,\"parent_id\":0,\"component\":\"\",\"name\":\"\",\"api_route\":\"\",\"icon\":\"\",\"type\":\"\",\"path\":\"\",\"auth\":\"\",\"locale\":\"\",\"tenant_id\":1}}', '', '0.006秒', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.95 Safari/537.36', 'POST');
INSERT INTO `operate_logs` VALUES (4560, '2025-01-22 19:19:33.142', '2025-01-22 19:19:33.142', NULL, 1, '114.252.153.14', 2, '/tenant/list', '系统管理--租户管理--列表（分页）', '{\"keyword\":\"\",\"order_by\":[{\"field\":\"id\",\"is_desc\":true}],\"is_delete\":false,\"model\":{\"sign\":\"\",\"domain\":\"\",\"status\":0,\"tenant_id\":1,\"id\":0,\"name\":\"\"},\"page\":1,\"limit\":15,\"start_time\":\"\",\"end_time\":\"\"}', '', '0.001秒', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.95 Safari/537.36', 'POST');
INSERT INTO `operate_logs` VALUES (4561, '2025-01-22 19:19:33.961', '2025-01-22 19:19:33.961', NULL, 1, '114.252.153.14', 2, '/organization/tree-select', '系统管理--组织架构--列表（树）', '{\"limit\":15,\"start_time\":\"\",\"end_time\":\"\",\"keyword\":\"\",\"order_by\":[{\"field\":\"id\",\"is_desc\":false}],\"is_delete\":false,\"model\":{\"id\":0,\"parent_id\":0,\"name\":\"\",\"description\":\"\",\"tenant_id\":1},\"page\":1}', '', '0.001秒', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.95 Safari/537.36', 'POST');

-- ----------------------------
-- Table structure for organization
-- ----------------------------
DROP TABLE IF EXISTS `organization`;
CREATE TABLE `organization`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime(3) NULL DEFAULT NULL,
  `update_time` datetime(3) NULL DEFAULT NULL,
  `delete_time` datetime(3) NULL DEFAULT NULL,
  `tenant_id` bigint(20) NULL DEFAULT NULL COMMENT '租户id',
  `parent_id` bigint(20) NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_organization_delete_time`(`delete_time`) USING BTREE,
  INDEX `idx_organization_tenant_id`(`tenant_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of organization
-- ----------------------------
INSERT INTO `organization` VALUES (1, '2025-01-15 15:08:03.185', '2025-01-15 15:08:03.185', NULL, 1, 0, '总裁办', '');
INSERT INTO `organization` VALUES (2, '2025-01-17 09:52:17.849', '2025-01-17 09:52:17.849', '2025-01-20 17:38:08.416', 1, 1, '助理', '');
INSERT INTO `organization` VALUES (3, '2025-01-17 11:54:36.852', '2025-01-17 11:54:36.852', NULL, 1, 0, '部长', '');
INSERT INTO `organization` VALUES (4, '2025-01-17 14:02:53.093', '2025-01-17 14:02:53.093', NULL, 1, 1, '44', '44');
INSERT INTO `organization` VALUES (5, '2025-01-20 17:38:20.952', '2025-01-20 17:38:20.952', NULL, 1, 4, 'ceshi', 'ces');
INSERT INTO `organization` VALUES (6, '2025-01-20 17:38:29.856', '2025-01-20 17:38:29.856', NULL, 1, 5, 'uu', 'u');
INSERT INTO `organization` VALUES (7, '2025-01-22 08:46:41.116', '2025-01-22 08:46:41.116', '2025-01-22 08:46:50.098', 1, 4, '55', '');
INSERT INTO `organization` VALUES (8, '2025-01-22 08:47:01.993', '2025-01-22 08:47:01.993', NULL, 1, 3, '66', '');
INSERT INTO `organization` VALUES (9, '2025-01-22 08:47:15.345', '2025-01-22 08:47:15.345', NULL, 1, 8, '77', '');

-- ----------------------------
-- Table structure for tenant
-- ----------------------------
DROP TABLE IF EXISTS `tenant`;
CREATE TABLE `tenant`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime(3) NULL DEFAULT NULL,
  `update_time` datetime(3) NULL DEFAULT NULL,
  `delete_time` datetime(3) NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户名',
  `sign` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '租户唯一标识',
  `domain` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '单独的域名',
  `status` tinyint(1) NULL DEFAULT NULL COMMENT '状态 -1 禁用 1启用',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uni_tenant_sign`(`sign`) USING BTREE,
  INDEX `idx_tenant_delete_time`(`delete_time`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tenant
-- ----------------------------
INSERT INTO `tenant` VALUES (1, NULL, NULL, NULL, 'Six', 'six', NULL, 1);
INSERT INTO `tenant` VALUES (2, '2025-01-17 17:18:27.119', '2025-01-17 17:18:27.119', NULL, 'test', 'aaa', '', 1);
INSERT INTO `tenant` VALUES (3, '2025-01-22 10:01:12.848', '2025-01-22 10:01:12.848', NULL, '9', '9', '', 1);

-- ----------------------------
-- Table structure for tests
-- ----------------------------
DROP TABLE IF EXISTS `tests`;
CREATE TABLE `tests`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime(3) NULL DEFAULT NULL,
  `update_time` datetime(3) NULL DEFAULT NULL,
  `delete_time` datetime(3) NULL DEFAULT NULL,
  `tenant_id` bigint(20) NULL DEFAULT NULL COMMENT '租户id',
  `te_str` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '测试普通字符串',
  `te_int` int(11) NOT NULL DEFAULT 0 COMMENT '测试普通整形',
  `te_bigint` bigint(20) NOT NULL DEFAULT 0 COMMENT '测试大整形',
  `te_float` float NOT NULL DEFAULT 0 COMMENT '测试浮点数',
  `te_decimal` decimal(10, 2) NOT NULL DEFAULT 0.00 COMMENT '测试高精度浮点',
  `te_select` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '测试单选下拉框',
  `te_select_many` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '测试多选下拉框',
  `te_radio` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '单选框',
  `te_checkbox` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '复选框',
  `te_switch` tinyint(1) NOT NULL DEFAULT 0 COMMENT '开关',
  `te_timepicker` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '时间选择器',
  `te_datepicker` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '日期选择器',
  `te_datetimepicker` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '日期时间选择器',
  `te_image_one` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '上传图片',
  `te_image_many` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '上传图片（多图）',
  `te_video` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '上传视频',
  `te_file` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '上传文件',
  `te_text` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '文本域',
  `te_editor` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '富文本',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_auth_user_delete_time`(`delete_time`) USING BTREE,
  INDEX `idx_auth_relation_tenant_id`(`tenant_id`) USING BTREE,
  INDEX `idx_tests_te_str`(`te_str`) USING BTREE,
  INDEX `idx_tests_te_int`(`te_int`) USING BTREE,
  INDEX `idx_tests_te_bigint`(`te_bigint`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of tests
-- ----------------------------
INSERT INTO `tests` VALUES (1, '2025-01-18 17:37:53.264', '2025-01-18 17:37:53.264', NULL, 1, 'asdasd', 99, 999999, 9999.99, 9999.99, '2', '[\"sss\"]', '2', '[\"sss\"]', 1, '17:37:24', '2025-01-18', '2025-01-18 17:37:28', '[\"uploads/b955b6eb175e4818948022e26f7cc71c.png\"]', '[\"uploads/b955b6eb175e4818948022e26f7cc71c.png\",\"uploads/7a5ef5baa48be2256c6f6ac64789b56a.png\",\"uploads/b955b6eb175e4818948022e26f7cc71c.png\"]', '[]', '[\"uploads/b955b6eb175e4818948022e26f7cc71c.png\",\"uploads/23e7fc78ef2cc6886915c867f64d4203.jpg\"]', 'asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少', '<p>asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少asdasdas收到工时费胜多负少</p>');

-- ----------------------------
-- Table structure for xxx
-- ----------------------------
DROP TABLE IF EXISTS `xxx`;
CREATE TABLE `xxx`  (
  `price` tinyint(4) NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of xxx
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
