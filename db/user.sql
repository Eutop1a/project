
CREATE TABLE user (
    id BIGINT PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL,  -- 使用bcrypt加密存储
    is_admin BOOLEAN NOT NULL DEFAULT FALSE,  -- 是否是管理员
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP NULL DEFAULT NULL

);

-- 可选：为deleted_at添加索引（提升查询性能）
CREATE INDEX idx_user_deleted_at ON user(deleted_at);