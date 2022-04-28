# maps

maps are go's [associative data type](http://en.wikipedia.org/wiki/Associative_array); aka python's dict

0. to create an empty map use `make(map[key-type]val-type)`
- set: key/value pairs using name[key] = val syntax
- get: name[key]

1. printing a map (e.g. Println) will show all of its key/value pairs

2. len() returns the number of key/value pairs

3. delete() removes key/value pairs from a map

4. optional second return value when getting from a map indicates if the key was present
- this can be used to disambiguate between missing keys and keys with zero values
- we don't need the val itself, we ignore it using the _ identifier

5. we can declare and init a new map on the same line:

```
map[k-type]v-type{ k: v, ... }
```

6. maps appear in the form map[k:v ...] when printed with Println
