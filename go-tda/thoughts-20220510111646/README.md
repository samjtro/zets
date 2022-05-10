# why is there no good dataframe package for go?

i have been giving this a bit of thought lately; go kicks pythons ass in terms of data science speed and ability. the only thing python still excels at is ease of writing, obviously, but that doesn't take away from the very obvious advantages go has over python (more than speed/concurrency; native static typing, security, compilation, etc.). 

so why is there no decent dataframe package for go? go-gota is okay i suppose, but it is nowhere NEAR pandas in terms of usability, functionality, and general functionality.

for go-tda, i have decided to go for creating my own "dataframe" of sorts, by creating FRAME structs to put into a multi-dimensional slice formation. i believe this should effectively replicate dataframe functionality combined with the power of go structs, which i can make whatever i wish.
