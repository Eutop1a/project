USE project_question;

-- 创建主表：存储所有题目的公共属性
CREATE TABLE question_base (
                               `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
                               `type` VARCHAR(20) NOT NULL COMMENT '题型（choice/fill/judge/essay）',
                               `content` TEXT NOT NULL COMMENT '题目文本内容',
                               `image_path` VARCHAR(255) COMMENT '图片存储路径',
                               `difficulty` FLOAT CHECK (`difficulty` BETWEEN 1 AND 5) COMMENT '难度系数（1-5）',
                               `knowledge_id` BIGINT COMMENT '关联知识点分类',
                               `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP,        # 创建时间
                               `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                               `deleted_at` TIMESTAMP NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建知识点表
CREATE TABLE knowledge (
                           `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
                           `name` VARCHAR(255) NOT NULL COMMENT '知识点名称',
                           `description` TEXT COMMENT '知识点描述',
                           `parent_id` BIGINT COMMENT '父知识点 ID，用于构建知识点树结构',
                           `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
                           `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                           `deleted_at` TIMESTAMP NULL DEFAULT NULL COMMENT '删除时间，用于软删除',
    -- 添加外键约束，关联自身的父知识点
                           FOREIGN KEY (`parent_id`) REFERENCES knowledge(`id`) ON DELETE SET NULL ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建中间表，用于存储题目和知识点的多对多关系
CREATE TABLE question_knowledge_relation (
                                             `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
                                             `question_id` BIGINT NOT NULL,
                                             `knowledge_id` BIGINT NOT NULL,
                                             FOREIGN KEY (`question_id`) REFERENCES question_base(`id`) ON DELETE CASCADE ON UPDATE CASCADE,
                                             FOREIGN KEY (`knowledge_id`) REFERENCES knowledge(`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 为常用查询字段添加索引
CREATE INDEX idx_type ON question_base(`type`);
CREATE INDEX idx_difficulty ON question_base(`difficulty`);
CREATE INDEX idx_knowledge ON question_base(`knowledge_id`);

-- 创建选择题子表
CREATE TABLE question_choice (
                                 `id` BIGINT PRIMARY KEY,
                                 `option_a` TEXT NOT NULL,
                                 `option_b` TEXT NOT NULL,
                                 `option_c` TEXT NOT NULL,
                                 `option_d` TEXT NOT NULL,
                                 `answer` CHAR(1) NOT NULL CHECK (`answer` IN ('A', 'B', 'C', 'D')),
                                 FOREIGN KEY (`id`) REFERENCES question_base(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建填空题子表
CREATE TABLE question_fill (
                               `id` BIGINT PRIMARY KEY,
                               `answer` TEXT NOT NULL COMMENT '支持正则表达式（如 "北京|上海"）',
                               FOREIGN KEY (`id`) REFERENCES question_base(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建判断题子表
CREATE TABLE question_judge (
                                `id` BIGINT PRIMARY KEY,
                                `answer` BOOLEAN NOT NULL COMMENT 'TRUE/FALSE',
                                FOREIGN KEY (`id`) REFERENCES question_base(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建简答题/大题子表
CREATE TABLE question_essay (
                                `id` BIGINT PRIMARY KEY,
                                `reference` TEXT NOT NULL COMMENT '参考答案',
                                FOREIGN KEY (`id`) REFERENCES question_base(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建试卷表
CREATE TABLE exam_paper (
                            `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
                            `title` VARCHAR(255) NOT NULL COMMENT '试卷标题',
                            `questions` JSON NOT NULL COMMENT '题目列表（格式：[{"type":"choice","id":1}, ...]）',
                            `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP,        # 创建时间
                            `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                            `deleted_at` TIMESTAMP NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 为 question_base 表的 knowledge_id 字段添加外键约束，关联 knowledge 表的 id 字段
ALTER TABLE question_base
    ADD CONSTRAINT fk_question_base_knowledge
        FOREIGN KEY (`knowledge_id`) REFERENCES knowledge(`id`)
            ON DELETE SET NULL
            ON UPDATE CASCADE;
