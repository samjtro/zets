# [Beginner Boost 2021 Day 5](https://www.youtube.com/watch?v=YcBIUOSOAeg&list=PLrK9UeDMcQLre1yPasCnuKvWvyXKzmKhW&index=6)

## What's a container and why should you care?

["kubernetes is the new operating system"](https://www.infoworld.com/article/3322120/sorry-linux-kubernetes-is-now-the-os-that-matters.html)

**Timeline:**

0. development in the linux kernel allowed you to seperate the kernel services and hardware from the processes, files, etc.
1. at the same time, virtualization is occuring
- running the operating system with a hypervisor, and apps inside of that os
- you could run clusters of vms for deployment
2. before any of this -- chroot was the environment people used to isolate services or programs into its own mem/fs (linux jail)

> TO LEARN: diff between a container and a virtual machine?

**Containers allow for management of server-side microservices**

- allows you to combine all components of the back-end application into one component
- seperation of concerns & encapsulation also allows for secure operations
- a container takes all of that, encapsulates it, and allows for running of the program without giving explicit access
- can punch holes through it -- can let it access one file, or one socket
- but MUST EXPLICITLY POKE THOSE HOLES
- by default, it has access to itself

**Containers are images, which allow you to bundle up all runtime neccessities while writing project code for deployment**

- it has been created, the image is similar to a tarball
- not active until you run it: when it's been run, either by kubernetes or docker or containerd
- the resources are downloaded and installed, ports are established, and it is running == process
- you can pause that process, and restart it whenever you want

**Kubernetes is an operating system, cont.**

- each sensor is a container
- those containers connect as nodes to a structure that is deployed and managed with k8s
- graph theory comes into play with large scale node architectures
- kubernetes architecture; nodes can all communicate with themselves and pass data
- edge computing is what happens when the car uses all sensors as nodes in the cluster; or the mars rover, sensors running as nodes communicate through mesh
- k8s is the  architecture for computers to talk to each other in a variety of ways

**Cybersecurity**

- build home/cloud/hybrid k8s cluster
- point nodes to internet, syncing communication and cracking vulnerabilities to make money
- can create supercomputer power clusters with aws/gcp

**Entry points**

- application that runs, different from a command

> Learn YAML
> Learn POSIX Shell - For DOCKERFILE
> Makes much more sense to learn shell before OOP
> Forces you to understand how a computer thinks v High-level paradigms
