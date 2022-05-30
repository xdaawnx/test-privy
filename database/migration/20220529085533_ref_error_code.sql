-- +goose Up
-- +goose StatementBegin
CREATE TABLE `ref_error_codess` (
  `code` char(3) NOT NULL COMMENT 'Error codes',
  `message` varchar(255) DEFAULT NULL COMMENT 'Message of error',
  `messageID` varchar(255) DEFAULT NULL,
  `description` text COMMENT 'Description',
  `created_date` datetime DEFAULT NULL,
  `created_by` int NOT NULL DEFAULT '1',
  `updated_date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `updated_by` int NOT NULL DEFAULT '1',
  PRIMARY KEY (`code`)
) ENGINE=InnoDB;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP table `ref_error_codes`;
-- +goose StatementEnd
