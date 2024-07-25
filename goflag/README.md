# Goflag

A lightweight feature flagging application.

## Setup the database

The server uses a MySQL database to store flag data. To setup the database, run the following commands:

Create the database
```sh
mysql -u root -p
```
```sql
CREATE SCHEMA goflag_dev DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;
```
Run the schema script
```sh
mysql -u root -p goflag_dev < scripts/db/schema.sql
```

## Run the server

To run the server, execute:
```sh
go run cmd/api/main.go
```
The servers default port is `8080`.
