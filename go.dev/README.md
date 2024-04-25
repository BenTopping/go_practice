## go.dev
This directory contains a series of notes and code snippets from the [go.dev](https://go.dev/) tutorials.
This README documents what I worked on and when.

### 04/04/24
- Directories created/used:
    - `hello_world`: hello world program
    - `hello`: module that imports greetings and prints a message
    - `greetings`: module for greeting messages with tests
- Covered:
    - Basics: syntax, types, and variables
    - Modules: creation, imports, and versioning
    - Tests: writing and running
Up to: https://go.dev/doc/tutorial/add-a-test

### 11/04/24
- Directions created/used:
    - `workspaces`: workspace with multiple modules
    - `data-access`: module for database access and actions
- Covered:
    - Compilation and execution
    - Workspaces
    - Database access
        - Connecting to a DB
        - Multi-row and single-row queries
        - Adding data
Up to: https://go.dev/doc/tutorial/web-service-gin

### 18/04/24
- Directions created/used:
    - `web-service-gin`: web service with Gin
    - `generics`: generics in Go
    - `fuzz`: fuzzing in Go
- Covered:
    - Creating a basic REST API using Gin
        - Basic GET and POST requests
    - Go generics
        - Non-generic functions
        - Generic functions taking multiple types
        - Type constraints
    - Fuzzing
        - Writing and running fuzz tests
Up to: https://go.dev/doc/tutorial/govulncheck

### 25/04/24
- Directories created/used
    - `vuln-tutorial`: vulnerability checking in Go
- Covered:
    - Checking for vulnerabilities in Go code
        - Using `govulncheck`
