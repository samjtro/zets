# slices

slices are a key data type in go
more powerful interface to sequences than arrays
slices are typed only by the elements they contain (not num of elements)
blog post [re: design+implementation](http://blog.golang.org/2011/01/go-slices-usage-and-internals.html) of slices in go

0. create an empty slice with non-zero length using make()
1. append returns a slice containing one or more new values accepted by var
2. slices can be copy'd; first by creating empy slice c (with len == s) and then using copy(new,old)
3. slices also support slicing, slice[lo:hi]
4. slices can also be multi-dimensional
5. the len of the inner slice my vary, unlike with multi-d arrays

