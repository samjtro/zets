# week 0

## representation

base-2 = binary: 0,1
    - electricity is there, or it's not
bit: a digit of binary, 0,1
transistor: a switch

binary math ex:

```
0 0: 0
0 1: 1
1 0: 2
1 1: 3
1 0 0: 4
1 0 1: 5
1 1 1: 7
1 0 0 0: 8
1 0 0 1: 9
1 0 1 1: 11
1 1 1 1: 15
1 0 0 0 0: 16
```

base-10: decimal
0   0  0
100 10 1
10^2 10^1 10^0

base-2: binary
0   0   0
4   2   1
2^2 2^1 2^0

byte: a bit

8-bit: 00000000 = 0
8-bit: 11111111 = 255 (256-1)

e.g.: A: 65: 01000001

ASCII: 8-bit
    - For other languages, we need more

Unicode: 8-32-bit
    - Superset of ASCII

11110000100111111001100010000010: crying laughing face emoji
U+1F602: base-16, hex

Represent color as RGB

72 73 33 = Yellow Color, Hi!, etc.
Depends on the context they are represented in

Pixel = 3 8-bit bytes

input -> algorithms -> output

Code is a computerized implementation of algorithms

Algorithm to search phonebook:

```
n: n*1

n/2: n*2

log2n: n/2
```

## pseudocode

log2n:

```
1. pick up phone book
2. open to middle of phone book
3. look at page
4. if person is on page
5.      call
6. else if person is earlier
7.      open to middle of left half of book
8.      go back to line 3
9. else if person is later
10.     open to middle of right half of book
11.     go back to line 3
12. else
13.     quit
```


