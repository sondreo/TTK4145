# Result of exercise 1 - part 4

### What is happening?

> Two threads are spawned and are being executed concurrently. One would think that
the global variable `i` should be counted up and down an equal amount of times so the end result should be zero. This is, however, not the case. Instead, the end sum of the variable is a seemingly random integer. The result differs between each execution of the programs.

### Why is this happening?

> The global variable is a shared recourse. Since there is not implemented any form of locking
mechanism to insure that only one thread works with the variable at a time, the result may differ.
The incrementing/decrementing is not an atomic operation, so after one of the threads has incremented the variable, the other could fetch the variable and decrement it before the first thread has been able to store the modified value. This means the value of the variable may be changed from the operation made by one thread, without it being registered by the other thread.
