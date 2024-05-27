# [Docker - Getting Started p3]()

## Updating our app

> implement the change from "empty text" to "You have no todo items yet! Add one above!"

### Updating the source code

> update src/static/js/app.js

0. Let's build the updated version

`docker build -t getting-started .`

1. Run the new container

`docker run -dp 3000:3000 getting-started`

```ERROR 
docker: Error response from daemon: driver failed programming external connectivity on endpoint quizzical_ellis (ee7f3ef0499194d66d8555cef5f5ad2ecabb71d98b3397f4d00a147f05c66b19): Bind for 0.0.0.0:3000 failed: port is already allocated.
```

The old container is still running!

### Removing the old container

0. Get the container ID utilizing:

`docker ps`

1. Stop the container:

`docker stop <id>`

2. Remove it:

`docker rm <id>`

### Start the new container

0. Run the updated app

`docker run -dp 3000:3000 getting-started`

*localhost:3000 should have updated*

### Recap

- When the update occurred, all of the existing items added to the list vanished
- We can make code updates without needing to rebuild/start anew after every change


