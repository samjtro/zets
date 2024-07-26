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

