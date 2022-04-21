#[Beginner Boost 2021 Day 5](https://www.youtube.com/watch?v=YcBIUOSOAeg&list=PLrK9UeDMcQLre1yPasCnuKvWvyXKzmKhW&index=6)

##Whats a container and why should you care?

"kubernetes is the new operating system" [Link](https://www.infoworld.com/article/3322120/sorry-linux-kubernetes-is-now-the-os-that-matters.html)

**timeline:**

0. development in the linux kernel allowed you to seperate the kernel services and hardware from the processes, files, etc.
1. at the same time, virtualization is occuring
- running the operating system with a hypervisor, and apps inside of that os
- you could run clusters of vms for deployment
2. before any of this -- chroot was the environment people used to isolate services or programs into its own mem/fs (linux jail)

> TO LEARN: diff between a container and a virtual machine?

**containers allow for management of server-side microservices**
- allows you to combine all components of the back-end application into one component
- seperation of concerns & encapsulation also allows for secure operations

**a container takes all of that, encapsulates it, and allows for running of the program without giving explicit access**
- can punch holes through it -- can let it access one file, or one socket
- but MUST EXPLICITLY POKE THOSE HOLES
- by default, it has access to itself

> containers allow you to bundle up all runtime neccessities while writing project code for deployment

##Install Docker


