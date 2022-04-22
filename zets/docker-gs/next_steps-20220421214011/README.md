# What's Next?

## Container Orchestration

Running containers in prod is tough
- Using a single machine to run docker run / docker-compose up can run into issues when containers die or you want to scale across several machines

Container orchestration helps solve this problem; kubernetes, swarm, nomad & ecs all help in different ways
- General idea: have a manager who receive expected state; this state might be "run 2 instances of my web app and expost port 80"
- The managers look at all machines in a cluster and delegate work to "worker" nodes
- Managers watch for changes (like a container quitting) and work to make actual state reflect expected state

## Cloud Native Computing Foundation Projects

[Graduated/Incubated Projects](https://www.cncf.io/projects/)

[CNCF Landscape](https://landscape.cncf.io/)
