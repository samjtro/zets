# closures

go supports [anonymous functions](http://en.wikipedia.org/wiki/Anonymous_function)
- these can form [closures](http://en.wikipedia.org/wiki/Closure_(computer_science))
- anonymous functions are useful when you want to define a function inline without naming it

0. functions that return another function can be defined anymously in the body of the parent
- the returned function closes over the variable returned to form a closure

```
func parent() func() int {
    i := 0
    return func() int {
        i++
        return i
    }

}
```

0. call the parent function (assigned to a result var)
- this function value captures its own return value which is then updated each successive call

1. that state is unique to the function
