---
sidebar_position: 3
title: Docker Commands
description: A collection of Docker commands to help you get started with Docker.
tags: ["Docker", "Containerization", "DevOps"]
keywords: ["Docker", "Containerization", "DevOps"]
slug: "/docker/commands"
---

### Docker Basics

- Check the Docker version

```bash
docker version
```

- To check all the available images

```bash
docker images
```

- Pull or download an image from a registry to your local machine.

```bash
docker pull <image name>
//Eg: docker pull nginx
```

- Run a container. Docker will pull the image first if it is not available locally.
  - NOTE: If you only provide the image name, Docker uses the default tag, usually `latest`. You can also specify a version such as `nginx:1.14`.
  - Common flags:
  
     - `--name <name> `- To give a name to the container.
     - `-p <Host port:container port>`- To forward the port.
     - `-d` - To run in detached mode
     - `-it` - For an interactive environment
     - `-e` - For environment variable

```bash
docker run <image-name>
//Eg: docker run nginx
```

- We can also pass a complete `.env` file

```bash
--env-file <path-to-env-file>
Eg: --env-file ./.env
```

### Docker Container

- To stop a running container

```bash
docker stop <container-ID/name>
``` 

- To resume a stopped container

```bash
docker start <container-ID/name>
```

- To check the running processes inside a container.

```bash
docker top <container-name/id>
```

- To check stats of running container.

```bash
docker stats <container-name/id>
```

- Check the config and info of a container.

```bash
docker inspect <container-name/id>
//Eg: docker inspect mynginx
```

- Check all the container running.

```bash
docker ps
or
docker container ls
```

- To start and interactive session and attach terminal with the container.

  - NOTE: every image does not support `bash` so we should use `sh`

```
docker exec -it <container-ID/name> bash/sh
```

- To check which ports has been exposed and forwarded

```bash
docker port <image-name>
```

- To check all the containers (stopped, paused and running)

```bash
docker ps -a
```

- Check logs of a container

```bash
docker logs <container-ID/name>
```

- Delete all the stopped containers

```bash
docker container prune -f
```

- Delete all the containers (stopped, paused and running)
```bash
docker rm -f $(docker ps -aq)
```
- Delete all the images 
```bash
docker rmi -f $(docker images -q)
```

- Remove unused images
```bash
docker image prune -all
```

- Auto cleanup when the container exits

```bash
 docker container run —rm <image-name>
```

### Docker Network

- Check the list of available networks.

```bash
docker network ls
```

- Inspect a network components, like which container are attached to that network.

```bash
docker network inspect <network-name>
```

- Run a container on a specific or user-created network.

```
docker run --network <network-name> <image-name>
```

```
docker inspect --format "{{.NetworkSettings.IPAddress}}" <container-name>
```

### Docker Images

- Remove an image

```bash
docker rmi <image name> -f
```

- Remove all the images at once

```bash
docker rmi $(docker images -q)
```

- To inspect an image layers and other info

```bash
docker inspect  <image-name/id>
```

- Check the image layers formation

```bash 
docker history <image-name/id>
```

- Create your own tagged image from an existing image.

```
docker image tag <image-name:tag> <new-image-name:tag>
docker image tag nginx pradumna/nginx:hello
docker image tag ubuntu:18.04 pradumna/ubuntu:example
```

### Docker Volume

- Create a bind mount.
  - This helps you sync local files with a container.
  

- Sync local changes into a container with a bind mount.
    - `-v` is used to define the mount. You can also mount a second path as read-only when needed.

```bash
docker run -v <path-to-folder-on-local-machine>:<path-to-folder-on-container> -p <host-port>:<container-port> -d --name docker-node docker-node
docker
```

```bash
docker run -v <path-to-folder-on-local-machine>:<path-to-folder-on-container> -v <path-to-file/folder-on-container> -p <local-machine-port>:<container-port> -d --name docker-node docker-node
```
To make the mount read-only, add `:ro` at the end of the mount definition.

- Mount the Docker socket into a container when you need that container to talk to the host Docker daemon, for example in some Jenkins setups.
```bash
docker run -it -v /var/run/docker.sock:/var/run/docker.sock docker-image:version bin/bash
```

### Docker Compose

- Run a Docker Compose project.
  Note: If you do not pass `-f`, Docker looks for `compose.yaml` or `docker-compose.yaml` in the working directory.

```bash
docker compose up -d
```

```bash
docker compose down
```

- Rebuild the image with new changes.

```bash
docker compose up --build
```

- Combine multiple Compose files.

```bash
docker compose -f compose.yaml -f compose.dev.yaml
```

###  Docker Swarm and Services

- Initialize Swarm

```bash
docker swarm init
```

- Check all available nodes

```bash
docker node ls
```

- To add a node as a manager

```bash
docker node update --role manager <node-name>
```

- To create an overlay network

```bash
docker network create -d overlay backend
```

- Create a service. Also we can add flags for further customization.

    - `--name` - to give a service name
    - `--replicas` - to define how many running instance of the same image.
    - `-p` - for port forwarding

```bash
docker service create -p 8080:80 --name vote --replicas 2 nginx
```

- Get all task containers running across nodes

```bash
docker service ps <service-name/id>
```

> SERVICE UPDATE

- To scale up the service (i.e increasing the no of replicas)

```bash
docker service scale <service name>=<no to scale>
docker service scale mynginx=5
```

- To update the image in running service

```bash
docker service update --image <updated image> <service name>
docker service update --image mynginx:1.13.6  web
```

- To update the port

We can't directly update the port We have to add and remove the ports

```
docker service update --publish-rm 8080 --publish-add 808180 <service name>
docker service update --publish-rm 8080 --publish-add 808180 mynginx
```

### Docker Stack

- To deploy a stack file

```bash
docker stack deploy -c <file-name.yaml> <stackname>
```

- To remove running stack

```
docker stack rm <stack name>
```

- To check list of stacks running

```
docker stack ls
```

**STACK -> SERVICES -> TASKS -> CONTAINERS**

- Check which services are running inside a stack

```
docker stack services <stack name>
```

- Check which tasks are running inside a stack

```
docker stack ps <stack name> 
```

> Registry

```
127.0.0.0:5000/v2/_catalog
```

### Tips and Short hands

- Run the command with the container creation

```bash
docker run <image-name> <command>
// Eg: `docker run ubuntu:16.04 echo hey`
```


- Creating our Own image and container.

```
Step 1 - create Dockerfile
Step 2 - docker build -t myimage:1.0 <path-of-dockerfile> (-t for tag)
Step 3 - docker run myimage:1.0
```

## Read next

- [Docker Introduction](./introduction.md) - Review the core Docker concepts behind these commands.
- [Learning Resources](./learning-resources.md) - Continue with courses and videos for deeper practice.
- [Other Resources](./other-resources.md) - Explore sample Dockerfiles and Compose files.
