# variables

`let` for vars, `const` for constants

# data types and conditionals

## strings
```
let string = "test";
 string2 = 'test';
 string3 = `test`;
```

\`\` are template literals, you can embed js and declare template literals over multiple lines

use `${}` to embed js in template literals

concatenation ex:
```
let one = 'test';
let two = `${one}`;
```

### concatenation in context

html:
```
<button>press</button>
<div id="greeting"></div>
```

js:
```
const button = document.querySelector("button");

function greet() {
    const name = prompt("name: ");
    const greeting = document.querySelector("#greeting");
    greeting.textContext = `hello ${name}`
}

button.addEventListener("click", greet);
```

you can only use `${}` in template literals; you can concatenate normal strings using `+` but template literals are more readable

### multiline strings

template literals respect line breaks in source code

### numbers as strings

you can use numbers as strings; built-in `Number()` function converts string -> numbers

## if... else...

bools in js are `true` `false`

strings are evaluated on a lexicographical, dictionary, order

## conditionals

`if (condition) { then } else if (condition) { then } else { then }`

### conditional operator: ?

`let result = condition ? val1 : val2`

if it's truthy, `val1` is returned; else, `val2`

`if (age > 18) { accessAllowed=true; } else { accessAllowed=false; }` == `let accessAllowed = (age > 18) ? true : false;`

#### multiple ?s

sequences of `?` operators return a value that depends on a branch of conditions:
```
let age = prompt('age: ', 18);
let message = (age < 3) ? 'goo goo ga ga' :
  (age > 18) ? 'hi' :
  (age < 100) ? 'greetings' :
  'unusual!';

alert(message);
```

#### non-traditional ?

can be used as a replacement for `if`

```
let co = prompt('which co created js: ');
(company == 'netscape') ?
  alert('correct!') : alert('false');
```

## switch

```
switch(x) {
  case'val':
    ...
    break;
  default:
    ...
    break;
}
```

### grouping of case

```
...
  case 3:
  case 4:
    ...
    break;
...
```

## js dev tools

[js dev tools](https://www.theodinproject.com/lessons/foundations-javascript-developer-tools)

## js call stack

[js call stack](https://www.javascripttutorial.net/javascript-call-stack/)

## functions

### built-in browser funcs

```
let i = "i am a string";
let i2 = i.replace("string", "sausage");
```

```
let arr = ["hi", "there"]
let arrString = arr.join(" ");
```

```
let num = Math.random();
```
### optional params

specify default args for optional parameters with `=`

### anon/arrow funcs

anon
```
(function () {
  alert("hi");
});
```

to run code when the user types in text box, call addEventListener(), which takes two params
  - name of event to listen for
  - function to run when event happens

when the user presses a key, the browser will call the function you provided and pass it a param with info about the event

without anon
```
function logKey(event) {
  console.log(`you pressed "${event.key}".`);
}

textBox.addEventListener("keydown", logKey);
```

anon
```
textbox.addEventListener("keydown", function (event) {
  console.log(`you pressed "${event.key}".`);
});
```

arrow
```
textBox.addEventListener("keydown", (event) => {
  ...
});
```

if the function contains only one line thats a return statement, you can omit the braces and `return`

```
let i = [1,2,3];
let i2 = i.map(item => item * 2);
console.log(i2) //2,4,6
```

above, `item => item * 2` ==
```
function (item) {
  return item * 2;
}
```

e.g.
```
textBox.addEventListener("keydown" (event) =>
  console.log(`you pressed: ${event.key}`);
);
```

# problem solving

from `think like a programmer`: `Problem solving is writing an original program that performs a particular set of tasks and meets all stated constraints.`

1. understand the problem
2. plan
3. *pseudocode
4. divide and conquer

## solving fizz buzz

```
```
