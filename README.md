# Go Simple Calculator

## Description
A simple command-line calculator in Go, supporting basic mathematical operations.

## Features
- Addition (+)
- Subtraction (-)
- Multiplication (*)
- Division (/)

## Prerequisites
- Go 1.23.2 or higher

## Installation

1. Clone the repository
```bash
git clone https://github.com/boysimon10/go-simple-calc.git
cd go-simple-calc
```

2. Initialize Go module
```bash
go mod init github.com/boysimon10/go-simple-calc
go mod tidy
```

## Usage
```bash
go run cmd/main.go
```

Example:
```
[number 1] [operator] [number 2]
5 + 3
```

## Project Structure
- `cmd/main.go`: Main entry point
- `pkg/calculator/`: Mathematical operations
- `pkg/handler/`: Input and calculation handling
- `pkg/utils/`: Utility functions
