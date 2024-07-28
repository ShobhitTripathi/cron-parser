# Cron Expression Parser

This is a command-line application that parses a cron string and expands each field to show the times at which it will run.

## Prerequisites

- [Go](https://golang.org/dl/) 1.16 or higher


## Installation

1. **Download or Clone the Project**: If you have a zip file, unzip it. If you're cloning from a repository, use `git clone`.

    ```sh
    unzip cron-parser.zip
    # or
    git clone https://github.com/ShobhitTripathi/cron-parser.git
    cd cron-parser
    ```

2. **Initialize and Download Dependencies**:

    ```sh
    go mod tidy
    ```

## Building the Application

1. **Navigate to the `cmd/cronparser` directory and build the application**:

    ```sh
    go build ./cmd/cronparser
    ```

## Running the Application

1. **Execute the built binary with a cron expression as an argument**:

    ```sh
    ./cronparser "*/15 0 1,15 * 1-5 /usr/bin/find"
    ```

   Example output:

    ```
    minute        0 15 30 45
    hour          0
    day of month  1 15
    month         1 2 3 4 5 6 7 8 9 10 11 12
    day of week   1 2 3 4 5
    command       /usr/bin/find
    ```

## Running Tests

- To run the tests, navigate to the project root and use the Go test tool:

    ```sh
    go test ./...
    ```

## Example Output

Given the input:

```sh
./cronparser "*/15 0 1,15 * 1-5 /usr/bin/find"
