# [Docker -- Getting Started](https://docs.docker.com/get-started/)

Available on [Github](https://github.com/docker/getting-started)

## Getting Started

### First Command

`docker run -d -p 80:80 docker/getting-started`

0. -d = run the container in *detached mode* (in the background)
1. -p 80:80 = map port 80 of the host to port 80 in the container
2. docker/getting-started = the image to use

> -d -p can be shortened to -dp

### What is a Container?

**A container is another proccess on the machine that has been isolated from all other processes**

- That isolation leverages [kernel namespaces and cgroups]
	- These are the foundational elements of the container
	- [Medium - Demystifying Kernel Space](https://medium.com/@saschagrunert/demystifying-containers-part-i-kernel-space-2c53d6979504)
- Docker leverages these capabilities and makes them easy to use
- [Full 'Containers from Scratch' Video](https://www.youtube.com/watch?v=8fi7uSYlOdc)

### What is a Container Image?

- When running a container, it uses an isolated fs provided by a **container image**.
- Since the image contains the container's fs; it must contain everything needed to run the app
	- Dependencies
	- Configuration
	- Scripts
	- Binaries
	- Configuration for the container; env vars, default commands to run, other metadata
	- etc.

**Think of a container as an extended version of [chroot](https://man7.org/linux/man-pages/man1/chroot.1.html)**

- The fs is simply coming from the image
- A container adds additional isolation not available when simply using chroot

## Our Application

This tutorial relies on a node.js todo app running at localhost:3000

### Building the App's Container Image

```Dockerfile
FROM node:12-alpine
RUN apk add --no-cache python2 g++ make
WORKDIR /app
copy . .
RUN yarn install --production
CMD ["node", "src/index.js"]
```

`docker build -t getting-started .`

**Image Building**

- The previous command uses Dockerfile to build a new container image
- "layers" were downloaded as a result
	- We instructed the builder that we wanted to start from the node:12-alpine image
	- That image was subsequently downloaded
- After the download, we copied and used yarn to install the app's dependencies
- CMD directive specifies the default command to run when starting the container
- -t flags the image
	- the name is getting-started, which is then refered to when running a container
- the . tells Docker to look for the Dockerfile in the current directory

### Starting an App Container

**Running the app from the built image**

0. Start the container using docker run

`docker run -dp 3000:3000 getting-started`

1. open [localhost:3000](http://localhost:3000)

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

## Sharing our App

To share Docker images, you have to use a Docker registry


