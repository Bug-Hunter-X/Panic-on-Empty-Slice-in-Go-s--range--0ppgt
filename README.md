# Go Range on Empty Slice

This example demonstrates a common error in Go: using `range` on an empty slice without proper error handling. Accessing indices of an empty slice during iteration will result in a runtime panic.

The `bug.go` file contains a function `myFunc` which demonstrates the error. The `bugSolution.go` contains corrected function.