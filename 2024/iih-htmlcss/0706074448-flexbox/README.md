# tl;dr


- Use `display: flex;` to create a flex container.
- Use `justify-content` to define the horizontal alignment of items.
- Use `align-items` to define the vertical alignment of items.
- Use `flex-direction` if you need columns instead of rows.
- Use the `row-reverse` or `column-reverse` values to flip item order.
- Use `order` to customize the order of individual elements.
- Use `align-self` to vertically align individual items.
- Use `flex` to create flexible boxes that can stretch and shrink.

# overview of flexbox

already covered in odin.

# aligning a flex item

inside of the flex container, use `justify-content: center;` - other options include `flex-start`, `flex-end`, `space-around` & `space-between`
this has the same effect as adding the `margin: 0 auto;` declaration to the flex item - however, manipulating items through their containers is a common theme in flexbox

# distributing multiple flex items

`space-around` & `space-between` allows you to distribute items equally within a container - either equally based on the space around, or the space between the items

![space-around space-between diagram](https://internetingishard.netlify.app/flex-justify-content-distribution-b0ee9c.3c71bf1f.png)

# grouping flex items

flex containers position elements one level deep - grouping flex items is a powerful tool

# cross-axis (vertical) alignment

flex containers can define the vertical alignment of their items

using `align-items: center` vertically centers the items of the div styled - available items include `flex-start` (top), `flex-end` (bot), `stretch` & `baseline`

![align-items diagram](https://internetingishard.netlify.app/flex-align-items-26abfd.9d4b350a.png)

# wrapping flex items

flexbox is a more powerful alternative to [float-based grids](https://internetingishard.netlify.app/html-and-css/floats/#floats-for-grids)

creating grids requires using `flex-wrap` properly

![flex-wrap diagram](https://internetingishard.netlify.app/flex-wrap-b960c1.73a3247a.png)

# flex container direction

`flex-direction` refers to whether a container renders its items horizontally `row` or vertically `column`
default: horizontal (mostly)

## alignment considerations

when rotating `direction`, you also rotate the baseline for `justify-content`

![direction diagram](https://internetingishard.netlify.app/flex-direction-axes-b30e85.d1bca75a.png)

aka - `justify-content` == `align-items` for `column` `direction`

# flex container order

`flex-direction` gives control over the order in which items appear via `row-reverse` & `column-reverse`

`column-reverse` can be especially useful for mobile layouts

# flex item order

adding an `order` property to a flex item defines its order without impacting the surrounding item

default val: 0, increase=right, decrease=left

align individually with `align-self`, takes the same as `align-items`

# flexible items

all of the items thus far have either fixed- or content-defined- widths

the point of flexbox is the flexible part

`flex` property defines the width of individual items in a flex container

## static item widths

we can combine flex and static width box; `flex: initial;` falls back to the element's explicit `width` property

# auto-margins

alternative to an extra `<div>` when aligning a group of items left/right

auto-margins eat up all extra space in a container
