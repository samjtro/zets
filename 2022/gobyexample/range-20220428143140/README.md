# range

range iterates over elements in a variety of data structures

0. range can be used to sum the numbers in a slice or array

1. range on arrays & slices provide the index & value for each entry
- if we don't need either we can ignore it with `_`

2. range on map iterates over k:v pairs

3. range can iterate over just k, or just v

```
for k,v := range map

// or

for k := range map

// or

for _,v := range map
```

4. range on strings iterates over unicode code points
- first val is the starting byte index of the rune, second val is the rune
- see [strings & runes](https://gobyexample.com/strings-and-runes)
