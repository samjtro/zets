# [Docker - Getting Started p2]()

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


