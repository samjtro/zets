# [Docker - Getting Started p4]()

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


