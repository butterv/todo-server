CREATE TABLE users (
    id         VARCHAR(128) NOT NULL,
    email      VARCHAR(254) NOT NULL,
    created_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME,
    PRIMARY KEY (id),
    INDEX users_email (email)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE user_passwords (
    user_id       VARCHAR(128) NOT NULL,
    password_hash VARCHAR(128) NOT NULL,
    created_at    DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at    DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id),
    CONSTRAINT fk_user_passwords_user_id
        FOREIGN KEY (user_id) REFERENCES users (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE user_tokens (
    id         BIGINT       NOT NULL AUTO_INCREMENT,
    user_id    VARCHAR(128) NOT NULL,
    token      VARCHAR(255) NOT NULL,
    created_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME,
    PRIMARY KEY (id),
    CONSTRAINT fk_user_tokens_user_id
        FOREIGN KEY (user_id) REFERENCES users (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE labels (
    id         BIGINT      NOT NULL AUTO_INCREMENT,
    color      VARCHAR(32) NOT NULL,
    name       VARCHAR(64) NOT NULL,
    created_at DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME,
    PRIMARY KEY (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE projects (
    id         BIGINT      NOT NULL AUTO_INCREMENT,
    color      VARCHAR(32) NOT NULL,
    name       VARCHAR(64) NOT NULL,
    created_at DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME,
    PRIMARY KEY (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE todos (
    id          BIGINT       NOT NULL AUTO_INCREMENT,
    subject     VARCHAR(128) NOT NULL,
    label_id    BIGINT,
    project_id  BIGINT,
    description TEXT,
    created_at  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at  DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at  DATETIME,
    PRIMARY KEY (id),
    CONSTRAINT fk_todos_label_id
        FOREIGN KEY (label_id) REFERENCES labels (id),
    CONSTRAINT fk_todos_project_id
        FOREIGN KEY (project_id) REFERENCES projects (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE todo_comments (
    id         BIGINT   NOT NULL AUTO_INCREMENT,
    todo_id    BIGINT   NOT NULL,
    comment    TEXT     NOT NULL,
    created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME,
    PRIMARY KEY (id),
    CONSTRAINT fk_todo_comments_todo_id
        FOREIGN KEY (todo_id) REFERENCES todos (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;

CREATE TABLE user_todos (
    user_id    VARCHAR(128) NOT NULL,
    todo_id    BIGINT       NOT NULL,
    created_at DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at DATETIME,
    PRIMARY KEY (user_id, todo_id),
    CONSTRAINT fk_user_todos_user_id
        FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT fk_user_todos_todo_id
        FOREIGN KEY (todo_id) REFERENCES todos (id)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4;