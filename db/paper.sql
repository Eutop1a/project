USE project_paper;
SHOW databases;

CREATE TABLE `papers` (
                          `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '试卷ID',
                          `name` VARCHAR(255) NOT NULL COMMENT '试卷名称',
                          `creator_id` BIGINT UNSIGNED NOT NULL COMMENT '创建者ID',
                          `similarity_threshold` FLOAT DEFAULT 0.3 COMMENT '相似度阈值 (默认30%)',
                          `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
                          PRIMARY KEY (`id`),
                          KEY `idx_creator` (`creator_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

