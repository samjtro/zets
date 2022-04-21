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

> Most popular is [Docker Hub](hub.docker.com)

**Create a repo**

- If you haven't, login to Docker Hub:

`docker login -u <user>`

- Then, tag the image with the name of the Repo

`docker tag <img_name> <name>/<repo_name>`

- Now, push the container to Docker Hub:

`docker push <user>/<repo_name>`

### Running the shared Image on a New Instance

- Login to [Play with Docker](https://labs.play-with-docker.com/)
- Click "+ ADD NEW INSTANCE"
- Run the app:

`docker run -dp 3000:3000 <name>/<repo_name>`

## Persistence

Why does our data *not* persist normally?

### The Container's fs

When a container runs, it uses those layers from an image for its fs
- Each container gets "scratch space", to create/update/remove files
- Any changes made in one container using the same image won't be seen in another

Start an ubuntu container
- Creating a file named "/data.txt" 
- With a random number between 1-10000

`docker run -d ubuntu bash -c "shuf -i 1-10000 -n 1 -o /data.txt && tail -f /dev/null"`

To view the number:

`docker exec <id> cat /data.txt`

On another Ubuntu container, that '/data.txt' file will be gone.

> To **Force End** a running container, use `docker rm -f`

### Container Volumes

[Volumes](https://docs.docker.com/storage/volumes/) provide the ability to connect spec fs paths of the container with the host machine
- Container directory is mounted == changes in that dir are seen on the host machine
- If we mount the same dir across container restarts, we are able to see the same files

2 main types of volumes; named volumes & bind points

### Persisting data on the todo app

By default the todo app stores data in a SQLite db located at `/etc/todos/todo.db`

Because SQLite is a relational db that stores its data in a single file, we can persist that file on the host and make it available to the next container
- By creating a vol and attaching it to the dir the data is stored in

Going to use a Named Volume -- buckets of data
- Maintains the physical location of data on disk referenced by the name of the volume

0. Create a volume

`docker volume create todo-db`

1. Stop the container if neccesary, and start a new container adding the -v flag
- Specifies volume mount, name it and mount to `/etc/todos/` which captures files created at the path

`docker run -dp 3000:3000 -v todo-db:/etc/todos getting-started`

When restarting the container now, data persists in the mounted dir

### Diving into the volume

When using a named volume, you can use `docker volume inspect <vol_name>` to view where that data is actually stored

```
[
    {
        "CreatedAt": "2022-04-21T08:32:38-06:00",
        "Driver": "local",
        "Labels": null,
        "Mountpoint": "/var/lib/docker/volumes/todo-db/_data",
        "Name": "todo-db",
        "Options": null,
        "Scope": "local"
    }
]
```

`Mountpoint` is the location on disk where data is stored

## Using Bind Mounts

Named Volumes persist and are good for cases where we don't have to worry about where the data is stored
- With **bind points**, we control the exact host mountpoint
	- Can use this to provide additional data into containers & persist data
	- ie Mount our source code into the container to see if the code changes, responds, and lets us see changes immediately

With node [nodemon](https://npmjs.com/package/nodemon) is a good tool

### Quick Volume Type Comparisons

Bind Mounts / Named Volumes are the two types of volumes; there are additional volume drive, eg:
0. SFTP
1. Ceph
2. NetApp
3. S3, etc.

						Named Volumes		Bind Points

Host Location					Docker chooses		You control
Mount Example (-v)				my-vol:/usr/local/data	/path/to/data:/usr/local/data
Popultes new volume with container contents	yes			no
Supports Volume Drivers				yes			no

### Dev-mode Containers

0. Mount src code to container
1. Install all dependencies (including "dev" pkg)
2. Start nodemon to watch fs changes

**Starting a dev-mode container implementation**

0. cd into app src code
1. Run this command:

```
docker run -dp 3000:3000 \
    -w /app -v "$(pwd):/app" \
    node:12-alpine \
    sh -c "yarn install && yarn run dev"
```

- -w /app == sets the container's pwd where the command will print from
- -v "$(pwd):/app" == bind mount the host's present app dir to the container's /app dir
- node:12-alpine == the image to use
	- base image from the app from the dockerfile
- sh -c "yarn install && yarn run dev" == the command to be executed
	- start an sh shell
	- run yarn install to instal dependencies
	- yarn run dev to run it
		- inside the package.json, the dev script is starting nodemon

Watch logs using:

`docker logs -f <id>`

Make changes to your app, stop it and then rebuild the new image with the old `docker buiild -t <img_name>`.

Using bind mounts is very common for local dev setups. The advantage is that the dev machine doesn't need to have all of the build tools/environments installed
- With the single `docker run`; the dev env is pulled/ ready to go
- docker compose will help simplify these commands

## Multi-Container Apps


