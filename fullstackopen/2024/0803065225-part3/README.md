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

## about http request types

[http standard](https://www.rfc-editor.org/rfc/rfc9110.html#name-common-method-properties) talks about two props related to req types; safety and idempotency

safe; convention established that GET/HEAD methods should not have the significance of taking an action other than retrieval
- means that it must not cause any other side effects on the server

idempotent; N > 0 side effects of identical requests is the same for a single request
- should be the same no matter how many times the request is sent

express' `json-parser` is [middleware](https://expressjs.com/en/guide/using-middleware.html)
- funcs that are used for handling req/resp objects

json-parser takes raw data, parses to a js object and assigns it to the req as a new prop body
- can use several middlewares at the same time

# b - deploying app to the internet

copied frontend to the same dir, moved all part3 files to backend and attempted to initialize: encountered error:
- Access to XMLHttpRequest at 'http://localhost:3001/api/notes' from origin 'http://localhost:5173' has been blocked by CORS policy: No 'Access-Control-Allow-Origin' header is present on the requested resource.

## same origin policy and CORS

url origins are defined by the combination of protocol (scheme), hostname & port

```
http://example.com:80/index.html

protocol: http
host: example.com
port: 80
```

browser issues a request to the server, response contains refs to external assets hosted either on the same server, or a different one

if it was issued using the same url that the source html was fetched from, then the browser process the request. if not, browser checks the `Access-Control-Allow-Origin` header
- if it contains * in source, it processes the request, otherwise throwing an error

`same-origin policy` is a security mech implemented by browser to prevent session hijacking
to enable legitimate cross-origin requests, w3c came up with the CORS (cross-origin resource sharing) mechanism
- allows the webpage to embed cross-origin assets


