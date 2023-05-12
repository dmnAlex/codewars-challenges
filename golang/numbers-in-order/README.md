# Are the numbers in order?

## Task

In this Kata, your function receives an array of integers as input. Your task is to determine whether the numbers are in ascending order. An array is said to be in ascending order if there are no two adjacent integers where the left integer exceeds the right integer in value.

For the purposes of this Kata, you may assume that all inputs are valid, i.e. arrays containing only integers.

Note that an array of 0 or 1 integer(s) is automatically considered to be sorted in ascending order since all (zero) adjacent pairs of integers satisfy the condition that the left integer does not exceed the right integer in value.

For example:

```go
InAscOrder([]int{1, 2, 4, 7, 19}) // returns True
InAscOrder([]int{1, 2, 3, 4, 5}) // returns True
InAscOrder([]int{1, 6, 10, 18, 2, 4, 20}) // returns False
InAscOrder([]int{9, 8, 7, 6, 5, 4, 3, 2, 1}) // returns False because the numbers are in DESCENDING order
```

## Run tests

Use `go test -v` to run the test
