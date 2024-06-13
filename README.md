# graph-formatting
A function that takes a map of integer numbers and dates in Unix time as the first argument and the required interval for formatting as the second argument. Intervals are as follows: "MONTH", "WEEK", "DAY", "HOUR". Resulting slice of maps contains grouped pairs of integer numbers and dates in Unix time. Timestamps are changed to UTC as it was not specified which timezone I should use. Data for functions is represented by graph map in the main.go.

## Usage

In order to select desired interval - change interval variable in main.go on line 9 to any of this: "MONTH", "WEEK", "DAY", "HOUR"

## Run Test

Execute the following command to run unit test:

```shell
go test ./...
```

### Run Tests with coverage

```shell
go test -cover ./...
```

If you want to export coverage report, execute the following command:

```shell
go tool cover -html=coverage.out
```