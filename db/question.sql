USE project_question;

-- 创建主表：存储所有题目的公共属性
CREATE TABLE question_base (
                               `id` INT AUTO_INCREMENT PRIMARY KEY,
                               `type` VARCHAR(20) NOT NULL COMMENT '题型（choice/fill/judge/essay）',
                               `content` TEXT NOT NULL COMMENT '题目文本内容',
                               `image_path` VARCHAR(255) COMMENT '图片存储路径',
                               `difficulty` FLOAT CHECK (`difficulty` BETWEEN 1 AND 5) COMMENT '难度系数（1-5）',
                               `knowledge_id` INT COMMENT '关联知识点分类',
                               `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP,        # 创建时间
                               `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                               `deleted_at` TIMESTAMP NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 为常用查询字段添加索引
CREATE INDEX idx_type ON question_base(`type`);
CREATE INDEX idx_difficulty ON question_base(`difficulty`);
CREATE INDEX idx_knowledge ON question_base(`knowledge_id`);

-- 创建选择题子表
CREATE TABLE question_choice (
                                 `id` INT PRIMARY KEY,
                                 `option_a` TEXT NOT NULL,
                                 `option_b` TEXT NOT NULL,
                                 `option_c` TEXT NOT NULL,
                                 `option_d` TEXT NOT NULL,
                                 `answer` CHAR(1) NOT NULL CHECK (`answer` IN ('A', 'B', 'C', 'D')),
                                 FOREIGN KEY (`id`) REFERENCES question_base(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建填空题子表
CREATE TABLE question_fill (
                               `id` INT PRIMARY KEY,
                               `answer` TEXT NOT NULL COMMENT '支持正则表达式（如 "北京|上海"）',
                               FOREIGN KEY (`id`) REFERENCES question_base(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建判断题子表
CREATE TABLE question_judge (
                                `id` INT PRIMARY KEY,
                                `answer` BOOLEAN NOT NULL COMMENT 'TRUE/FALSE',
                                FOREIGN KEY (`id`) REFERENCES question_base(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建简答题/大题子表
CREATE TABLE question_essay (
                                `id` INT PRIMARY KEY,
                                `reference` TEXT NOT NULL COMMENT '参考答案',
                                FOREIGN KEY (`id`) REFERENCES question_base(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- 创建试卷表
CREATE TABLE exam_paper (
                            `id` INT AUTO_INCREMENT PRIMARY KEY,
                            `title` VARCHAR(255) NOT NULL COMMENT '试卷标题',
                            `questions` JSON NOT NULL COMMENT '题目列表（格式：[{"type":"choice","id":1}, ...]）',
                            `create_time` DATETIME DEFAULT CURRENT_TIMESTAMP,        # 创建时间
                            `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                            `deleted_at` TIMESTAMP NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;