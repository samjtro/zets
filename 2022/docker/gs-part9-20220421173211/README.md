# Image Building Best Practices

## Security Scanning

When you have built an image; it is good practice to scan it for security vulnerabilities using `docker scan`
- [Snyk](http://snyk.io/)

`docker scan <img_name>`

You can configure the scan using [commands](https://docs.docker.com/engine/scan/)
- As well, you can [configure Docker Hub](https://docs.docker.com/docker-hub/vulnerability-scanning/) to automatically scan newly pushed images automatically

## Image Layering

Using `docker image history <img_name>`, you can view the layers in the selected image as well as the command used to create each layer

- Using `docker image history getting-started` displays:

```
IMAGE          CREATED        CREATED BY                                      SIZE      COMMENT
de7f4001f055   12 hours ago   /bin/sh -c #(nop)  CMD ["node" "src/index.js…   0B
e7943b319923   12 hours ago   /bin/sh -c yarn install --production            83.8MB
95e06ce4d8b8   12 hours ago   /bin/sh -c #(nop) COPY dir:91bbd9c5da8507ee2…   58.1MB
4702ac55fc63   24 hours ago   /bin/sh -c #(nop) WORKDIR /app                  0B
0c749ad82487   24 hours ago   /bin/sh -c apk add --no-cache python2 g++ ma…   223MB
bb6d28039b8c   10 days ago    /bin/sh -c #(nop)  CMD ["node"]                 0B
<missing>      10 days ago    /bin/sh -c #(nop)  ENTRYPOINT ["docker-entry…   0B
<missing>      10 days ago    /bin/sh -c #(nop) COPY file:4d192565a7220e13…   388B
<missing>      10 days ago    /bin/sh -c apk add --no-cache --virtual .bui…   7.84MB
<missing>      10 days ago    /bin/sh -c #(nop)  ENV YARN_VERSION=1.22.18     0B
<missing>      10 days ago    /bin/sh -c addgroup -g 1000 node     && addu…   77.6MB
<missing>      10 days ago    /bin/sh -c #(nop)  ENV NODE_VERSION=12.22.12    0B
<missing>      2 weeks ago    /bin/sh -c #(nop)  CMD ["/bin/sh"]              0B
<missing>      2 weeks ago    /bin/sh -c #(nop) ADD file:5d673d25da3a14ce1…   5.57MB
```

Each line represents a new layer; the display shows the base at the bottom and newest on top
- You can easily then diagnose large images and quickly view the size of each layer

`--no-trunc` displays truncated layers

## Layer Caching

> Once a layer changes, all downstream layers **must** be recreated as well

Previous dockerfile:

```
FROM node:12-alpine
WORKDIR /app
COPY . .
RUN yarn install --production
CMD ["node", "src/index.js"]
```

Each command in the Dockerfile becomes a new layer in the image
- Yarn dependencies must be rebuilt around new layers

To fix shipping the same dependencies every time we build, we can restructure the Dockerfile to support dependency caching
- Node-based apps define dependencies in package.json
- We can copy that file first, install dependencies, then copy everything else
- In that way we only recreate the yarn dependencies if there was a change to the package.json

0. Update the Dockerfile to copy package.json --> Instal dependencies --> Copy everything else

```
FROM node:12-alpine
WORKDIR /app
COPY package.json yarn.lock ./
RUN yarn install --production
COPY . .
CMD ["node", "src/index.js"]
```

1. Create a file named .dockerignore in the same file as the Dockerfile with

```
node_modules
```

.dockerignore files selectively copy only image relevant files [Read More](https://docs.docker.com/engine/reference/builder/#dockerignore-file)
- In this case, `node_modules` is omitted in the second COPY step because otherwise it would overwrite files created by the command in RUN
- [More info](https://nodejs.org/en/docs/guides/nodejs-docker-webapp/) on dockerizing Node.js apps

2. Then, build the new image

`docker build -t <img_name>`

3. After making changes to index.html, rebuilding the output should change
- It will say "using cache" for steps 1-4 of the dockerfile, and build mach faster using components from cache

## Multi-Stage Builds

Multi-stage builds can be an incredibly powerful tool to use multiple stages to create an image; advantages include
- Seperating build-time dependencies from runtime dependencies
- Reduce overall image sie by shipping only what your app needs to run

#### Maven/Tomcat Example

When building Java apps, a JDK is needed to compile the source code to Java bytecode
- The JDK isn't subsequently needed in prod -- Or, Maven/Gradle

```
FROM maven AS build
WORKDIR /app
COPY . .
RUN mvn package

FROM tomcat
COPY --from=build /app/target/file.war /usr/local/tomcat/webapps
``` 

We use the maven build stage to perform the actual java build
- Then, FROM tomcat copies the files from the build stage and creates the image from those files

#### React Example

React apps need Node environments to compile JS code, SASS stylesheets and more into static HTML/JS/CSS
- Not using server-side rendering, we don't even need Node for the production build
- We can ship these static resources in a static nginx container

```
FROM node:12 AS build
WORKDIR /app
COPY package* yarn.lock ./
RUN yarn install
COPY public ./public
COPY src ./src
RUN yarn run build

FROM nginx:alpine
COPY --from=build /app/build /usr/share/nginx/html
```

We use node:12 to build, and then copy the output to an nginx container

