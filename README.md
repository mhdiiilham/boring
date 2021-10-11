# BORING

# The Go Approach
Don't communicate by sharing memory, share memory by communicating.

ie: You don't have some blob of memory and then put lock and mutexes
and condition variable around it to protect from parallel access, instead
use the channel to pass the data back and forth between goroutine and make
concurrent program to operate that way.

# 1. Generator
A function that return a channel
