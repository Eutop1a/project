USE project_question;
SHOW databases;

CREATE TABLE `questions` (
                             `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '题目ID',
                             `type` TINYINT NOT NULL COMMENT '题型 (1: 选择, 2: 填空, 3: 简答, 4: 大题)',
                             `content` TEXT NOT NULL COMMENT '题目内容 (JSON格式存储图文混合内容)',
                             `answer` TEXT NOT NULL COMMENT '参考答案',
                             `difficulty` FLOAT DEFAULT 0.5 COMMENT '难度系数 (0-1)',
                             `knowledge_point` VARCHAR(255) COMMENT '所属知识点',
                             `creator_id` BIGINT UNSIGNED NOT NULL COMMENT '创建者ID',
                             `created_at` DATETIME DEFAULT CURRENT_TIMESTAMP,
                             `updated_at` DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                             PRIMARY KEY (`id`),
                             KEY `idx_type` (`type`),
                             KEY `idx_difficulty` (`difficulty`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
