# pointers

go supports [pointers]() allowing you to pass refs to vals and records within your program

0. during a traditional functional call, you may utilize parameters as variables to set and use as you wish; these are copys of the passed parameter

1. likewise, a function can accept a `*type`, or a pointer. the pointer in the subsequent code then dereferences the pointer from its memory address, assigning it to the value at the current address - changing the value at the referenced address

2. `&i` gives the memory address of i - pointer to i

3. zeroval doesn't change the i in main, but zeroptr does because it has a ref to the mem address for that var
