# intro

## motivations and mindset

growth mindset: someone who believes they can get better at anything with effort+persistence

struggling = growth

### learning process

projects = the ultimate method for ensuring theoretical understanding == actuality

focus mode: conscious focus on learning, reading, watching, working
diffuse mode: subconscious focus on dishes, exercising, sleeping

here, your mind connects what it's been learning - where breakthroughs happen

take a break to refresh and let your subconcious work on making connections

understand, practice & teach it

## how the web works

computers connected via a network of undersea cables, locally/regionally connected via fiberoptic

# html, css & js

html: markup schema for basic structure
css: static style sheet language to improve visual styling
js: if html is structure, css is design, js is function

## html

### elements & tags

all elements on an html page are pieces of content wrapped in html tags

elements are containers for content

`<p>jello world</p>`
open/close tag indicate element, content in between

[mozilla html ref](https://developer.mozilla.org/en-US/docs/Web/HTML/Element)

### void elements

some html elements do not have a closing tag
aka `<br>`, `<img>`

void elements are known as such because they have nothing inside them

self-closing tags, aka `<br />` & `<img />`, are discouraged in the most recent version of HTML

[video on markdown](https://www.youtube.com/watch?v=LGQuIIv2RVA)

### boilerplate

`<!DOCTYPE html>`: HTML5
`<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN" "http://www.w3.org/TR/html4/loose.dtd">`: HTML4

```
<html lang="en">
</html>
```

#### head element

`<head>` is where the meta-info about the webpage, no elements that display content allowed

##### meta element

`<meta>` tag, always specify the charset encoding of the webpage, ie `<meta charset="utf-8">`

#### body element

always in html, under head

using vscode (ew) generation, this line is generated:
`<meta name="viewport" content="width=device-width, initial-scale=1.0">`

### working with text

`<p>`, `<h1-h6>`

`<strong>` makes text bold, semantically marking text as important which affect tools for those with visual impairments

`<em>` makes text italic, also semantically marking text as important

#### comments

`<!-- -->`

### lists

unordered:
```
<ul>
    <li></li>
</ul>
```

ordered:
```
<ol>
    <li></li>
</ol>
```

[shay howe on lists](https://learn.shayhowe.com/html-css/creating-lists/)

### links & images

to create links, use the anchor element `<a>`

however, there is no destination without using the `href` attribute

`<a href="www.samjtro.com">samjtro</a>`

#### open links in new tab

the `target` attribute specifies where the linked resource is to be opened

by default, it will take on the `_self` value - current tab

to open in a new tab/window, depending on browser settings, you can set it to `_blank`

`rel` attribute describes the relation between the current page and the linked doc

for opening links in new tabs, `rel` must be set to `"noopener noreferrer"`
    - `noopener`: prevents the opened link from ganing access to the webpage
    - `noreferrer`: prevents the opened link from knowing resource has a reference to it, including the noopener behaviour by default; thus, can be used by itself as well

for security: `noopener` prevents phishing attacks, `noreferrer` wish to not let the opened link know that your resource links to it

### absolute v relative links

#### absolute

`protocol://domain/path`

#### relative

protocol, domain implied
prepend `./` to relative links

##### metaphor

Think of your domain name (town.com) as a town, the directory in which your website is located (/museum) as a museum, and each page on your website as a room in the museum (/museum/movie_room.html and /museum/shops/coffee_shop.html). Relative links like ./shops/coffee_shop.html are directions from the current room (the museum movie room /museum/movie_room.html) to another room (the museum shop). Absolute links, on the other hand, are full directions including the protocol (https), domain name (town.com) and the path from that domain name (/museum/shops/coffee_shop.html): https://town.com/museum/shops/coffee_shop.html.

### images

`<img>` element displays an image; and is self closing
empty, self-closing html elements don't need a closing tag

`src` attribute works much like `href`, embedding an image using either absolute or relative paths

`alt` attribute is alt text

`height`, `width` are size attrs

[four main image types](https://internetingishard.netlify.app/html-and-css/links-and-images/#image-formats)
