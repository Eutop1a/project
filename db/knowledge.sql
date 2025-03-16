
-- 创建知识点表
CREATE TABLE knowledge (
                           `id` INT AUTO_INCREMENT PRIMARY KEY,
                           `name` VARCHAR(255) NOT NULL COMMENT '知识点名称',
                           `description` TEXT COMMENT '知识点描述',
                           `parent_id` INT COMMENT '父知识点 ID，用于构建知识点树结构',
                           `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                           `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                           `deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT '删除时间，用于软删除',
    -- 添加外键约束，关联自身的父知识点
                           FOREIGN KEY (`parent_id`) REFERENCES knowledge(`id`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
-- 为 question_base 表的 knowledge_id 字段添加外键约束，关联 knowledge 表的 id 字段
# ALTER TABLE question_base
#     ADD CONSTRAINT fk_question_base_knowledge
#         FOREIGN KEY (`knowledge_id`) REFERENCES knowledge(`id`)
#             ON DELETE SET NULL
#             ON UPDATE CASCADE;