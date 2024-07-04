## css

### basic syntax

```
selector.selected {
    property: value;
}
```

universal selector: `*`
type selectors: `div`
class selectors: `.name`
id selectors: `#name`
grouping selector:
```
.read,
.unread {
    property: value;
}

.read {
...
.unread {
...
```
chaining selectors:
```
.class1.class2
.class#id
```
descendant combinator:
```
.grandparent .parent .child {
...
.parent .child {
...
```

### properties to get started with

#### color, background-color

`color` sets an element's text color
`background-color` sets an element's background color

they accept keywords (black, white, transparent), HEX, RGB & HSL values
[css legal color values](https://www.w3schools.com/cssref/css_colors_legal.php)

#### typography, text-align

`font-family` can be a single value or a csv to determine which font an element uses
    - font family name: `"Times New Roman"`
    - generic family name: `serif`

should always provide a fallback, e.g. `font-family: "Times New Roman", serif`

`font-size: 22px;`
`font-weight` sets boldness, can be a keyword `bold`, number between 1-1000
`text-align: center`

#### image height, width

```
img {
    height: auto;
    width: 500px;
}
```

## adding css to html

### external

in `head`, create a self closing `<link>` as `<link rel="stylesheet" href="styles.css">`

### internal

embedded within the html itself
placed directly within `<style>` tags, which are then placed inside the `<head>`

### inline

added to the html element itself
`<div style="color: white; background-color: black">...</div>`

## the cascade of css

there are default styles provided by browsers
the cascade is what determines which rules get applied to the html

there are different factors that the cascade uses to determine:

specificity - more specific declarations take precedence over less specific ones
    - inline has the highest specificity
    - each selector has it's own specificity level

specificity will only be taken into account when an element has multiple conflicting declarations

`id` > `class` > `type`

when there is no declaration with a selector of higher specificity, rules with greater #s of selectors of the same type takes precedence

index.html
```
<div class="main">
    <div class="list" id="subsection"></div>
</div>
```
style.css
```
#subsection {
    background-color: yellow;
    color: blue;
}

/* higher priority */
.main #subsection {
    color: red;
}
```

### inheritance

certain css properties are inherited by that element's descendants

typography-based properties are usually inherited, most others aren't

when directly targeting an element, inheritance is overwritten

### rule order

whichever rule was last defined has priority, if all other checks fail

## inspecting html & css

chrome devtools

## the box model

