-- To create the database use the following command
-- CREATE SCHEMA baragoda_dev DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE barcode_group (
  id INT NOT NULL AUTO_INCREMENT PRIMARY KEY
, prefix VARCHAR(32) NOT NULL UNIQUE KEY
, sequence INT NOT NULL DEFAULT 0
, created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);