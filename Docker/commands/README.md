### Docker Basic

- To check Docker vesrion

```
docker version
```

- To check all the available images

```bash
docker images
```

- Pull/Downlaod the image from the Docker registry to local machine.

```bash
docker pull <image name>
//Eg: docker run nginx
```

- To run an container (It will 1st pull the image if not present in the local sytem)
  - NOTE: When we just provide the name of the image it will pull the lastest one, i.e `nginx:latest`. We can also specify the version `nginx:1.14`
  - Additioanly we can use flags
  
     - `--name <name> `- To give a name to the container.
     - `-p <Host port:container port>`- To fowrad the port.
     - `-d` - To run in detached mode
     - `-it` - For interactive envirnoment
     - `-e` - For environment variable

```bash
docker run <image name>
//Eg: docker run nginx
```

- We can aslo pass a complete `.env` file

```bash
--env-file <path-to-env-file>
Eg: --env-file ./.env
```

### Docker Container

- To stop a running container

```bash
docker stop <container ID/name>
``` 

- To resume a stopped container

```bash
docker start <container ID/name>
```

- To check the running processes inside a container.

```bash
docker top <container name/id>
```

- To check stats of running container.

```bash
docker stats <container name/id>
```

- Check the config and info of a container.

```bash
docker stats <container name/id>
//Eg: docker inspect mynginx
```

- Check all the container running.

```bash
docker ps
or
docker container ls
```

- To start and interactive session and get inside the container.

  - NOTE: every image does not support `bash` so we use `sh`

```
docker exec -it <container ID/name> bash/sh
```

- To check which ports has been exposed and forwarded

```bash
docker port <image name>
```

- To check all the stopped container

```bash
docker ps -a
```

- Check logs of a container

```bash
docker logs <container ID/name>
```

- Delete all the stopped container

```bash
docker container prune -f
```
- Auto cleanup when t

```bash
 docker container run —rm
```

### Docker Network

- Check list of avilable networks.

```bash
docker network ls
```

- Inspect a network components, like which container are attached to that network.

```bash
docker network inspect <network name>
```

- Run a container on a certian network/own careted network 

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
docker inspect  <image name/id>
```

- Check the image layers formation

```bash 
docker history <image-name>
```

- Create a our own image with an existing image.

```
docker image tag <image-name with tag> <new-image name with tag>
docker image tag nginx pradumna/nginx:hello
```

### Docker Volume

- Create bind mount
  - Help to sync our local files with help of Docker container.
  

- To sync our local machine changes with help of Docker volume (Bind mount)
    - `- v` is use to define volume, aslo we give another `-v` flag to override the changes so that it will not chnage in container.

```bash
docker run -v <path-on-folder-loacl-machine>:<path-to-folder-on-container> -p <local-machine-port>:<container-port> -d --name docker-node docker-node
docker
```

```bash
docker run -v <path-on-folder-loacl-machine>:<path-to-folder-on-container> -v <path-to-file/folder-on-container> -p <local-machine-port>:<container-port> -d --name docker-node docker-node
```
To make it read only so that when you add some files inside it, the container will not get created on your local machine use `-v port:port:ro`

### Docker Compose

- Run docker compose file.
  Note: By default it finds for the file name `docker-compose.yaml`, to give file with other naming use `-f <file-name.yaml>` command

```bash
docker compose up -d
```

```bash
docker compose down
```

- To rebuilt the new Image with thew new changes

```bash
docker compose up --build
```

- Override the existing of compose file

```bash
docker compose -f docker-compose.yaml  -f docker-compose.dev.yaml
```

###  Docker Swam and Services

- Initalize swarm

```bash
docker swarm init
```

- Check all the node available

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

- Create a service. Also we can add flags for further customiztaion.

    - `--name` - to give a service name
    - `--replicas` - to define how many running instance of the same image.
    - `-p` - for port forwarding

```bash
docker service create -p 8080:80 --name vote --replicas 2 nginx
```

- To get all task containers running on different node 

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

We can't direclty update the port We have to add and remove the ports

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

- To check which services are running inside a staacks

```
docker stack services <stack name>
```

- To check taks are running inside a stack

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
doc run <image-name> <command>
// Eg: `doc run ubuntu:16.04 echo hey`
```


- Creating our Own image and container.

```
Step 1 - create Dockerfile
Step 2 - docker build -t myimage:1.0 <location> (-t for tag)
Step 3 - docker run
```
