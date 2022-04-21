# [Docker - Getting Started p5]()

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
