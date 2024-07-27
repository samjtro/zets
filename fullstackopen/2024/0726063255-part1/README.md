# a - intro to react

## jsx compilation

```
const App = () => {
    const now = new Date()
    const a = 10
    const b = 20
    console.log(now, a+b)

    return (
        <div>
          <p>Hello world, it is {now.toString()}</p>
          <p>
            {a} plus {b} is {a + b}
          </p>
        </div>
    )
}
```

compiled by babel to:

```
const App = () => {
    const now = new Date()
    const a = 10
    const b = 20
    return React.createElement(
        'div',
        null,
        React.createElement(
            'p', null, 'Hello world, it is ', now.toString()
        ),
        React.createElement(
            'p', null, a, ' plus ', b, ' is ', a + b
        )
    )
}
```

## props: passing data to components

```
const Hello = (props) => {
    return (
        <div>
            <p>Hello, {props.name}</p>
        </div>
    )
}
```

The function defining the component now has param props, as an argument, receives an object of an arbitrary number of props with fields corresponding to user's defs

## root

Returns in React are usually rooted in a div

```
return (
    <div>
    ...
```

Alternatively:

```
return (
    <Element />
    ...
```

Adds divs automatically, making it messy. Alternatively:

```
return (
    <>
    ...
    </>
```

Has the same effect as wrapping in root

## do not render

# b - javascript

## intro

official name for JS spec is ECMAScript - [latest](https://262.ecma-international.org/)

## vars, arrays, objects, funcs

## object methods, this

assign methods to an object by defining props that are funcs

```
const arto = {
    ...
    greet: function() {
        console.log('hello, ' + this.name)
    ...
...
arto.greet()
```

methods can be assigned after creation

```
arto.growOlder = function() {
    this.age += 1
}
```

in js the value of this is defined based on how the method is called; thus, when attempting to run a method through a reference, e.g. `const refToGreet = arto.greet` you run into an error undefined

`bind` function is a potential fix; e.g. `arto.greet.bind(arto)`, this creates a new function where this is bound to point to Arto, independent of where and how the method is called

arrow funcs are also a possible fix to some of the related problems, but it shouldn't be used as a method for objects because then this won't work at all

## not classes

```
class Person {
    constructor(name, age) {
        this.name = name
        this.age = age
    }
    greet() {
        console.log('hello, ' + this.name)
    }
}

const sam = new Person('Sam', 23)
sam.greet()
```

resources:
- https://github.com/getify/You-Dont-Know-JS
- https://javascript.info/
- https://eloquentjavascript.net/
- https://www.youtube.com/playlist?list=PLlasXeu85E9cQ32gLCvAvr9vNaUccPVNP
- https://egghead.io/
- https://developer.mozilla.org/en-US/docs/Web/JavaScript

# c - component state, event handlers

## component helpers

```
const Hello = (props) => {
    const brnYear = () => return new Date().getFullYear() - props.age
```

## destructuring

indirect:

```
...
    ...
    const {name, age} = props
```

direct:

```
const Hello = ({name, age}) => {
    ...
```

## re-rendering

use the `refresh()` func

## stateful components

`const [counter, setCounter] = useState(0)`

## event handling

`<button onClick={() => setCounter(counter + 1)}>plus</button>`

remember, event handlers are functions, so you should define them as such

## passing state - to child components

lift state up - https://react.dev/learn/sharing-state-between-components

# more complex state, debugging react apps

## update of the state is asyncronous

## rules of hooks


