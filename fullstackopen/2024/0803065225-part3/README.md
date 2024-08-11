# a - node.js & express

## express, web servers

access web servers via http methods. can use express to create a minimal js web server (minimal for js**)

methods defined via `express().method(path, callback(request, response){})`
- use `response.send(value)` to send plain text back through response
- use `response.json(value)` to send json back through response

`response.end()`

## nodemon

watches files in dir for changes to update rather than having to restart server

## REST 

first introduced in [this paper](https://ics.uci.edu/~fielding/pubs/dissertation/rest_arch_style.htm)
- took notes on the paper [here](https://github.com/samjtro/zets/tree/main/rest/2024/0810155149-dissertation/README.md)
however, paper is not neccesarily as important as the more narrow idea of how restful apis are typically understood in web apps

resources = singular things in restful thinking, has unique url w resource address
convention: resource/id; e.g. notes/10

this is how we define a uniform interface
sometimes this is all referred to more simply as crud or https://en.wikipedia.org/wiki/Resource-oriented_architecture

## fetching a single resource

create a route to the resource via `express().get('path/:id', (req, resp) => {})`
can use `[].find(obj => obj.id === req.params.id` to find the resource with the specified id
`resp.json` still returns json, whereas you can send 404s with `resp.status(statusCode).end()`
you can additionally modify the status message by using `resp.statusMessage = ''`

