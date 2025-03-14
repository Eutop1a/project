
CREATE TABLE user (
                       id BIGINT PRIMARY KEY,
                       username VARCHAR(50) UNIQUE NOT NULL,
                       password_hash VARCHAR(255) NOT NULL,  -- 使用bcrypt加密存储
                       is_admin BOOLEAN NOT NULL DEFAULT FALSE  -- 是否是管理员
);