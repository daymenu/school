CREATE DATABASE school;

CREATE TABLE school_info
(
    `ID` BIGINT PRIMARY KEY AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL DEFAULT '' COMMENT '学校名称',
    `CreateAt` DATETIME ,
    
    INDEX idx_name(name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='学校信息表';