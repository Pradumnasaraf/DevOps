## Concepts


### Docker Images

- Images are made up of app binaries, dependencies, and metadata. Don't contain a full OS.
- Images are combination of mutiple layers.
- Each Image has it down unique ID and a tag for different version.

![Screenshot from 2022-11-02 11-57-19](https://user-images.githubusercontent.com/51878265/199414178-d59e8780-c140-4bf1-b27e-7e8f1c723afb.png)

### Dockerfile

> Common Commands:

- `FROM` (base image)
- `ENV` (environment variable)
- `RUN` (any arbitrary shell command)
- `EXPOSE` (open port from container to virtual network)
- `CMD` (command to run when container starts) 
