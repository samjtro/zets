# [Docker - Getting Started p1](https://docs.docker.com/get-started/)

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


