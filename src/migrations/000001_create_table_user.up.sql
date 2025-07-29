CREATE TABLE users
(
    id                CHAR(36) PRIMARY KEY,      -- UUID dạng chuỗi
    name              VARCHAR(255) NOT NULL,
    email             VARCHAR(255) NOT NULL UNIQUE,
    password          VARCHAR(255) NOT NULL,
    avatar            VARCHAR(255) DEFAULT NULL,
    phone             VARCHAR(20)  DEFAULT NULL,
    address           TEXT         DEFAULT NULL,
    email_verified_at INT          DEFAULT NULL, -- Unix timestamp hoặc NULL
    status            TINYINT      DEFAULT 1,
    created_at        INT          NOT NULL,     -- Unix timestamp (seconds since epoch)
    updated_at        INT          NOT NULL
);
