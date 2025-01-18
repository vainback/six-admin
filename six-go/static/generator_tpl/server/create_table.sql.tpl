CREATE TABLE `{{tpl - $table}}`  (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `create_time` datetime(3) NULL DEFAULT NULL,
  `update_time` datetime(3) NULL DEFAULT NULL,
  <tpl - $delete_time>
  <tpl - $tenant>
  <tpl - $fields>
  PRIMARY KEY (`id`) USING BTREE,
  <tpl - $fields-index>
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;