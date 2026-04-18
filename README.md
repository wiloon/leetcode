# LeetCode Practice

Personal Go solutions for LeetCode algorithm problems.

## Tech Stack

- **Language**: Go 1.26.2
- **Testing**: Go standard `testing` package (unit tests + benchmarks)

## Project Structure

Each problem is implemented in a single `.go` file, with a corresponding `_test.go` for unit tests and optionally a `_benchmark_test.go` for benchmarks.

```
leetcode/
├── running-sum-of-1d-array.go
├── running-sum-of-1d-array_test.go
└── running-sum-of-1d-array_benchmark_test.go
```

## Run Tests

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run benchmarks
go test -bench=. ./...
```

## Problems

| # | Title | File |
|---|-------|------|
| 1480 | Running Sum of 1d Array | [running-sum-of-1d-array.go](running-sum-of-1d-array.go) |
