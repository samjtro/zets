# Using [Docker Compose](https://docs.docker.com/compose/)

Docker Compose is a tool that allows you to create a YAML file, to define the services and spin everything up/tear down.

The advantage of using compose is that you can define the entire app stack in the file, keep it at the root of the project's repo, and enable someone to contribute to the project. Someone just needs to clone the repo and start the compose app.

## Creating our compose file

0. At the root of the app project, create a file named `docker-compose.yml`
1. Start by defining the schema version. The [Compose file reference](https://docs.docker.com/compose/compose-file/) shows the current schema versions and the compatibility matrix

`version: "3.8"`

2. Then, define the list of services/containers we want to run as part of the app

`services: `

Now, we migrate those services to the compose file.

#### Defining the App Service

0. Define the service entry and the image for the container; the name for the service can be anything which then becomes a network alias for defining MySQL services
1. The command closes to the image definiton, however there is no requirement on ordering
2. -p 3000:3000 is migrated here using [short syntax](https://docs.docker.com/compose/compose-file/compose-file-v3/#short-syntax-1), but you could also use [long syntax](https://docs.docker.com/compose/compose-file/compose-file-v3/#long-syntax-1)
3. Then, migrate the working directory (-w /app) + volume mapping (-v "$(pwd):/app") by using the `working_dir` + `volumes` definitions. Volumes also have [short](https://docs.docker.com/compose/compose-file/compose-file-v3/#short-syntax-3) / [long](https://docs.docker.com/compose/compose-file/compose-file-v3/#long-syntax-3) syntax
- Docker compose volume definitions allow relative paths from the current directory, as seen above
4. Finally; we migrate the environment var definitions using the `environment` key

```
services:
  app:
    image: node:12-alpine
    command: sh -c "apk --no-cache --virtual build-dependencies add python2 make g++ && yarn install && yarn run dev"
    ports:
      - 3000:3000
    working_dir: /app
    volumes:
      - ./:/app
    environment:
      MYSQL_HOST: mysql
      MYSQL_USER: root
      MYSQL_PASSWORD: secret
      MYSQL_DB: todos
```

#### Defining the MySQL service

0. Define a new service named "mysql" under "app" - this automatically gets the network alias mysql. At the same time spec the image to use
1. Define the volume mapping - when ran thru `docker run`, the named volume is created automatically - with docker compose, we need to define the volume in the top-level volumes: section; then spec the mountpoint in the service config
- Providing only the volume name means default options are used: there are [other](https://docs.docker.com/compose/compose-file/compose-file-v3/#volume-configuration-reference) options though
2. Finally, spec the env vars

```
version: "3.8"

services:
  app:

  mysql:
    image: mysql:5.7
    volumes:
      - todo-mysql-data:/var/lib/mysql
    environment:
      MYSQL_ROOT_PASSWORD: secret
      MYSQL_DATABASE: todos

volumes:
  todo-mysql-data:
```

## Running the App Stack

`docker-compose up -d` points to the docker-compose.yml file (-d flags discrete to run in the background)

> `docker-compose logs -f` shows each service interweaved into a single stream (-f gives a live output for debuffing)

## Tear it down

`docker-compose down`
