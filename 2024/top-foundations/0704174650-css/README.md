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

the most important skill to master from css is positioning and layout

### everything is a box

everything you see in a webpage is a rectangular box
other boxes can live inside of one another

using `padding`, `border` & `margin` we can manipulate their size and the space between them
    - `padding` increases the space between the margin and the padding
    - `border` adds space between the margin and the padding
    - `margin` increases the space between the borders of a box and the borders of adjacent boxes

[video](https://www.youtube.com/watch?v=HdZHcFWcAd8) and [article](https://developer.mozilla.org/en-US/docs/Learn/CSS/Building_blocks/The_box_model)

## block & inline

most of the elements thus far have been default style `display: block;`
block elements appear stack atop each other, starting on a new line

inline elements appear in line with whatever else they are placed beside
for example `<a>` is inline

inline-block elements behave like inline elements, but with block-style padding & margin

### divs and spans

css styling, positioning, etc.
div is block-level by default: divide the page into blocks
span is inline-level by default: group text content and inline html for styling, only when no other html element is appropriate

[normal flow](https://developer.mozilla.org/en-US/docs/Learn/CSS/CSS_layout/Normal_Flow)
[default block/inline](https://www.w3schools.com/html/html_blocks.asp)
[inline v inline-block](https://www.digitalocean.com/community/tutorials/css-display-inline-vs-inline-block)
[learn css layout](https://learnlayout.com/no-layout.html)
