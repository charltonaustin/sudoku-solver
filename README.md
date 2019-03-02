# sudoku-solver

## build
`mkdir -p ./target && go build -o target/solver cmd/sudoku_solver.go`

## run
`./target/solver --help`

## tested with
go version go1.12 darwin/amd64

## run tests
`go test ./...`

## input
```
0 4 0 0 0 0 1 7 9
0 0 2 0 0 8 0 5 4
0 0 6 0 0 5 0 0 8
0 8 0 0 7 0 9 1 0
0 5 0 0 9 0 0 3 0
0 1 9 0 6 0 0 4 0
3 0 0 4 0 0 7 0 0
5 7 0 1 0 0 2 0 0
9 2 8 0 0 0 0 6 0
```
Save the above to a file and pass it to the solver via
`./target/solver --file=/path/to/file.txt`