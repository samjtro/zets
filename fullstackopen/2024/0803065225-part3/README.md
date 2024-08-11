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
