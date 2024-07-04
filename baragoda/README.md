# Baragoda

A partial rewrite in Go of The Wellcome Sanger Institute's [Baracoda](https://github.com/sanger/baracoda).

## Setup the database

The server uses a MySQL database to store barcode groups and barcodes. To setup the database, run the following commands:

Create the database
```sh
mysql -u root -p
```
```sql
CREATE SCHEMA baragoda_dev DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;
```
Run the schema script
```sh
mysql -u root -p baragoda_dev < scripts/db/schema.sql
```

## Run the server

To run the server, execute:
```sh
go run cmd/api/main.go
```
The servers default port is `8080`.

## API

<summary><code>GET</code>  <code><b>/barcode_groups</b></code> Returns a list of all barcode groups</summary>

<summary><code>GET</code>  <code><b>/barcode_groups/{barcode_group_prefix}</b></code> Returns a barcode group by prefix</summary>

<summary><code>POST</code> <code><b>/barcode_groups/{barcode_group_prefix}/new</b></code> Takes a prefix and sequence and creates a new barcode group </summary>

<summary><code>POST</code> <code><b>/barcodes/{barcode_group_prefix}/new</b></code> Returns a unique barcode from a barcode group given its prefix</summary>

<summary><code>GET</code>  <code><b>/barcodes/{barcode_group_prefix}/last</b></code> Returns the last barcode from a barcode group</summary>


### Examples

Creating a new barcode group
```sh
curl --header "Content-Type: application/json" --request POST --data '{"sequence": "1", "prefix": "new" }' --location http://localhost:8080/barcode_groups/new
```

Creating a new barcode from a barcode group with the prefix sqpd
```sh
curl --location --request POST 'http://localhost:8080/barcodes/sqpd/new'
```

Getting the last barcode from a barcode group with the prefix sqpd
```sh
curl --location 'http://localhost:8080/barcodes/sqpd/last'
```

### TODO

- [ ] General testing
- [ ] Dockerize the application
- [ ] internal/handlers/create_barcode_group.go should return an appropriate error message when the barcode group already exists
