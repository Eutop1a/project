USE project_user;
SHOW databases;

CREATE TABLE `paper_questions` (
                                   `paper_id` BIGINT UNSIGNED NOT NULL,
                                   `question_id` BIGINT UNSIGNED NOT NULL,
                                   `order` INT DEFAULT 0 COMMENT '题目在试卷中的顺序',
                                   PRIMARY KEY (`paper_id`, `question_id`),
                                   FOREIGN KEY (`paper_id`) REFERENCES `papers`(`id`) ON DELETE CASCADE,
                                   FOREIGN KEY (`question_id`) REFERENCES `questions`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
