---
sidebar_position: 1
title: GitHub Actions Introduction
description: A guide to learn about GitHub Actions.
tags: ["GitHub Actions", "DevOps", "CI/CD"]
keywords: ["GitHub Actions", "DevOps", "CI/CD"]
slug: "/github-actions"
---

GitHub Actions lets you automate software workflows directly inside a GitHub repository. You define workflows in YAML files, trigger them from events such as `push`, `pull_request`, or `workflow_dispatch`, and run jobs on GitHub-hosted or self-hosted runners.

### Overview

- [Docs](https://docs.github.com/en/actions/learn-github-actions/understanding-github-actions)

- **Workflow** - A workflow is a configurable automated process made up of one or more jobs. Workflows are defined in `.yml` files in the `.github/workflows` directory of your repository.

- **Jobs** - A job is a set of steps that execute on the same runner. Runner is a server that has the GitHub Actions runner application installed.

- **Steps** - A step is an individual task that can run shell commands or reusable actions like `actions/checkout@v4`. Each step in a job executes on the same runner, so files are available across steps in that job.

> Summary: The workflow is a set of jobs and each job is a set of steps. Each step can be an action or a shell command.

Basic workflow file:

```yaml
name: CI

on:
  push:
    branches: [main]
  pull_request:
    branches: [main]

jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - name: Run a one-line script
        run: echo "Hello, world!"
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

- We can trigger the workflow only when certain files are changed.

```yaml
on:
  push:
    paths:
      - "src/index.js"
```

- We can trigger the workflow only when certain file types are changed.

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

- We can trigger the workflow only when files in a specific directory are changed.

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

Services are containers that GitHub Actions starts for a job. They are useful for integration tests that need supporting dependencies such as Redis or Postgres.

```yaml
services:
  redis:
    image: redis:7
    ports:
      - 6379:6379

  postgres:
    image: postgres:16
    env:
      POSTGRES_PASSWORD: postgres
    ports:
      - 5432:5432
```

### Using output from one step in another step

We can use the output from one step in another step using the `id` keyword.

```yaml
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - name: Get the version
        id: get_version
        run: echo "version=1.0.0" >> "$GITHUB_OUTPUT"

      - name: Use the version
        run: echo ${{ steps.get_version.outputs.version }} # steps.<step_id>.outputs.<output_name>
```

GitHub now recommends environment files such as `GITHUB_OUTPUT` instead of the deprecated `::set-output` command.

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

We can set environment variables in the workflow file using the `env` keyword.

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

Matrix is helpful when we want to run the same job on multiple versions of a tool or across multiple operating systems.

```yaml
jobs:
  example_matrix:
    strategy:
      matrix:
        os: [ubuntu-22.04, ubuntu-24.04]
        version: [20, 22]
    runs-on: ${{ matrix.os }}
    steps:
      - uses: actions/setup-node@v4
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
      - uses: actions/checkout@v4
      - id: deploy-preview
        run: echo "url=preview_url" >> "$GITHUB_OUTPUT"

  data:
    runs-on: ubuntu-latest
    needs: deploy
    steps:
      - name: Load preview URL
        run: echo ${{ needs.deploy.outputs.url }}
```

Useful references:

- [Understanding GitHub Actions](https://docs.github.com/en/actions/learn-github-actions/understanding-github-actions)
- [Workflow syntax for GitHub Actions](https://docs.github.com/en/actions/reference/workflow-syntax-for-github-actions)
- [Workflow commands and `GITHUB_OUTPUT`](https://docs.github.com/en/actions/reference/workflows-and-actions/workflow-commands)

## Artifacts

It allows you to share data between jobs in a workflow and store data once the workflow has been completed. We can use the `upload-artifact` and `download-artifact` actions to upload and download artifacts.

```yaml
name: Share data between jobs

on: [push]

jobs:
  job_1:
    name: Generate Random Number
    runs-on: ubuntu-latest
    steps:
      - shell: bash
        run: |
          echo $((RANDOM % 100)) > random-number.txt
      - name: Upload random number for job 1
        uses: actions/upload-artifact@v4
        with:
          name: random_number
          path: random-number.txt

  job_2:
    name: Echo Random Number
    needs: job_1
    runs-on: ubuntu-latest
    steps:
      - name: Download random number from job 1
        uses: actions/download-artifact@v4
        with:
          name: random_number
      - shell: bash
        run: |
          number=$(cat random-number.txt)
          echo The random number is $number
``` 


### What's next?

- [Scenarios](./scenarios.md) - A collection of GitHub Actions workflow files I use and created to help you understand the concepts better.
- [Learning Resources](./learning-resources.md) - A list of resources to learn more about GitHub Actions.
- [Other Resources](./other-resources.md) - A list of extra resources that you can use to learn more about GitHub Actions.
