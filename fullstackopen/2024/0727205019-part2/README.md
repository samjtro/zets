# a - rendering a collection, modules

## js arrays

going to be using functional operators of arrays

great [resource](https://www.youtube.com/playlist?list=PL0zVEGEvSaeEd9hlmCXrk5yUyqUag-n84)

## rendering collections

say we want to render an array of objects, with schema:
```
{
  id: 1,
  content: "HTML is easy",
  important: true
},
...
```
we could display them in a list manually, indexing each element, or:
`{notes.map(note => <li>{note.content}</li>)}` | `{notes.map(function(note) => {return ...})}`
`.map()` creates a new array, iterates over the old array, mapping each element with the function given as a parameter

this, however, generates a warning:
`each child in an array or iterator should have a unique "key" prop`
fixed via: `notes.map(note => <li key={note.id}...`

## anti-pattern; array indexes as keys

```
notes.map((note, i) =>
  <li key={i}>
    {note.content}
  </li>
```

is not reccomended

# b - forms

after storing notes in state, how do we access the form's input

through the use of [controlled components](https://react.dev/reference/react-dom/components/input#controlling-an-input-with-a-state-variable)

