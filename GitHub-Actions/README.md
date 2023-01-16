## GitHub Actions

GitHub Actions is a feature that allows you to automate your software development workflows. You can write individual tasks, called actions, and combine them to create a custom workflow. Workflows are custom automated processes that you can set up in your repository to build, test, package, release, or deploy any code project on GitHub.

### Resources

- [GitHub Actions]( https://docs.github.com/en/actions )
- [GitHub Actions: Getting Started]( https://docs.github.com/en/actions/learn-github-actions/introduction-to-github-actions )
- [GitHub Action Tutorial - Tech World With Nana](https://youtu.be/R8_veQiYBjI) 

### Overview

- [Docs](https://docs.github.com/en/actions/learn-github-actions/understanding-github-actions)

- **Workflow** - A workflow is a configurable automated process made up of one or more jobs. Workflows are defined in `.yml` files in the `.github/workflows` directory of your repository.

- **Jobs** - A job is a set of steps that execute on the same runner. Runner is a server that has the GitHub Actions runner application installed. 

- **Steps** - A step is an individual task that can run commands or actions like `actions/checkout@v2`. Each step in a job executes on the same runner, allowing for direct file sharing.

> Summary: The workflow is a set of jobs and each job is a set of steps. Each step can be an action or a shell command.

Basic worflow file.

```yaml
name: CI # name of the workflow

on: [push] # triggers the workflow on push or pull request events but only for the master branch

jobs:
  build: # name of the job
    runs-on: ubuntu-latest # runs-on is the type of machine to run the job on - runner
    steps: # steps are the individual tasks that make up a job

    - uses: actions/checkout@v2 
    - name: Run a one-line script # name is the name of the step
      run: echo Hello, world!
```

- **Action** - An action is a custom application for the GitHub Actions platform that performs a complex but frequently repeated task.

- **Event** - An event is a specific activity that triggers a workflow. For example, activity that occurs on GitHub, such as opening a pull request or pushing a commit.

### Triggers

The `on` keyword is used to trigger the workflow. It can be triggered by a push, pull request, or a schedule. 

- [Docs](https://docs.github.com/en/actions/reference/events-that-trigger-workflows)

```yaml
on:
  push:
    branches: [ main ]
  pull_request:
    branches: [ main ]
```

#### Trigger Filters

- We can trigger the workflow when only when certain files are changed.

```yaml
on:
  push:
    paths:
      - 'src/index.js'
```

- We can trigger the workflow when only when certain files with an extension are changed.

```yaml
on:
  push:
    paths:
      - '**.js'
      - '**.css'
  pull_request:
    paths:
      - '**.js'
```

- We can trigger the workflow when only when certain files in a directory are changed.

```yaml
on:
  push:
    paths:
      - 'src/**'
```

### Services

Services are external resources that can be used in a workflow. We can use services like `redis` or `postgres` in our workflow. We can run a service in a container and use it in our workflow. We can also use custom images.

```yaml
services:
  redis:
    image: redis # Docker image
    ports:
      - 6379:6379 # port mapping
  
  mongodb:
    image: mongo
    ports:
      - 27017:27017

  sqlite:
    image: sqlite
    ports:
      - 3306:3306
```

### Using ouput from one step in another step

We can use the output from one step in another step using the `id` keyword.

```yaml
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - name: Get the version
        id: get_version
        run: echo ::set-output name=version::1.0.0

      - name: Use the version
        run: echo ${{ steps.get_version.outputs.version }} # {{ steps.<step_id>.outputs.<output_name> }}
```


### GitHub context

The `github` context is available to you in any workflow or action you create on GitHub. You can use the context to get information about the workflow run, the repository, the event that triggered the workflow run, and more.

