# debug building

It's test golang program for debug building. The program is creating shopping cart, and add/remove item from it. 

You can build debug code with `-tags debug`. In debug build, the program will print the errors to console. 

The debug output
```BASH
$ go build -tags debug -o d
$ ./d
Added item apple 1
Out of stock
Num is smaller zero
Added item banana 2
Item not found
Item removed banana
Added item banana 2
```

The original output (should be empty)
```BASH
$ go build -o a
$ ./a
```
