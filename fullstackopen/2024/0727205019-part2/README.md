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

# c - getting data from a server

data was fetched manually in part 0

```
const xhttp = new XMLHttpRequest()

xhttp.onreadystatechange = function() {
  if (this.readyState == 4 && this.status == 200) {
    const data = JSON.parse(this.responseText)
      // handle the response that is saved in variable data
    }
  }

xhttp.open('GET', '/data.json', true)
xhttp.send()
```

js engines or follow an [asynchronous model](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Event_loop); this requires all [io operations](https://en.wikipedia.org/wiki/Input/output) to be executed as non-blocking, meaning code execution happens immediately after calling an io func
when an asynchronous op is completed, the js engine calls the event handlers registered to the operation
js engines are single-threaded: the non-blocking model is required, otherwise the browser would freeze for each request

[web workers can enable parallelized code in js](https://developer.mozilla.org/en-US/docs/Web/API/Web_Workers_API/Using_web_workers), but still handled by a [single thread](https://medium.com/techtrument/multithreading-javascript-46156179cf9a)

## npm, axios & promises

could use fetch, but we will use axios

add devel devependencies `npm install package --save-dev`

axios' `get` method returns a [promise](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Using_promises)

`a promise is an object representing an async op, can have 3 states: pending, fulfilled, rejected`

problematic data call in main.jsx
```
axios
  .get('http://localhost:3001/notes')
  .then(response => {
    ReactDOM.createRoot(document.getElementById('root')).render(
      <App notes={response.data} />
    )
  })
```

### effect hooks

effects let a component connect to and synx with external systems

# d - altering data in server

## rest

using axios, check `./example/db.js`

individual data objects are resouces, with its own url; e.g. `/notes/3`

use `.catch` after `.then` to handle errors on exec

# e - adding styles to react apps

[refresher](https://developer.mozilla.org/en-US/docs/Learn/Getting_started_with_the_web/CSS_basics)

selector: declaration

## inline styles

aka the use of the `<style>` tag
