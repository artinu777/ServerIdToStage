-- +goose Up
-- SQL in this section is executed when the migration is applied.

CREATE TABLE stages(
                       id INT NOT NULL AUTO_INCREMENT,
                       server_id INT NULL,
                       stage INT NULL,
                       updated_at TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW(),
                       created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                       deleted_at TIMESTAMP null,
                       PRIMARY KEY(id)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.

DROP TABLE stages;
