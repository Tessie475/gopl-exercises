# gopl-exercises

My solutions to the exercises from [*The Go Programming Language*](https://www.gopl.io/) by Alan A. A. Donovan and Brian W. Kernighan, worked through as I learn Go.

## Layout

Solutions are grouped by chapter, and each exercise lives in its own directory with a runnable `package main`:

```
ch01/
  ex1.1/   echo, also printing os.Args[0] (the command name)
  ex1.2/   echo, printing the index and value of each argument, one per line
```

## Running an exercise

Each exercise is a standalone program. From the repo root:

```sh
go run ./ch01/ex1.1 hello world
go run ./ch01/ex1.2 hello world
```

## Progress

| Exercise | Description | Status |
| -------- | ----------- | ------ |
| 1.1 | Modify `echo` to also print `os.Args[0]` | Done |
| 1.2 | Modify `echo` to print the index and value of each argument, one per line | Done |
