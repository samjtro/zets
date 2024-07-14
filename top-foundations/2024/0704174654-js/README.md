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

# understanding errors

[what went wrong](https://developer.mozilla.org/en-US/docs/Learn/JavaScript/First_steps/What_went_wrong)

# DOM manipulation & events

DOM == `document object model`; a tree-like rep of contents of a webpage

```
<div id="1">
  <div class="2"></div>
  <div class="3"></div>
</div>
```

use selectors to grab nodes to work with. to refer to `2`, you can use css-style selectors

`div.2`, `.2`, `#1 > .2` & `div#container > div.2`

you can also use relational selectors, e.g. `firstElementChild` or `lastElementChild`

when HTML is parsed by a browser, it is converted to the DOM. these nodes are JS objects that have props and methods, which are the primary tools we use to manipulate our webpage

## query selectors

`element.querySelector(selector)` - returns a reference to the4 first match of selector
`element.querySelectorAll(selectors` - returns a 'NodeList' containing references to all of the matches of the selectors
  - `NodeList` != array

convert NodeList -> array using `Array.from()` or the spread operator

## element creation

`document.createElement(tagName, [options])` creates a new element of tag type tagName in memory; you manipulate the element, and then place the element into DOM with one of the following methods

## append elements

`parentNode.appendChild(childNode)`
`parentNode.insertBefore(newNode, referenceNode)` into parentNode before ref

## remove elements

`parentNode.removeChild(child)`

## altering elements

`add`, `remove` or `alter` element properties

`let div = document.createElement('div');` create element, then:

`div.style.color = "blue";` use a specific setter
`div.style.cssText = "color: blue; background: white;";` add several style rules under a specific setter
`div.setAttribute("style", "color: blue; background: white;");` add several style rules under the generic setter

when accessing a kebab-cased CSS prop, aka `background-color`, use camelcase with dot or bracket notation
bracket notation can utilise either camel or kebab case, but the property must be a string

`dot notation, kebab case: doesn't work`
`div.style.background-color`, equivalent to `div.style.background - color`

`dot notation, camel case: works`
`div.style.backgroundColor`

`bracket notation, kebab case: works`
`div.style["background-color"]`

`bracket notation, camel case: works`
`div.style["backgroundColor"]`

## editing attrs

`div.setAttribute("id", "theDiv")` if `id` exists, update to `theDiv`; else, create `id` with value `theDiv`

`div.getAttribute("id")` returns value of specified attr, e.g. `theDiv`

`div.removeAttribute("id")` removes specified attr

[mdn on html attrs](https://developer.mozilla.org/en-US/docs/Web/HTML/Attributes)

## working with classes

`div.classList.add("new")` adds class "new" to the div

`div.classList.remove("new")` removes class "new" from the div

`div.classList.toggle("new")` add/remove class "new" from the div

standard (and easier) to toggle rather than add/remove

## adding text content

`div.textContent = "hello,world";` creates a text node and inserts in div

## adding html content

`div.innerHTML = "<span>hello</span>";` renders the html inside div

*textContent is preferred over innerHTML, potential css vuln

all of these methods are used by the browser at runtime to modify the DOM, not the html file

# events

events make DOM manipulation happen dynamically - 3 ways to do this:
  - function attrs directly in HTML
  - props in the form of `on<eventType` such as `onClick` or `onmousedown` in DOM nodes
  - event listeners attached to DOM nodes

event listeners are preferred, but all 3 are used

create a button that alerts "hello" when clicked

## method 1

`<button onclick="alert('hello')>click</button>`

## method 2

html:
`<button id="btn">click</button>`

js:
```
const btn = document.querySelector("#btn");
btn.onclick = () => alert("hello");
```

## method 3

html:
`<button id="btn">click</button>`

js:
```
const btn = document.querySelector("#btn");
btn.addEventListener("click", () => {
  alert("hello");
});
```

named function implementation: [here](https://www.theodinproject.com/lessons/foundations-dom-manipulation-and-events#exercise)

we can pass the event as a param to the func

```
btn.addEventListener("click", function (e) {
  console.log(e);
});
```

`e` contains an object with a ref to the `event`
specify components of event, e.g. `e.target`
or change style of target, e.g. `e.target.style.background = "blue";`

## attaching listeners to groups of nodes

use `querySelectorAll('')` to get a list of all components with the selector, then iterate thru list

```
const buttons = document.querySelectorAll("button");

buttons.forEach((button) => {
  button.addEventListener("click", () => {
    alert(button.id);
  });
});
```

## other useful events

- `click`
- `dblclick`
- `keydown`
- `keyup`

[event ref page - w3schools](https://www.w3schools.com/jsref/dom_obj_event.asp)

[additional resources](https://www.theodinproject.com/lessons/foundations-dom-manipulation-and-events#assignment)

# objects - javascript.info

object can be created with figure brackets and an optional list of properties

properties are `k:v` pairs - `key` is a string, `value` can be anything

```
let user = new Object(); // object constructor
let user = {};           // object literal
```

```
let user = {
  name: "John",
  age: 30
};

user.isAdmin = true;
delete user.age;
```

multiword prop names must be quoted
```
let user = {
  ...
  "is admin": true,
};
```

## square brackets

for multiword props, dot access doesn't work - instead, using bracket notation

```
user["is admin"] = true;
```

```
let user = {
  name: "sam",
  age: 24,
};

let key = prompt("what do you want to know", "name");
alert( user[key] ); //if enter "name", John
```

## computed properties

you can use square brackets in an object literal when creating an object, called computed properties

```
let fruit = prompt("fruit? ", "apple");

let bag = {
  [fruit]: 5,
};

alert(bag.apple); //if fruit="apple", 5
```

==

```
let fruit = prompt("fruit? ", "apple");
let bag = {};

bag[fruit] = 5;
```

can use complex expressions inside square brackets:

```
let fruit = 'apple';
let bag = {
  [fruit + 'Computers']: 5 // bag.appleComputers = 5
};
```

we can use existing vars as vals for prop names

```
function makeUser(name, age) {
  return {
    name: name,
    age: age,
  };
}

let user = makeUser("john", 30);
alert(user.name);
```

shorthand `name:name` == `name`, e.g.

```
...
  return {
    name,
    age,
  };
...
```

## prop name limitations

vars cannot be named language-reserved words `for`, `let`, etc. - for objects, there's no such restriction

other types are converted to strings when used as prop keys

`0` becomes `"0"`

special prop `__proto__` can be set to a non-object val

## property existence test, "in"

`object.key == undefined` == `"key" in object`

if the key's val is undefined, we still want it to return "false" if key doesn't exist. thus, we use the `in` operator to avoid this pitfall

## walk over all keys, for...in

```
for (key in object) {
  ...
}
```

if keys are ints, they are sorted in ascending order. else, they are listed in creation order

make a non-int by adding non-int before/after

# objects - the odin project

```
function sumOfTripledEvens(arr) {
  let tripledEven = 0;
  for(let i = 0; i < arr.length; i++) {
    if(arr[i]%2 == 0) {
      tripledEven += arr[i] * 3;
    }
  }
  return tripledEven;
}
```

## map method

`map` function expects a callback as an arg
returns a new array, running a for,in loop and executing the passed function

```
function addOne(num) {
  return num + 1;
}
const arr = [1,2,3,4,5];
const mappedArr = arr.map(addOne);
console.log(mappedArr);
```

can also do this inline

```
...
const mappedArr = arr.map((num) => num + 1);
...
```

## filter method

similar to map but rather than transforming the vals in the arr, it returns the original vals only if the callback returns true

e.g., `isOdd` returns `true`/`false`

`filter` expects `callback` to return `true`/`false`

e.g. if we wanted to remove all even nums from `[1,2,3,4,5]`

```
function isOdd(num) {
...
const oddNums = arr.filter(isOdd);
```

## reduce method

expects cb, which takes two args
  - `accumulator`: current val of the result at that point in the loop, set to either the `initialValue` or if none is provided, the first element
  - `initialValue`: starts from 0 [1st element] if none provided, otherwise, if we wanted to add 10

```
...
const productOfAllNums = arr.reduce((total, currentItem) => {
  return total * currentItem;
}, 1);
```

`reduce` goes through every element in arr applying cb, then changes total without changing the array and returns total

# rewrite sum of tripled evens

```
function sumTripledEvens(arr) {
  const evens = arr.filter(arr%2 == 0);
  const tripledEvens = arr.map((num) => num * 3);
  return arr.reduce((total, currentItem) => return total + currentItem);
}
```

simplified as

```
...
  return arr
    .filter((num) => num % 2 === 0)
    .map((num) => num * 3)
    .reduce((acc, curr) => acc + curr);
}
```
