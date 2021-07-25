# Crea's Train Routes App

## How to run application

* Install Go v1.16 or higher
* Build the application

```bash
go build -o app
```

* Then you can run application and you can specify csv file by using flag --file. If you not specify filepath, application
will default to routes.csv

```bash
./app --file=routes.csv
```

## How to run tests

```bash
go test ./...
```

## Improvement

* For the graph we can use adjacency list instead of metrix to improve memory usage.
* We can use priority queue instead of array to lower cpu for finding lowest distance node.
