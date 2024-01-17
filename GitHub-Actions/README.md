## GitHub Actions

GitHub Actions is a feature that allows you to automate your software development workflows. You can write individual tasks, called actions, and combine them to create a custom workflow. Workflows are custom automated processes that you can set up in your repository to build, test, package, release, or deploy any code project on GitHub.

### Resources

- [GitHub Actions](https://docs.github.com/en/actions)
- [GitHub Actions: Getting Started](https://docs.github.com/en/actions/learn-github-actions/introduction-to-github-actions)
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
    branches: [main]
  pull_request:
    branches: [main]
```

#### Trigger Filters

- We can trigger the workflow when only when certain files are changed.

```yaml
on:
  push:
    paths:
      - "src/index.js"
```

- We can trigger the workflow when only when certain files with an extension are changed.

```yaml
on:
  push:
    paths:
      - "**.js"
      - "**.css"
  pull_request:
    paths:
      - "**.js"
```

- We can trigger the workflow when only when certain files in a directory are changed.

```yaml
on:
  push:
    paths:
      - "src/**"
```

- Manual trigger

```yaml
on:
  workflow_dispatch:
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

### Jobs

We can have multiple jobs in a workflow. Each job runs in parallel by default. We can also run jobs sequentially using the `needs` keyword. Each step in a job runs on the same runner.

We can run a job when a job passes or fails using the `needs` keyword.

```yaml
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - name: Get the version
        run: echo "Hello World"

  test:
    needs: build
    runs-on: ubuntu-latest
    steps:
      - name: Get the version
        run: echo "Hello World"
```

### GitHub context

The `github` context is available to you in any workflow or action you create on GitHub. You can use the context to get information about the workflow run, the repository, the event that triggered the workflow like pull request number, etc.

**Some useful properties of the `github` context are:**

- `github.event_name` - The name of the webhook event that triggered the workflow.
- `github.sha` - The commit SHA that triggered the workflow.
- `github.ref` - The branch or tag ref that triggered the workflow.
- `github.repository` - The owner and repository name. For example, `Pradumnasaraf/DevOps`.
- `github.actor` - The name of the person or app that initiated the workflow.
- `github.job` - The name of the job that's currently running.
- `github.run_number` - A unique number for each run of a particular workflow in a repository. This number begins at 1 for the workflow's first run, and increments with each new run. This number does not change if you re-run the workflow run.
- `github.run_id` - A unique number for each run of any workflow in a repository. This number begins at 1 for the workflow's first run, and increments with each new run. This number does not change if you re-run the workflow run.
- `github.workflow` - The name of the workflow.
- `github.action` - The unique identifier (id) of the action.
- `github.event` - The event payload. For example, the issue or pull request object that triggered the workflow run.

### Environment variables

We can set environment variables in the workflow file using the `env` keyword. WE

```yaml
env:
  MY_NAME: "Pradumna"
```

### Secrets

We can store sensitive data like API keys, passwords, etc. in the repository settings. We can access the secrets using the `secrets` context.

```yaml
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - name: Get the version
        run: echo ${{ secrets.MY_SECRET }}
```

### Matrix

Matrix is helpful when we want to run a job on multiple versions of a tool.For eg. we want to run a job on multiple versions of Node.js. We can use the `matrix` keyword to run a job on multiple versions of a tool. We can also use the `matrix` keyword to run a job on multiple operating systems.

```yaml
jobs:
  example_matrix:
    strategy:
      matrix:
        os: [ubuntu-22.04, ubuntu-20.04]
        version: [14, 16, 18]
    runs-on: ${{ matrix.os }}
    steps:
      - uses: actions/setup-node@v3
        with:
          node-version: ${{ matrix.version }}
```

### Outputs

We can use the `outputs` keyword to output data from a job. We can use the output from one job in another job using the `needs` keyword.

```yaml
jobs:
  deploy:
    runs-on: ubuntu-latest
    outputs:
      url: ${{ steps.deploy-preview.outputs.url }}
    steps:
      - uses: actions/checkout@v3
      - id: deploy-preview
        run: echo "url=preview_url" >> $GITHUB_OUTPUT

    data:
      runs-on: ubuntu-latest
      needs: deploy
      steps:
        - name: load json files
          run: echo ${{ needs.deploy.outputs.url }}
```


