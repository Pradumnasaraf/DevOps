### Docker Basic

- To check Docker vesrion

```
docker version
```

- To check all the images available

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
     - `--name <name> `- To give a name to the conatiner.
     - `-p <Hot port:container port>`- To fowrad the port.
     - `-d` - To run in detached mode
     - `-it` - For interactive envirnoment

```bash
docker run <image name>
//Eg: docker run nginx
```

### Docker Container

- To stop a running conatiner

```bash
docker stop <container ID/name>
``` 

- To resume a stopped conatiner

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

- Check the config and info of a conatiner.

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

- To start and interactive session and get inide the container.

  - NOTE: every image does not support `bash` so we use `sh`

```
docker exec -it <container ID/name> bash/sh
```

- To check which ports has been exposed and forwarded

```bash
docker port <container name>
```

- Check all the stopped container

```bash
docker ps -a
```

- Check logs of a container

```bash
docker logs <container ID/name>
```

- Delete all the stopped conatiner

```bash
docker container prune -f
```
- Auto cleanup when t

```
 docker container run —rm

### Docker Network

- Check list of avilable networks.

```bash
docker network ls
```

- Inspect a network components

```bash
docker network inspect <network name>
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

### Docker Volume

- Create bind mount
  - Help to sync our local files with help of Docker container.
  

- To sync our local machine chnages with help of Docker volume (Bind mount)
    - `- v` is use to define volume, aslo we give another `-v` flag to override the changes so that it will not chnage in container.

```bash
docker run -v <path-on-folder-loacl-machine>:<path-to-folder-on-container> -p <local-machine-port>:<container-port> -d --name docker-node docker-node
docker
```

```bash
docker run -v <path-on-folder-loacl-machine>:<path-to-folder-on-container> -v <path-to-file/folder-on-conatiner> -p <local-machine-port>:<container-port> -d --name docker-node docker-node
```
To make it read only so that when you add some files inside it the container and it will not get created on you local machine use `-v port:port:ro`


- To override the and ENV of a docker container, here PORT
```
-e PORT=3500
```
or file

```bash
--env-file <path-to-env-file>
Eg: --env-file ./.env
```

TO run docker compose file

```bash
docker compose up -d
```

```bash
docker compose down
```
- When we run docker compose while with the existing image it will not create build the image even tho there is some changes. It runs the stale version.

```
docker compose up --build
```

To override the existing config:

```bash
docker compose -f docker-compose.yaml  -f docker-compose.dev.yaml
```

- To list all the networks

```bash
docker network ls
```

To inspect a particular network

```bash
docker inspect <network-id>
```

to check which port are exposed in a container

```bash
docker container port <Container-name>
````

```
docker inspect --format "{{.NetworkSettings.IPAddress}}" <conatiner-name>
```

- To inspect which conatiners are attached to the a particalr newtork

```
docker network inspect <network-name>
docker network inspect bridge
```

- To create a network (It will create a bridge)

```
docker network create <network-name>
```

- To run a container on a certian network/own careted network 

```
docker run --network <network-name> <image-name>
```

- To connect a conatanier to a another network

```
docker connect network <network-name> <conatiner-name>
```

- To disconnect a conatanier from another network

```
docker disconnect network <network-name> <conatiner-name>
```

- Check the image layers formation

```bash 
docker history <image-name>
```

- Inspect the meta data of an image

```
docker inspect <image-name>
```

- To create a our own tag with some image

```
docker image tag <image-name with tag> <new-image name with tag>
docker image tag nginx pradumna/nginx:hello
```

> DOCKER SWARM and SERVICE

- To initalize swarm

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

- To craete a service

    - `--name` - to give a service name
    - `--replicas` - to define how many running instance of the same image.
    - `-p` - for port forwarding

```
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

We ca't direclty update the port We have to add and remove the ports

```
docker service update --publish-rm 8080 --publish-add 808180 <service name>
docker service update --publish-rm 8080 --publish-add 808180 mynginx
```

> DOCKER STACK

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
