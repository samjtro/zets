# [Docker - Getting Started p6]()

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
